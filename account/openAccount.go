package account

import (
	"encoding/binary"
	"fmt"
	"math/rand"
)

func (d *databaseImpl) registerAccount(name string, pwd string, currency string, balance uint64) uint64 {
	accountNumber := uint64(rand.Int63n(10000))
	d.records[accountNumber] = &Account{
		Name:      name,
		Password:  pwd,
		Currency:  currency,
		Balance:   balance,
		AccNumber: accountNumber,
	}

	return accountNumber
}

func RegisterAccount(content []byte) []byte {
	nameSize := binary.BigEndian.Uint16(content[:2])
	pwdSize := binary.BigEndian.Uint16(content[2:4])
	currencySize := binary.BigEndian.Uint16(content[4:6])

	pwdIndex := 6 + nameSize
	name := string(content[6:pwdIndex])

	currencyIndex := pwdIndex + pwdSize
	pwd := string(content[pwdIndex:currencyIndex])

	balanceIndex := currencyIndex + currencySize
	currency := string(content[currencyIndex:balanceIndex])

	balance := binary.BigEndian.Uint64(content[balanceIndex : balanceIndex+8])
	accountNumber := Database.registerAccount(name, pwd, currency, balance)

	fmt.Println("New account registered with account number:", accountNumber)
	dispatchOpenAccountEvent(accountNumber)

	resp := make([]byte, 8)
	binary.BigEndian.PutUint64(resp, accountNumber)
	return resp
}

func dispatchOpenAccountEvent(accountNumber uint64) {
	account := Database.records[accountNumber]
	s := fmt.Sprintf("AccNumber: %d, Name: %s, Password: %s, Currency: %s, Balance: %d",
		account.AccNumber, account.Name, account.Password, account.Currency, account.Balance)

	clientsTrackingImpl.dispatchEvent([]byte(s))
}
