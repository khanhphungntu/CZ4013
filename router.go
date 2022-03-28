package main

import (
	"DS/account"
	"net"
)

type serviceType uint8

var (
	REGISTER_ACCOUNT        serviceType = 0
	REGISTER_MONITOR_CLIENT serviceType = 3
)

// RouterImpl : Assume the first byte is used to identify the api
func RouterImpl(content []byte, addr *net.UDPAddr) []byte {
	switch serviceType(content[0]) {
	case REGISTER_ACCOUNT:
		return account.RegisterAccount(content[1:])
	case REGISTER_MONITOR_CLIENT:
		account.RegisterMonitorClient(content[1:], addr)
		return nil
	}

	return nil
}
