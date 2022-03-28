package account

import "encoding/binary"

type dwRequest struct {
	isDeposit bool
	amount    uint64
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

	req.amount = binary.BigEndian.Uint64(data[1:9])
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
	balance uint64
}

func (res *dwResponse) marshal() []byte {
	arr := make([]byte, 8)
	binary.BigEndian.PutUint64(arr, res.balance)
	return arr
}

type dwMonitorResponse struct {
	isDeposit bool
	accNumber uint64
	amount    uint64
}

func (res *dwMonitorResponse) marshal() []byte {
	arr := make([]byte, 17)
	if res.isDeposit {
		arr[0] = byte(1)
	} else {
		arr[0] = byte(0)
	}
	binary.BigEndian.PutUint64(arr[1:9], res.accNumber)
	binary.BigEndian.PutUint64(arr[9:17], res.amount)
	return arr
}

func DepositWithdraw(content []byte) (StatusCode, []byte, []byte) {
	req := &dwRequest{}
	req.unmarshal(content)

	status := Database.authenticate(req.accNumber, req.name, req.password)
	if status != SUCCESS {
		return status, nil, nil
	}

	if req.isDeposit {
		Database.records[req.accNumber].Balance += req.amount
	} else {
		// withdraw
		curBalance := Database.records[req.accNumber].Balance
		if curBalance < req.amount {
			return INSUFFICIENT_BALANCE, nil, nil
		}
		Database.records[req.accNumber].Balance -= req.amount
	}

	// Prepare response
	res := &dwResponse{
		balance: Database.records[req.accNumber].Balance,
	}
	monitorRes := &dwMonitorResponse{
		isDeposit: req.isDeposit,
		accNumber: req.accNumber,
		amount:    req.amount,
	}
	return SUCCESS, res.marshal(), monitorRes.marshal()
}
