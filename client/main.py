from client.constants import ST_REGISTER_ACCOUNT
from client.open_account import OpenAccountRequest
from request import dispatch_request


def register_account(name: str, pwd: str, currency: str, balance: int):
    acc = OpenAccountRequest(name, pwd, currency, balance)
    req = bytearray(ST_REGISTER_ACCOUNT.to_bytes(1, 'big'))
    req.extend(acc.marshal())
    dispatch_request(req)


if __name__ == '__main__':
    for i in range(20):
        print(register_account(
            name="Khanh",
            pwd="123",
            currency="SGD",
            balance=1000
        ))
