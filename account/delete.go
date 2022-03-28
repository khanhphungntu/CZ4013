package account

import (
	"encoding/binary"
	"fmt"
)

type deleteRequest struct {
	name      string
	accNumber uint64
	password  string
}

func (req *deleteRequest) unmarshal(data []byte) {

	// need to decide the format of the request
	nameSize := binary.BigEndian.Uint16(data[:2])
	pwdSize := binary.BigEndian.Uint16(data[2:4])
	// accountNumberSize := binary.BigEndian.Uint16(content[4:6])

	pwdIndex := 4 + nameSize
	name := string(data[4:pwdIndex])

	accountNumberIndex := pwdIndex + pwdSize
	pwd := string(data[pwdIndex:accountNumberIndex])

	accountNumber := binary.BigEndian.Uint64(data[accountNumberIndex : accountNumberIndex+8])

	*req = deleteRequest{
		name:      name,
		accNumber: accountNumber,
		password:  pwd,
	}
}

type deleteMonitorResponse struct {
	accNumber uint64
}

func (res *deleteMonitorResponse) marshal() []byte {
	// 1 for statusCode, 8 for accNumber,
	toRet := make([]byte, 8)
	binary.BigEndian.PutUint64(toRet, res.accNumber)
	return toRet
}

func (d *databaseImpl) deleteAccount(delReq *deleteRequest) StatusCode {
	authCode := d.authenticate(delReq.accNumber, delReq.name, delReq.password)
	if authCode == SUCCESS {
		delete(d.records, delReq.accNumber)
	}
	return authCode
}

func DeleteAccount(content []byte) (StatusCode, []byte, []byte) {
	// need to decide the format of the request
	deleteReq := &deleteRequest{}
	deleteReq.unmarshal(content)
	authCode := Database.authenticate(deleteReq.accNumber, deleteReq.name, deleteReq.password)
	var delMontinorResContent []byte = nil

	if authCode == SUCCESS {
		delete(Database.records, deleteReq.accNumber)
		delMonitorRes := &deleteMonitorResponse{
			accNumber: deleteReq.accNumber,
		}
		delMontinorResContent = delMonitorRes.marshal()
	}
	fmt.Println("Delete status code:", authCode)

	return authCode, nil, delMontinorResContent
}
