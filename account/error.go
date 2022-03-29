package account

type StatusCode uint8

const (
	SUCCESS StatusCode = 0

	// delete
	ACCOUNT_NUMBER_NOT_FOUND StatusCode = 1
	WRONG_USER_NAME          StatusCode = 2
	WRONG_PASSWORD           StatusCode = 3

	// deposit & withdraw
	INSUFFICIENT_BALANCE      StatusCode = 4
	WRONG_CURRENCY            StatusCode = 5
	WRONG_RECIPIENT_CURRENCY  StatusCode = 6
	INVALID_RECIPIENT_ACCOUNT StatusCode = 7
)
