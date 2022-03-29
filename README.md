# CX4013

CX4013 - Distributed System

## Environment

### Operating System

- Linux, MacOS

### Programming Languages

- Python (Client)
- Golang (Server)

## Getting Started

### Client (Java)

```
cd client
javac -d . -cp .:lib/* *.java
java -cp .:lib/* client.UDPClient -h <HOST NAME> -p <PORT> [-al] [-am] [-fr <FAILURE RATE>] [-to <TIMEOUT>] [-mt <MAX TIMEOUT COUNT>] [-v]
```

#### Note:

```bash
Options:
-al,--atleast             Enable at least once invocation semantic
-am,--atmost              Enable at most once invocation semantic
-fr,--failurerate <arg>   Set failure rate (float)
-h,--host <arg>           Server host
-mt,--maxtimeout <arg>    Set timeout max count
-p,--port <arg>           Server port
-to,--timeout <arg>       Set timeout in millisecond
-v,--verbose              Enable verbose print for debugging
```

### Server (C++)

```
cd server
g++ -o server -std=c++11 main.cpp udp_server.cpp utils.cpp Handler.cpp AccountManager.cpp Account.cpp Admin.cpp
./server <PORT> <MODE> <FAULT> <LIMIT>
```

#### Note:

```<MODE>``` is the invocation semantic. Possible values:

- 0: no ack
- 1: at-least-once
- 2: at-most-once

```<FAULT>``` is the probability that server fails to reply

```<LIMIT>``` is the limit of retries

## Supported Service

### Open New Account

#### Request

| Params   | Type         |
| -------- | ------------ |
| Name     | ```string``` |
| Password | ```string``` |
| Currency | ```string``` |
| Balance  | ```float```  |

#### Response

| Params         | Type       |
| -------------- | ---------  |
| Account Number |```uint64```|

### Close Existing Account

#### Request

| Params         | Type         |
| -------------- | ------------ |
| Account Number | ```uint64``` |
| Name           | ```string``` |
| Password       | ```string``` |

### Deposit/Withdraw Money

#### Request

| Params         | Type         |
| -------------- | ------------ |
| Is Deposit     | ```boolean```|
| Account Number | ```uint64``` |
| Name           | ```string``` |
| Password       | ```string``` |
| Currency       | ```string``` |
| Balance        | ```float```  |

#### Response

| Params   | Type        |
| -------- | ----------- |
| Currency | ```string```|
| Balance  | ```float``` |

### Monitor Updates

#### Request

| Params   | Type      |
| -------- | --------- |
| Duration | ```uint64```|

### Transfer Money

#### Request

| Params                   | Type         |
| ------------------------ | ------------ |
| Account Number           | ```uint64``` |
| Name                     | ```string``` |
| Recipient Account Number | ```uint64``` |
| Password                 | ```string``` |
| Currency                 | ```string``` |
| Balance                  | ```float```  |

#### Response

| Params   | Type        |
| -------- | ----------- |
| Currency | ```string```|
| Balance  | ```float``` |

### Change Password

#### Request

| Params         | Type         |
| -------------- | ------------ |
| Name           | ```String``` |
| Account Number | ```int```    |
| Old Password   | ```String``` |
| New Password   | ```String``` |

#### Response

| Params | Type          |
| ------ | ------------- |
| ACK    | ```boolean``` |