import random
from socket import *

import constants
import deposit_withdraw
import get_account_info
import open_account
import transfer_money

s = socket(type=SOCK_DGRAM)
PACKET_SIZE = 1024
TIME_OUT = 1


def dispatch_request(service_type: int, payload: bytes):
    req_id = random.randint(1, 2 ** 16 - 1)
    no_bytes = len(payload)
    serialized = bytearray(req_id.to_bytes(2, 'big'))
    serialized.extend(no_bytes.to_bytes(2, 'big'))

    # Service type then payload
    serialized.extend(bytearray(service_type.to_bytes(1, 'big')))
    serialized.extend(payload)

    s.settimeout(TIME_OUT)
    retry = 0

    while True:
        s.sendto(serialized, (constants.IP, constants.PORT))
        try:
            data, addr = s.recvfrom(PACKET_SIZE)

            if int.from_bytes(data[:2], 'big') != req_id:
                raise Exception("Invalid request response received")
            break
        except Exception as e:
            print(e)
            print('Request time out, retrying: {}'.format(retry))

        retry += 1

    resp_size = int.from_bytes(data[2:4], 'big')
    resp_content = data[4: 4 + resp_size]
    # Print response from server directly
    print(unmarshal(resp_content))


def unmarshal(data: bytes) -> str:
    # Status code
    status = int.from_bytes(data[0:1], 'big')
    if status != constants.SUCCESS:
        return constants.ERROR_MAPPING[status]

    # Service type
    service = int.from_bytes(data[1:2], 'big')
    if service == constants.ST_REGISTER_ACCOUNT:
        klass = open_account.OpenAccountResponse
    elif service == constants.ST_DELETE_ACCOUNT:
        return "Account is deleted successfully"
    elif service == constants.ST_DEPOSIT_WITHDRAW:
        klass = deposit_withdraw.DWResponse
    elif service == constants.ST_TRANSFER_MONEY:
        klass = transfer_money.TransferResponse
    elif service == constants.ST_GET_ACCOUNT_INFO:
        klass = get_account_info.GetAccInfoResponse
    else:
        raise NotImplementedError

    return klass.unmarshal(data[2:])
