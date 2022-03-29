package account

import "encoding/binary"

type transferRequest struct {
	amount       uint64
	accNumber    uint64
	accNumberDst uint64
	name         string
	password     string
	currency     string
}

func (req *transferRequest) unmarshal(data []byte) {
	req.amount = binary.BigEndian.Uint64(data[0:8])
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
	balance uint64
}

func (res *transferResponse) marshal() []byte {
	arr := make([]byte, 8)
	binary.BigEndian.PutUint64(arr, res.balance)
	return arr
}

type transferMonitorResponse struct {
	accNumber    uint64
	accNumberDst uint64
	amount       uint64
}

func (res *transferMonitorResponse) marshal() []byte {
	arr := make([]byte, 24)
	binary.BigEndian.PutUint64(arr[:8], res.accNumber)
	binary.BigEndian.PutUint64(arr[8:16], res.accNumberDst)
	binary.BigEndian.PutUint64(arr[16:24], res.amount)
	return arr
}

func TransferMoney(content []byte) (StatusCode, []byte, []byte) {
	req := &transferRequest{}
	req.unmarshal(content)

	// Validation
	authCode := Database.authenticate(req.accNumber, req.name, req.password)
	if authCode != SUCCESS {
		return authCode, nil, nil
	}

	account := Database.records[req.accNumber]
	accountDst, ok := Database.records[req.accNumberDst]
	if !ok {
		return INVALID_RECIPIENT_ACCOUNT, nil, nil
	}

	if req.currency != account.Currency {
		return WRONG_CURRENCY, nil, nil
	}
	if req.currency != accountDst.Currency {
		return WRONG_RECIPIENT_CURRENCY, nil, nil
	}

	if account.Balance < req.amount {
		return INSUFFICIENT_BALANCE, nil, nil
	}

	account.Balance -= req.amount
	accountDst.Balance += req.amount

	// Prepare response
	res := &transferResponse{balance: account.Balance}
	monitorRes := &transferMonitorResponse{
		accNumber:    req.accNumber,
		accNumberDst: req.accNumberDst,
		amount:       req.amount,
	}
	return SUCCESS, res.marshal(), monitorRes.marshal()
}
