import request
from constants import ST_DELETE_ACCOUNT


class DeleteRequest:
    def __init__(self, accNumber: int, name: str, pwd: str):
        self.name = name
        self.pwd = pwd
        self.accNumber = accNumber

    def marshal(self) -> bytes:
        name_size = len(self.name)
        pwd_size = len(self.pwd)

        serialized = bytearray(name_size.to_bytes(2, 'big'))
        serialized.extend(pwd_size.to_bytes(2, 'big'))
        serialized.extend(bytearray(self.name, 'utf-8'))
        serialized.extend(bytearray(self.pwd, 'utf-8'))

        serialized.extend(self.accNumber.to_bytes(8, 'big'))
        return serialized


def delete_account(accNumber: int, name: str, pwd: str):
    delReq = DeleteRequest(accNumber, name, pwd)
    request.dispatch_request(ST_DELETE_ACCOUNT, delReq.marshal())
