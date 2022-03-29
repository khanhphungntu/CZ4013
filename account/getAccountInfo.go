package account

import (
	"encoding/binary"
	"math"
)

type getAccRequest struct {
	accNumber uint64
	name      string
	password  string
}

func (req *getAccRequest) unmarshal(data []byte) {
	req.accNumber = binary.BigEndian.Uint64(data[0:8])

	nameSize := binary.BigEndian.Uint16(data[8:10])
	pwdSize := binary.BigEndian.Uint16(data[10:12])

	pwdIndex := 12 + nameSize
	req.name = string(data[12:pwdIndex])

	req.password = string(data[pwdIndex : pwdIndex+pwdSize])
}

type getAccResponse struct {
	AccNumber uint64
	Balance   float64
	Name      string
	Currency  string
}

func (res *getAccResponse) marshal() []byte {
	nameSize := len(res.Name)
	currencySize := len(res.Currency)
	size := nameSize + currencySize + 8 + 8 + 2 + 2
	arr := make([]byte, size)

	binary.BigEndian.PutUint64(arr[0:8], res.AccNumber)
	binary.BigEndian.PutUint64(arr[8:16], math.Float64bits(res.Balance))

	binary.BigEndian.PutUint16(arr[16:18], uint16(len(res.Name)))
	binary.BigEndian.PutUint16(arr[18:20], uint16(len(res.Currency)))

	currencyIndex := 20 + nameSize
	copy(arr[20:currencyIndex], res.Name)
	copy(arr[currencyIndex:currencyIndex+currencySize], res.Currency)
	return arr
}

func GetAccInfo(content []byte) (StatusCode, []byte) {
	req := &getAccRequest{}
	req.unmarshal(content)

	status := Database.authenticate(req.accNumber, req.name, req.password)
	if status != SUCCESS {
		return status, nil
	}

	// Prepare response
	account := Database.records[req.accNumber]
	res := &getAccResponse{
		AccNumber: account.AccNumber,
		Balance:   account.Balance,
		Name:      account.Name,
		Currency:  account.Currency,
	}
	return SUCCESS, res.marshal()
}
