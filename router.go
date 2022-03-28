package main

import (
	"DS/account"
	"net"
)

type serviceType uint8

var (
	REGISTER_ACCOUNT        serviceType = 0
	DELETE_ACCOUNT          serviceType = 1
	DEPOSIT_WITHDRAW        serviceType = 2
	REGISTER_MONITOR_CLIENT serviceType = 3
)

// RouterImpl : Assume the first byte is used to identify the api
func RouterImpl(content []byte, addr *net.UDPAddr) (res []byte) {
	var monitorRes []byte
	switch serviceType(content[0]) {
	case REGISTER_ACCOUNT:
		return account.RegisterAccount(content[1:])
	case DEPOSIT_WITHDRAW:
		// res, monitorRes = account.DepositWithdraw(content[1:])\
	case DELETE_ACCOUNT:
		res, monitorRes = account.DeleteAccount(content[1:])
	case REGISTER_MONITOR_CLIENT:
		account.RegisterMonitorClient(content[1:], addr)
		return nil
	}

	if monitorRes != nil {
		account.DispatchEvent(monitorRes)
	}
	return
}
