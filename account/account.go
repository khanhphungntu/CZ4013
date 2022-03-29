package account

import (
	"encoding/binary"
	"math"
)

type Account struct {
	Name      string
	Password  string
	Currency  string
	Balance   float64
	AccNumber uint64
}

type databaseImpl struct {
	records map[uint64]*Account
}

func (d *databaseImpl) authenticate(accNumber uint64, name string, password string) StatusCode {
	if userAccount, ok := d.records[accNumber]; ok {
		// wrong accoutnt name
		if name != userAccount.Name {
			return WRONG_USER_NAME
		}
		// wrong password
		if password != userAccount.Password {
			return WRONG_PASSWORD
		}
		return SUCCESS
		//do something here
	}
	// cannot find account
	return ACCOUNT_NUMBER_NOT_FOUND
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
	binary.BigEndian.PutUint64(balanceBytes, math.Float64bits(a.Balance))

	var accountNumberBytes = make([]byte, 8)
	binary.BigEndian.PutUint64(accountNumberBytes, a.AccNumber)

	return append(serialized, accountNumberBytes...)
}
