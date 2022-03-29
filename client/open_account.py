import struct


class OpenAccountRequest:
    def __init__(self, name: str, pwd: str, currency: str, balance: float):
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

        serialized.extend(struct.pack('>d', self.balance))
        return serialized


class OpenAccountResponse:
    def __init__(self, acc_no: int):
        self.acc_no = acc_no

    def __str__(self):
        return f"Your account is created successfully with Number: {self.acc_no}"

    @classmethod
    def unmarshal(cls, data) -> str:
        acc_no = int.from_bytes(data, 'big')
        return str(OpenAccountResponse(acc_no))
