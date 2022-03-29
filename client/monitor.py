import random
import time
from socket import *

REGISTER_CALLBACK_REQUEST_ID = 3

s = socket(type=SOCK_DGRAM)
PACKET_SIZE = 1024


def register_callback(interval: int):
    payload = bytearray(REGISTER_CALLBACK_REQUEST_ID.to_bytes(1, 'big'))
    payload.extend(interval.to_bytes(8, 'big'))

    req_id = random.randint(1, 2 ** 16 - 1)
    no_bytes = len(payload)
    serialized = bytearray(req_id.to_bytes(2, 'big'))
    serialized.extend(no_bytes.to_bytes(2, 'big'))
    serialized.extend(payload)

    s.settimeout(interval + 10)
    s.sendto(serialized, ('localhost', 8000))
    expire = time.time() + interval + 10
    while time.time() < expire:
        try:
            data, addr = s.recvfrom(PACKET_SIZE)
            if data:
                resp_size = int.from_bytes(data[:2], 'big')
                resp_content = data[2: 2 + resp_size]
                print(resp_content)
        except timeout:
            break

        time.sleep(1)


register_callback(30)
