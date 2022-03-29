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

	nameSize := binary.BigEndian.Uint16(data[9:11])
	pwdSize := binary.BigEndian.Uint16(data[11:13])

	pwdIndex := 13 + nameSize
	req.name = string(data[13:pwdIndex])

	req.password = string(data[pwdIndex : pwdIndex+pwdSize])
}

type getAccResponse struct {
	AccNumber uint64
	Balance   float64
	Name      string
	Currency  string
}

func (res *getAccResponse) marshal() []byte {
	size := len(res.Name) + len(res.Currency) + 8 + 8
	arr := make([]byte, size)

	binary.BigEndian.PutUint64(arr[0:8], res.AccNumber)
	binary.BigEndian.PutUint64(arr[8:16], math.Float64bits(res.Balance))

	binary.BigEndian.PutUint16(arr[16:18], uint16(len(res.Name)))
	binary.BigEndian.PutUint16(arr[18:20], uint16(len(res.Currency)))

	arr = append(arr, []byte(res.Name)...)
	arr = append(arr, []byte(res.Currency)...)
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
