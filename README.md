# CX4013

CX4013 - Distributed System

## Environment

### Operating System

- Linux, MacOS

### Programming Languages

- Python 3.7+ (Client)
    - Install link: https://www.python.org/downloads/
- Golang 1.16+ (Server)
    - Install link: https://go.dev/doc/install

## Getting Started

### Server (Golang)

```bash
go run .
```

### Client (Python)

```bash
cd client
python ui.py
```

## Supported Service

### Open New Account

#### Request

| Params   | Type          |
|----------|---------------|
| Name     | ```string```  |
| Password | ```string```  |
| Currency | ```string```  |
| Balance  | ```float64``` |

#### Response

| Params         | Type         |
|----------------|--------------|
| Account Number | ```uint64``` |

### Close Existing Account

#### Request

| Params         | Type         |
|----------------|--------------|
| Account Number | ```uint64``` |
| Name           | ```string``` |
| Password       | ```string``` |

### Deposit/Withdraw Money

#### Request

| Params         | Type          |
|----------------|---------------|
| Is Deposit     | ```boolean``` |
| Account Number | ```uint64```  |
| Name           | ```string```  |
| Password       | ```string```  |
| Currency       | ```string```  |
| Balance        | ```float64``` |

#### Response

| Params   | Type          |
|----------|---------------|
| Currency | ```string```  |
| Balance  | ```float64``` |

### Monitor Updates

#### Request

| Params   | Type         |
|----------|--------------|
| Duration | ```uint64``` |

### Transfer Money

#### Request

| Params                   | Type          |
|--------------------------|---------------|
| Account Number           | ```uint64```  |
| Name                     | ```string```  |
| Recipient Account Number | ```uint64```  |
| Password                 | ```string```  |
| Currency                 | ```string```  |
| Balance                  | ```float64``` |

#### Response

| Params   | Type          |
|----------|---------------|
| Currency | ```string```  |
| Balance  | ```float64``` |

### View Account Information

#### Request

| Params         | Type         |
|----------------|--------------|
| Account Number | ```uint64``` |
| Name           | ```string``` |
| Password       | ```string``` |

#### Response

| Params         | Type          |
|----------------|---------------|
| Account Number | ```uint64```  |
| Balance        | ```float64``` |
| Name           | ```string```  |
| Currency       | ```string```  |