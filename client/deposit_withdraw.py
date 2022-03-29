import struct

import constants
import request
from constants import CurrencyEnum


class DWRequest:
    def __init__(self, is_deposit: bool, amount: float, acc_no: int, name: str,
                 pwd: str, currency: str):
        assert currency in CurrencyEnum.list()

        self.is_deposit = is_deposit
        self.amount = amount
        self.acc_no = acc_no
        self.name = name
        self.pwd = pwd
        self.currency = currency

    def marshal(self):
        deposit = 1 if self.is_deposit else 0
        serialized = bytearray(deposit.to_bytes(1, 'big'))
        serialized.extend(struct.pack('>d', self.amount))
        serialized.extend(self.acc_no.to_bytes(8, 'big'))

        name_size = len(self.name)
        pwd_size = len(self.pwd)
        currency_size = len(self.currency)

        serialized.extend(name_size.to_bytes(2, 'big'))
        serialized.extend(pwd_size.to_bytes(2, 'big'))
        serialized.extend(currency_size.to_bytes(2, 'big'))

        serialized.extend(bytearray(self.name, 'utf-8'))
        serialized.extend(bytearray(self.pwd, 'utf-8'))
        serialized.extend(bytearray(self.currency, 'utf-8'))
        return serialized


class DWResponse:
    def __init__(self, balance: float):
        self.balance = balance

    def __str__(self):
        return f"Your updated balance is {self.balance}"

    @classmethod
    def unmarshal(cls, data) -> str:
        balance = struct.unpack('>d', data[:8])[0]
        return str(DWResponse(balance))


def deposit_withdraw(is_deposit: bool, amount: float, acc_no: int, name: str,
                     pwd: str, currency: str):
    req = DWRequest(
        is_deposit=is_deposit,
        amount=amount,
        acc_no=acc_no,
        name=name,
        pwd=pwd,
        currency=currency,
    )
    request.dispatch_request(constants.ST_DEPOSIT_WITHDRAW, req.marshal())


if __name__ == '__main__':
    deposit_withdraw(True, 10, 3551, "Nhan", "1234", "SGD")
