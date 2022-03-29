package account

import (
	"encoding/binary"
	"fmt"
	"math"
)

type transferRequest struct {
	amount       float64
	accNumber    uint64
	accNumberDst uint64
	name         string
	password     string
	currency     string
}

func (req *transferRequest) unmarshal(data []byte) {
	req.amount = math.Float64frombits(binary.BigEndian.Uint64(data[0:8]))
	req.accNumber = binary.BigEndian.Uint64(data[8:16])
	req.accNumberDst = binary.BigEndian.Uint64(data[16:24])

	nameSize := binary.BigEndian.Uint16(data[24:26])
	pwdSize := binary.BigEndian.Uint16(data[26:28])
	currencySize := binary.BigEndian.Uint16(data[28:30])

	pwdIndex := 30 + nameSize
	req.name = string(data[30:pwdIndex])

	currencyIndex := pwdIndex + pwdSize
	req.password = string(data[pwdIndex:currencyIndex])

	req.currency = string(data[currencyIndex : currencyIndex+currencySize])
}

type transferResponse struct {
	balance  float64
	currency string
}

func (res *transferResponse) marshal() []byte {
	currencySize := len(res.currency)
	arr := make([]byte, 8+2+currencySize)
	binary.BigEndian.PutUint64(arr[0:8], math.Float64bits(res.balance))
	binary.BigEndian.PutUint16(arr[8:10], uint16(currencySize))

	copy(arr[10:], res.currency)
	return arr
}

func TransferMoney(content []byte) (StatusCode, []byte) {
	req := &transferRequest{}
	req.unmarshal(content)

	// Validation
	authCode := Database.authenticate(req.accNumber, req.name, req.password)
	if authCode != SUCCESS {
		return authCode, nil
	}

	account := Database.records[req.accNumber]
	accountDst, ok := Database.records[req.accNumberDst]
	if !ok {
		return INVALID_RECIPIENT_ACCOUNT, nil
	}

	convertedAmount := convertAmount(req.amount, req.currency, account.Currency)
	convertedAmountDst := convertAmount(req.amount, req.currency, accountDst.Currency)

	if account.Balance < convertedAmount {
		return INSUFFICIENT_BALANCE, nil
	}

	account.Balance -= convertedAmount
	accountDst.Balance += convertedAmountDst

	// Prepare monitor dispatch
	s := fmt.Sprintf("Amount %f %s is transferred from Account number %d to Account number %d",
		req.amount, req.currency, req.accNumber, req.accNumberDst)
	clientsTrackingImpl.dispatchEvent([]byte(s))
	fmt.Println(s)

	// Prepare response
	res := &transferResponse{balance: account.Balance, currency: account.Currency}
	return SUCCESS, res.marshal()
}
