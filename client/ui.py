import constants
from constants import CurrencyEnum
from delete import delete_account
from deposit_withdraw import deposit_withdraw
from get_account_info import get_acc_info
from monitor import register_callback
from open_account import register_account
from transfer_money import transfer_money


def getInputWithType(type, msg: str):
    while True:
        try:
            toRet = type(input(msg))
            return toRet
        except:
            print("You did not enter a valid input")


def getBoolean(msg: str):
    while True:
        input = getInputWithType(str, msg + " Y/N: ")
        if (input == "Y"):
            return True
        if (input == "N"):
            return False
        print("PLease enter Y or N")


def getMoneyAmount(msg: str):
    while True:
        money_amount = getInputWithType(float, msg)
        if money_amount > 0:
            return money_amount
        else:
            print("The amount of money must be higher than 0")


def getCurrency():
    currencyList = CurrencyEnum.list()
    print("Please enter a number for currency ")
    for i in range(0, len(currencyList)):
        print(f"{i}: {currencyList[i]}")
    while True:
        try:
            toRet = int(input("Please enter the currency: "))
            if toRet >= len(currencyList) or toRet < 0:
                print(" Your input is not a valid currency")
                continue
            return currencyList[toRet]
        except:
            print("You did not enter a valid input")


def getBasicInfo():
    accNum = getInputWithType(int, "Please enter your account number: ")
    name = getInputWithType(str, "Please enter your user name: ")
    pwd = getInputWithType(str, "Please enter your password: ")
    return (accNum, name, pwd)


def UI():
    ip = getInputWithType(str, "Please enter the ip: ")
    port = getInputWithType(int, "Please enter the port: ")
    if (ip != "LOCAL"):
        constants.IP = ip.strip()
        constants.PORT = port

    userInput = ""
    while (userInput != "Q"):
        print("\n")
        print("Welcome to ZZZ bank")
        print("Press A to add new account")
        print("Press R to remove your account")
        print("Press D to deposit money")
        print("Press W to withdraw money")
        print("Press T to transfer money")
        print("Press M to start monitor")
        print("Press V to view account info")
        print("Press Q to quit ")
        print("\n")
        userInput = input("Please enter your input: ")
        if userInput == "A":
            print("Welcome to register account")
            name = getInputWithType(str, "Please enter your user name: ")
            pwd = getInputWithType(str, "Please enter your password: ")
            currency = getCurrency()
            balance = getMoneyAmount("Please enter your balance: ")
            print("\n")
            register_account(name, pwd, currency, balance)
        if (userInput == "R"):
            print("Please enter your information to delete account")
            (accNum, name, pwd) = getBasicInfo()
            print("\n")
            delete_account(accNum, name, pwd)
        if (userInput == "D"):
            print("Please enter the information to deposit money")
            (accNum, name, pwd) = getBasicInfo()
            amount = getMoneyAmount("Please enter the amount of money: ")
            currency = getCurrency()
            print("\n")
            deposit_withdraw(
                True,
                amount,
                accNum,
                name,
                pwd,
                currency,
            )
        if (userInput == "W"):
            print("Please enter the information to withdrawn money")
            (accNum, name, pwd) = getBasicInfo()
            amount = getMoneyAmount("Please enter the amount of money: ")
            currency = getCurrency()
            print("\n")
            deposit_withdraw(
                False,
                amount,
                accNum,
                name,
                pwd,
                currency,
            )
        if (userInput == "T"):
            print("Please enter information for transfer money")
            (accNum, name, pwd) = getBasicInfo()
            acc_no_dst = getInputWithType(int, "Please enter destination account number: ")
            amount = getMoneyAmount("Please enter the amount of money: ")
            currency = getCurrency()
            print("\n")
            transfer_money(
                amount=amount,
                acc_no=accNum,
                acc_no_dst=acc_no_dst,
                name=name,
                pwd=pwd,
                currency=currency,
            )
        if (userInput == "M"):
            monitor_time = getInputWithType(int, "Please enter the interval you want to monitor in second: ")
            register_callback(monitor_time)
        if (userInput == "V"):
            print("Please enter information to view account")
            (accNum, name, pwd) = getBasicInfo()
            print("\n")
            get_acc_info(
                accNum,
                name,
                pwd,
            )


UI()
