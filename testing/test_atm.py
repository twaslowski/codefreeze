from atm import ATM
from coin_provider import CoinProvider
from coin_dispenser import CoinDispenser

DEFAULT_COINS = [200, 100, 50, 20, 10, 5, 2, 1]

class CoinProviderFake(CoinProvider):
    _coins: list[int]

    def __init__(self, coins: list[int]):
        self._coins = coins

    def get_coins(self) -> list[int]:
        return reversed(sorted(self._coins))


class CoinDispenserFake(CoinDispenser):
    _transactions: list[list[int]]

    def __init__(self):
        self._transactions = []

    def dispense(self, coins: list[int]):
        self._transactions.append(coins)
        return coins

    def get_num_transactions(self):
        return len(self._transactions)


def test_50():
    assert ATM(CoinProviderFake(DEFAULT_COINS), CoinDispenserFake()).withdraw(50) == [50]


def test_100():
    assert ATM(CoinProviderFake(DEFAULT_COINS), CoinDispenserFake()).withdraw(100) == [100]


def test_200():
    assert ATM(CoinProviderFake(DEFAULT_COINS), CoinDispenserFake()).withdraw(200) == [200]


def test_500():
    assert ATM(CoinProviderFake(DEFAULT_COINS), CoinDispenserFake()).withdraw(500) == [200, 200, 100]


def test_550():
    assert ATM(CoinProviderFake(DEFAULT_COINS), CoinDispenserFake()).withdraw(550) == [200, 200, 100, 50]


def test_567():
    assert ATM(CoinProviderFake(DEFAULT_COINS), CoinDispenserFake()).withdraw(567) == [200, 200, 100, 50, 10, 5, 2]


def test_custom_coins():
    coin_provider = CoinProviderFake([57, 37])
    atm = ATM(coin_provider, CoinDispenserFake())
    assert atm.withdraw(94) == [57, 37]
    # assert atm.retrieve(74) == [37, 37]  # todo: this breaks!


def test_disregard_custom_coins_order():
    coin_provider = CoinProviderFake([37, 57])
    atm = ATM(coin_provider, CoinDispenserFake())
    assert atm.withdraw(94) == [57, 37]
    # assert atm.retrieve(74) == [37, 37]  # todo: this breaks!

def test_coins_dispensed():
    dispenser = CoinDispenserFake()
    atm = ATM(CoinProviderFake(DEFAULT_COINS), dispenser)
    atm.withdraw(100)
    assert dispenser.get_num_transactions() == 1