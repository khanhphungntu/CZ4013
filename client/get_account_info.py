import struct

import constants
import request


class GetAccInfoRequest:
    def __init__(self, acc_no: int, name: str, pwd: str):
        self.acc_no = acc_no
        self.name = name
        self.pwd = pwd

    def marshal(self):
        serialized = bytearray(self.acc_no.to_bytes(8, 'big'))

        name_size = len(self.name)
        pwd_size = len(self.pwd)

        serialized.extend(name_size.to_bytes(2, 'big'))
        serialized.extend(pwd_size.to_bytes(2, 'big'))

        serialized.extend(bytearray(self.name, 'utf-8'))
        serialized.extend(bytearray(self.pwd, 'utf-8'))
        return serialized


class GetAccInfoResponse:
    def __init__(self, accNumber: int, balance: float, name: str, currency: str):
        self.accNumber = accNumber
        self.balance = balance
        self.name = name
        self.currency = currency

    def __str__(self):
        return f"The account information: " \
               f"Account number: {self.accNumber} " \
               f"Name: {self.name} " \
               f"Balance: {self.balance} " \
               f"Currency: {self.currency}"

    @classmethod
    def unmarshal(cls, data) -> str:
        accNum = int.from_bytes(data[:8], 'big')
        balance = struct.unpack(">d", data[8:16])[0]

        nameSize = int.from_bytes(data[16:18], 'big')
        currencySize = int.from_bytes(data[18:20], 'big')
        currency_index = 20 + nameSize

        name = data[20:currency_index].decode('utf-8')
        currency = data[currency_index: currency_index + currencySize].decode('utf-8')
        return str(GetAccInfoResponse(
            accNumber=accNum,
            balance=balance,
            name=name,
            currency=currency,
        ))


def get_acc_info(acc_no: int, name: str, pwd: str):
    req = GetAccInfoRequest(
        acc_no=acc_no,
        name=name,
        pwd=pwd,
    )
    request.dispatch_request(constants.ST_GET_ACCOUNT_INFO, req.marshal())


if __name__ == '__main__':
    get_acc_info(4091, "Nhan", "1234")
