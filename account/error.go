package account

type statusCode uint8

var (
	SUCCESS statusCode = 0

	// delete
	ACCOUNT_NUMBER_NOT_FOUND statusCode = 1
	WRONG_USER_NAME          statusCode = 2
	WRONG_PASSWORD           statusCode = 3
)
