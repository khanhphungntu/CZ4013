package main

import (
	"DS/account"
	"fmt"
	"net"
)

// RouterImpl : Assume the first byte is used to identify the api
func RouterImpl(content []byte, addr *net.UDPAddr) []byte {
	switch content[0] {
	case 0:
		return account.RegisterAccount(content[1:])
	case 3:
		account.RegisterMonitorClient(content[1:], addr)
		return nil
	}

	return nil
}

func main() {
	addr := net.UDPAddr{
		Port: 8000,
		IP:   net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}

	account.RegisterServerWithClientMonitor(ser)
	proxy := &Proxy{
		Semantic: 1,
		WaitTime: 3,
	}
	connManager := NewConnectionManager(ser, RouterImpl, proxy)
	connManager.Run()
}
