package account

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
)

type openAccountRequest struct {
	name     string
	password string
	currency string
	balance  float64
}

func (req *openAccountRequest) unmarshal(content []byte) {
	nameSize := binary.BigEndian.Uint16(content[:2])
	pwdSize := binary.BigEndian.Uint16(content[2:4])
	currencySize := binary.BigEndian.Uint16(content[4:6])

	pwdIndex := 6 + nameSize
	req.name = string(content[6:pwdIndex])

	currencyIndex := pwdIndex + pwdSize
	req.password = string(content[pwdIndex:currencyIndex])

	balanceIndex := currencyIndex + currencySize
	req.currency = string(content[currencyIndex:balanceIndex])

	req.balance = math.Float64frombits(binary.BigEndian.Uint64(content[balanceIndex : balanceIndex+8]))
}

func (d *databaseImpl) registerAccount(name string, pwd string, currency string, balance float64) uint64 {
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

func RegisterAccount(content []byte) (StatusCode, []byte) {
	req := &openAccountRequest{}
	req.unmarshal(content)
	accountNumber := Database.registerAccount(req.name, req.password, req.currency, req.balance)

	fmt.Println("New account registered with account number:", accountNumber)
	dispatchOpenAccountEvent(accountNumber)

	resp := make([]byte, 8)
	binary.BigEndian.PutUint64(resp, accountNumber)
	return SUCCESS, resp
}

func dispatchOpenAccountEvent(accountNumber uint64) {
	account := Database.records[accountNumber]
	s := fmt.Sprintf("Account number is created with info "+
		"AccNumber: %d, Name: %s, Password: %s, Currency: %s, Balance: %f",
		account.AccNumber, account.Name, account.Password, account.Currency, account.Balance)

	clientsTrackingImpl.dispatchEvent([]byte(s))
}
