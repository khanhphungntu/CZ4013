# CZ4013

## Primitive marshal

- `string`: 2 bytes for len, the rests are characters
- Account number, balance: uint64 (8 bytes)

## Packet format

- 2 bytes request ID
- 2 bytes payload size
- Rest: Payload

### Payload

- 1 byte for status code
- 1 byte for service type
- Rest: Content