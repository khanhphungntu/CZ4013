class DWRequest:
    def __init__(self, is_deposit: bool, amount: int, acc_no: int, name: str,
                 pwd: str, currency: str):
        self.is_deposit = is_deposit
        self.amount = amount
        self.acc_no = acc_no
        self.name = name
        self.pwd = pwd
        self.currency = currency

    def marshal(self):
        deposit = 1 if self.is_deposit else 0
        serialized = bytearray(deposit.to_bytes(1, 'big'))
        serialized.extend(self.amount.to_bytes(8, 'big'))
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
    def __init__(self, balance: int):
        self.balance = balance

    def __str__(self):
        return f"Your updated balance is {self.balance}"

    @classmethod
    def unmarshal(cls, data) -> str:
        balance = int.from_bytes(data[:8], 'big')
        return str(DWResponse(balance))


class DWMonitorResponse:
    def __init__(self, is_deposit: bool, acc_no: int, amount: int):
        self.is_deposit = is_deposit
        self.acc_no = acc_no
        self.amount = amount

    def __str__(self):
        action = "deposited" if self.is_deposit else "withdrawn"
        return f"Account number {self.acc_no} is {action} by amount {self.amount}"

    @classmethod
    def unmarshal(cls, data: bytes) -> str:
        deposit = int.from_bytes(data[0:1], 'big')
        is_deposit = True if deposit == 1 else False
        acc_no = int.from_bytes(data[1:9], 'big')
        amount = int.from_bytes(data[9:17], 'big')
        return str(DWMonitorResponse(
            is_deposit=is_deposit,
            acc_no=acc_no,
            amount=amount
        ))
