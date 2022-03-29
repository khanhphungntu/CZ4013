
from numpy import int64
from request import dispatch_request
from CZ4013.client.constants import ST_DELETE_ACCOUNT

class DelReq:
  def __init__(self, accNumber: int64, name: str, pwd: str):
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


class TransferMonitorResponse:
    def __init__(self, acc_no: int):
        self.acc_no = acc_no

    def __str__(self):
        return f"Account {self.acc_no} is deleted"

    @classmethod
    def unmarshal(cls, data: bytes) -> str:
        acc_no = int.from_bytes(data[0:8], 'big')
        return str(TransferMonitorResponse(
          acc_no=acc_no
        ))

def delAccount(accNumber: int64, name: str, pwd: str):
    delReq = DelReq(accNumber, name, pwd)
    req = bytearray(ST_DELETE_ACCOUNT.to_bytes(1, 'big'))
    req.extend(delReq.marshal())
    dispatch_request(req)
    
    # statusCode = 
    # return int.from_bytes(resp, 'big')

def delUI():
  accNumber = input("Please enter the your account number: ")
  try:
    accNumber = int64(accNumber)
  except:
    print("You did not enter a valid account number")
    return
  name = input("Please enter your name: ")
  pwd = input("Please eneter your password: ")
  # delAccount(accNumber, name, pwd)
