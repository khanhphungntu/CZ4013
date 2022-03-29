from client.constants import ST_REGISTER_ACCOUNT
from client.open_account import OpenAccountRequest
from request import dispatch_request


def register_account(name: str, pwd: str, currency: str, balance: int):
    acc = OpenAccountRequest(name, pwd, currency, balance)
    req = bytearray(ST_REGISTER_ACCOUNT.to_bytes(1, 'big'))
    req.extend(acc.marshal())
    resp = dispatch_request(req)

    return int.from_bytes(resp, 'big')


print(register_account(
    name="Khanh",
    pwd="123",
    currency="SGD",
    balance=1000
))
