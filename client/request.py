from socket import *
import random

s = socket(type=SOCK_DGRAM)
PACKET_SIZE = 1024
TIME_OUT = 1


def dispatch_request(payload: bytes):
    req_id = random.randint(1, 2 ** 16 - 1)
    no_bytes = len(payload)
    serialized = bytearray(req_id.to_bytes(2, 'big'))
    serialized.extend(no_bytes.to_bytes(2, 'big'))
    serialized.extend(payload)

    s.settimeout(TIME_OUT)
    retry = 0

    while True:
        s.sendto(serialized, ('localhost', 8000))
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
    return resp_content


def unmarshal(data: bytes, is_monitor=False) -> str:
    # Status code
    status = int.from_bytes(data[0:1], 'big')
    # Service type
    service = int.from_bytes(data[1:2], 'big')

    # Content

