from enum import Enum

# Service type
ST_REGISTER_ACCOUNT = 0
ST_DELETE_ACCOUNT = 1
ST_DEPOSIT_WITHDRAW = 2
ST_REGISTER_MONITOR_CLIENT = 3
ST_TRANSFER_MONEY = 4

# Status Code
SUCCESS = 0

ACCOUNT_NUMBER_NOT_FOUND = 1
WRONG_USER_NAME = 2
WRONG_PASSWORD = 3

INSUFFICIENT_BALANCE = 4
WRONG_CURRENCY = 5
WRONG_RECIPIENT_CURRENCY = 6
INVALID_RECIPIENT_ACCOUNT = 7

ERROR_MAPPING = {
    ACCOUNT_NUMBER_NOT_FOUND: "",
    WRONG_USER_NAME: "",
    WRONG_PASSWORD: "",
    INSUFFICIENT_BALANCE: "",
    WRONG_CURRENCY: "",
    WRONG_RECIPIENT_CURRENCY: "",
    INVALID_RECIPIENT_ACCOUNT: "",
}


class CurrencyEnum(Enum):
    USD = "USD"
    SDG = "SGD"

    @classmethod
    def list(cls):
        return [x.value for x in cls]
