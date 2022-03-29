from constants import ST_DELETE_ACCOUNT
from request import dispatch_request


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


def delAccount(accNumber: int, name: str, pwd: str):
    delReq = DeleteRequest(accNumber, name, pwd)
    req = bytearray(ST_DELETE_ACCOUNT.to_bytes(1, 'big'))
    req.extend(delReq.marshal())
    dispatch_request(req)

    # statusCode = 
    # return int.from_bytes(resp, 'big')


def delUI():
    accNumber = input("Please enter the your account number: ")
    try:
        accNumber = int(accNumber)
    except:
        print("You did not enter a valid account number")
        return
    name = input("Please enter your name: ")
    pwd = input("Please eneter your password: ")
    # delAccount(accNumber, name, pwd)
