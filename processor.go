package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

const PacketSize = 1024

type SemanticChoice int

const (
	AtLeastOneSemantic SemanticChoice = 0
	AtMostOneSemantic                 = 1
)

// Router is the function registered to receive the content of the request, this function returns a serialized response
type Router func(content []byte, addr *net.UDPAddr) []byte

type ConnectionManager interface {
	Run()
}

type Proxy struct {
	Semantic SemanticChoice
	WaitTime int64
}

func (p *Proxy) execute() {
	time.Sleep(time.Duration(p.WaitTime) * time.Second)
}

// Packet format
// first two bytes indicate the request id, which is randomly assigned by client -> upto 2^16 id
// third & fourth byte indicates number of byte in the payload -> upto 2^16 byte
// remaining bytes are for the payload of the message
type Packet []byte

type connectionManagerImpl struct {
	ser     *net.UDPConn
	connMap map[uint16]Packet
	router  Router
	proxy   *Proxy
}

func NewConnectionManager(ser *net.UDPConn, router Router, proxy *Proxy) ConnectionManager {
	return &connectionManagerImpl{
		connMap: make(map[uint16]Packet),
		ser:     ser,
		router:  router,
		proxy:   proxy,
	}
}

func (p Packet) getRequestId() uint16 {
	return binary.BigEndian.Uint16(p[:2])
}

func (p Packet) getPacketSize() uint16 {
	return binary.BigEndian.Uint16(p[2:4])
}

func (p Packet) getPacketContent() []byte {
	return p[4 : 4+p.getPacketSize()]
}

// readFromPacket returns the content of the message and a boolean value to indicate if the
// message's payload has fully received or not
func (c *connectionManagerImpl) readFromPacket(p Packet, addr *net.UDPAddr) (resp []byte, reqId uint16) {
	reqId = p.getRequestId()

	if _, ok := c.connMap[reqId]; ok && c.proxy.Semantic == AtMostOneSemantic {
		fmt.Println("Duplicate request detected: ", reqId)
		copy(resp, c.connMap[reqId])
		return
	}

	content := p.getPacketContent()
	resp = c.router(content, addr)

	if c.proxy.Semantic == AtMostOneSemantic {
		c.connMap[reqId] = resp
	}
	return
}

// sendResponse chunks the content into packet
// each packet has format: first two bytes indicate the request id
// next two byte indicates size of the packet -> upto 2^16 bytes in the payload
// remaining bytes are payload
func (c *connectionManagerImpl) sendResponse(content []byte, reqId uint16, addr *net.UDPAddr) {
	var packet = make([]byte, 4)
	binary.BigEndian.PutUint16(packet, reqId)

	size := len(content)
	binary.BigEndian.PutUint16(packet[2:], uint16(size))
	packet = append(packet, content...)

	_, err := c.ser.WriteToUDP(packet, addr)

	if err != nil {
		fmt.Println("Couldn't send response", err)
	}
}

func (c *connectionManagerImpl) Run() {
	for {
		fmt.Println("Waiting for client request")
		p := make([]byte, PacketSize)
		_, remoteAddr, err := c.ser.ReadFromUDP(p)

		c.proxy.execute()
		fmt.Printf("Read a message from %v %s \n\n", remoteAddr, p)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}

		resp, reqId := c.readFromPacket(p, remoteAddr)
		c.sendResponse(resp, reqId, remoteAddr)
	}
}
