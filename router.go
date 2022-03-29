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
	TRANSFER_MONEY          serviceType = 4
)

// RouterImpl : Assume the first byte is used to identify the api
func RouterImpl(content []byte, addr *net.UDPAddr) []byte {
	status := account.SUCCESS
	var res []byte

	service := serviceType(content[0])
	switch service {
	case REGISTER_ACCOUNT:
		status, res = account.RegisterAccount(content[1:])
	case DEPOSIT_WITHDRAW:
		status, res = account.DepositWithdraw(content[1:])
	case DELETE_ACCOUNT:
		status, res = account.DeleteAccount(content[1:])
	case TRANSFER_MONEY:
		status, res = account.TransferMoney(content[1:])
	case REGISTER_MONITOR_CLIENT:
		account.RegisterMonitorClient(content[1:], addr)
		return nil
	}

	// Append status to front
	res = append([]byte{byte(status), byte(service)}, res...)
	return res
}
