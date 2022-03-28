package main

import (
	"DS/account"
	"net"
)

type serviceType uint8

const (
	REGISTER_ACCOUNT        serviceType = 0
	DELETE_ACCOUNT          serviceType = 1
	DEPOSIT_WITHDRAW        serviceType = 2
	REGISTER_MONITOR_CLIENT serviceType = 3
)

// RouterImpl : Assume the first byte is used to identify the api
func RouterImpl(content []byte, addr *net.UDPAddr) []byte {
	status := account.SUCCESS
	var res []byte
	var monitorRes []byte
	switch serviceType(content[0]) {
	case REGISTER_ACCOUNT:
		return account.RegisterAccount(content[1:])
	case DEPOSIT_WITHDRAW:
		status, res, monitorRes = account.DepositWithdraw(content[1:])
	case DELETE_ACCOUNT:
		status, res, monitorRes = account.DeleteAccount(content[1:])
	case REGISTER_MONITOR_CLIENT:
		account.RegisterMonitorClient(content[1:], addr)
		return nil
	}

	// Check if need to dispatch monitor
	if status == account.SUCCESS && monitorRes != nil {
		// Append status to front
		monitorRes = append([]byte{byte(status)}, monitorRes...)
		account.DispatchEvent(monitorRes)
	}

	// Append status to front
	res = append([]byte{byte(status)}, res...)
	return res
}
