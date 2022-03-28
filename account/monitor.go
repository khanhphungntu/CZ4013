package account

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"time"
)

type clientsTracking struct {
	lock           *sync.Mutex
	clients        map[*net.UDPAddr]time.Time
	schedulePeriod time.Duration
	ser            *net.UDPConn
}

func (c *clientsTracking) registerClient(addr *net.UDPAddr, expire time.Time) {
	c.lock.Lock()
	c.clients[addr] = expire
	c.lock.Unlock()
}

func (c *clientsTracking) cleanUp() {
	c.lock.Lock()
	var toDeleteKeys []*net.UDPAddr
	for key, value := range c.clients {
		if time.Now().After(value) {
			toDeleteKeys = append(toDeleteKeys, key)
		}
	}

	for _, key := range toDeleteKeys {
		fmt.Println("Deleting key", key)
		delete(c.clients, key)
	}

	c.lock.Unlock()
}

func (c *clientsTracking) scheduleCleanUp() {
	go func() {
		ticker := time.NewTicker(c.schedulePeriod)

		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Cleaning up", t)
				c.cleanUp()
			}
		}
	}()
}

func (c *clientsTracking) dispatchEvent(content []byte) {
	fmt.Println("Dispatching event...")

	var packet = make([]byte, 2)
	binary.BigEndian.PutUint16(packet, uint16(len(content)))
	packet = append(packet, content...)

	c.lock.Lock()
	for client, _ := range c.clients {
		_, err := c.ser.WriteToUDP(packet, client)

		if err != nil {
			fmt.Println("Failed to send event", content, client)
		} else {
			fmt.Println("Send to", client)
		}
	}

	c.lock.Unlock()
}

func newClientsTracking(schedulePeriod time.Duration) clientsTracking {
	c := clientsTracking{
		lock:           &sync.Mutex{},
		clients:        make(map[*net.UDPAddr]time.Time),
		schedulePeriod: schedulePeriod,
	}

	c.scheduleCleanUp()
	return c
}

var clientsTrackingImpl = newClientsTracking(30 * time.Second)

// RegisterServerWithClientMonitor must be called before any client can register
func RegisterServerWithClientMonitor(ser *net.UDPConn) {
	clientsTrackingImpl.ser = ser
}

func RegisterMonitorClient(content []byte, addr *net.UDPAddr) {
	if clientsTrackingImpl.ser == nil {
		panic("Server has not been registered!")
	}
	expire := binary.BigEndian.Uint64(content[:8])
	clientsTrackingImpl.registerClient(addr, time.Now().Add(time.Duration(expire)*time.Second))
}

func DispatchEvent(content []byte) {
	clientsTrackingImpl.dispatchEvent(content)
}
