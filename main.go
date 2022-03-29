package main

import (
	"DS/account"
	"fmt"
	"net"
)

func main() {
	addr := net.UDPAddr{
		Port: 8000,
		IP:   net.ParseIP("0.0.0.0"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}

	account.RegisterServerWithClientMonitor(ser)
	proxy := &Proxy{
		Semantic:     1,
		WaitTime:     0,
		RespDropRate: 0,
		ReqDropRate:  50,
	}
	connManager := NewConnectionManager(ser, RouterImpl, proxy)
	connManager.Run()
}
