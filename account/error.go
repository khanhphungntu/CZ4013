package account

type StatusCode uint8

const (
	SUCCESS StatusCode = 0

	// delete
	ACCOUNT_NUMBER_NOT_FOUND StatusCode = 1
	WRONG_USER_NAME          StatusCode = 2
	WRONG_PASSWORD           StatusCode = 3

	// deposit & withdraw
	INSUFFICIENT_BALANCE StatusCode = 4
)
