package account

import (
	"encoding/binary"
	"fmt"
	"math"
)

type dwRequest struct {
	isDeposit bool
	amount    float64
	accNumber uint64
	name      string
	password  string
	currency  string
}

func (req *dwRequest) unmarshal(data []byte) {
	if uint8(data[0]) == 0 {
		req.isDeposit = false
	} else {
		req.isDeposit = true
	}

	req.amount = math.Float64frombits(binary.BigEndian.Uint64(data[1:9]))
	req.accNumber = binary.BigEndian.Uint64(data[9:17])

	nameSize := binary.BigEndian.Uint16(data[17:19])
	pwdSize := binary.BigEndian.Uint16(data[19:21])
	currencySize := binary.BigEndian.Uint16(data[21:23])

	pwdIndex := 23 + nameSize
	req.name = string(data[23:pwdIndex])

	currencyIndex := pwdIndex + pwdSize
	req.password = string(data[pwdIndex:currencyIndex])

	req.currency = string(data[currencyIndex : currencyIndex+currencySize])
}

type dwResponse struct {
	balance  float64
	currency string
}

func (res *dwResponse) marshal() []byte {
	currencySize := len(res.currency)
	arr := make([]byte, 8+2+currencySize)
	binary.BigEndian.PutUint64(arr[0:8], math.Float64bits(res.balance))
	binary.BigEndian.PutUint16(arr[8:10], uint16(currencySize))

	copy(arr[10:10+currencySize], res.currency)
	return arr
}

func DepositWithdraw(content []byte) (StatusCode, []byte) {
	req := &dwRequest{}
	req.unmarshal(content)

	// Validation
	authCode := Database.authenticate(req.accNumber, req.name, req.password)
	if authCode != SUCCESS {
		return authCode, nil
	}

	account := Database.records[req.accNumber]
	convertedAmount := convertAmount(req.amount, req.currency, account.Currency)

	if req.isDeposit {
		account.Balance += convertedAmount
	} else {
		// withdraw
		if account.Balance < convertedAmount {
			return INSUFFICIENT_BALANCE, nil
		}
		account.Balance -= convertedAmount
	}

	// Prepare monitor update
	var action string
	if req.isDeposit {
		action = "deposited to"
	} else {
		action = "withdrawn from"
	}
	s := fmt.Sprintf("Amount %f %s is %s Account number %d",
		req.amount, req.currency, action, req.accNumber)
	clientsTrackingImpl.dispatchEvent([]byte(s))
	fmt.Println(s)

	fmt.Printf("Account number %d Updated balance: %f %s\n",
		account.AccNumber, account.Balance, account.Currency)
	// Prepare response
	res := &dwResponse{
		balance: account.Balance,
	}

	return SUCCESS, res.marshal()
}
