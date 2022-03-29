import struct

import constants
import request
from constants import CurrencyEnum
from deposit_withdraw import DWResponse


class TransferRequest:
    def __init__(self, amount: float, acc_no: int, acc_no_dst: int,
                 name: str, pwd: str, currency: str):
        assert currency in CurrencyEnum.list()

        self.amount = amount
        self.acc_no = acc_no
        self.acc_no_dst = acc_no_dst
        self.name = name
        self.pwd = pwd
        self.currency = currency

    def marshal(self):
        serialized = bytearray(struct.pack('>d', self.amount))
        serialized.extend(self.acc_no.to_bytes(8, 'big'))
        serialized.extend(self.acc_no_dst.to_bytes(8, 'big'))

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


class TransferResponse(DWResponse):
    pass


def transfer_money(amount: float, acc_no: int, acc_no_dst: int,
                   name: str, pwd: str, currency: str):
    req = TransferRequest(
        amount=amount,
        acc_no=acc_no,
        acc_no_dst=acc_no_dst,
        name=name,
        pwd=pwd,
        currency=currency
    )
    request.dispatch_request(constants.ST_TRANSFER_MONEY, req.marshal())


if __name__ == '__main__':
    transfer_money(5.1, 9410, 3551, "Nhan", "1234", "SGD")
