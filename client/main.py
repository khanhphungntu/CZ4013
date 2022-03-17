from request import dispatch_request

REGISTER_ACCOUNT_REQUEST_ID = 0


class Account:
    def __init__(self, name: str, pwd: str, currency: str, balance: int):
        self.name = name
        self.pwd = pwd
        self.currency = currency
        self.balance = balance

    def marshal(self) -> bytes:
        name_size = len(self.name)
        pwd_size = len(self.pwd)
        currency_size = len(self.currency)

        serialized = bytearray(name_size.to_bytes(2, 'big'))
        serialized.extend(pwd_size.to_bytes(2, 'big'))
        serialized.extend(currency_size.to_bytes(2, 'big'))
        serialized.extend(bytearray(self.name, 'utf-8'))
        serialized.extend(bytearray(self.pwd, 'utf-8'))
        serialized.extend(bytearray(self.currency, 'utf-8'))

        serialized.extend(self.balance.to_bytes(8, 'big'))
        return serialized


def register_account(name: str, pwd: str, currency: str, balance: int):
    acc = Account(name, pwd, currency, balance)
    req = bytearray(REGISTER_ACCOUNT_REQUEST_ID.to_bytes(1, 'big'))
    req.extend(acc.marshal())
    resp = dispatch_request(req)

    return int.from_bytes(resp, 'big')


print(register_account(
    name="Khanh",
    pwd="123",
    currency="SGD",
    balance=1000
))
