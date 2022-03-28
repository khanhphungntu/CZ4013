package account

import (
	"encoding/binary"
)

type Account struct {
	Name      string
	Password  string
	Currency  string
	Balance   uint64
	AccNumber uint64
}

type databaseImpl struct {
	records map[uint64]*Account
}

var Database = databaseImpl{records: make(map[uint64]*Account)}

// Marshal returns array of bytes where the first 6 bytes represent the size of Name,
// Password, Currency
func (a *Account) Marshal() []byte {
	nameSize := len(a.Name)
	pwdSize := len(a.Password)
	currencySize := len(a.Currency)

	var serialized = make([]byte, 6)
	binary.BigEndian.PutUint16(serialized, uint16(nameSize))
	binary.BigEndian.PutUint16(serialized[2:], uint16(pwdSize))
	binary.BigEndian.PutUint16(serialized[4:], uint16(currencySize))

	serialized = append(serialized, []byte(a.Name)...)
	serialized = append(serialized, []byte(a.Password)...)
	serialized = append(serialized, []byte(a.Currency)...)

	var balanceBytes = make([]byte, 8)
	binary.BigEndian.PutUint64(balanceBytes, a.Balance)

	var accountNumberBytes = make([]byte, 8)
	binary.BigEndian.PutUint64(accountNumberBytes, a.AccNumber)

	return append(serialized, accountNumberBytes...)
}
