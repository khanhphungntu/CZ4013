from enum import Enum

# Service type
ST_REGISTER_ACCOUNT = 0
ST_DELETE_ACCOUNT = 1
ST_DEPOSIT_WITHDRAW = 2
ST_REGISTER_MONITOR_CLIENT = 3
ST_TRANSFER_MONEY = 4
ST_GET_ACCOUNT_INFO = 5

# Status Code
SUCCESS = 0

ACCOUNT_NUMBER_NOT_FOUND = 1
WRONG_USER_NAME = 2
WRONG_PASSWORD = 3

INSUFFICIENT_BALANCE = 4
INVALID_RECIPIENT_ACCOUNT = 5

ERROR_MAPPING = {
    ACCOUNT_NUMBER_NOT_FOUND: "Cannot find the account number",
    WRONG_USER_NAME: "The user name is incorrect",
    WRONG_PASSWORD: "The password is incorrect",
    INSUFFICIENT_BALANCE: "Your balance is insufficient for this request",
    INVALID_RECIPIENT_ACCOUNT: "The recipient's account is incorrect",
}


class CurrencyEnum(Enum):
    USD = "USD"
    SDG = "SGD"

    @classmethod
    def list(cls):
        return [x.value for x in cls]
