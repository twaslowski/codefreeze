from coin_provider import CoinProvider
from coin_dispenser import CoinDispenser


class ATM:
    coin_provider: CoinProvider
    coin_dispenser: CoinDispenser

    def __init__(self, coin_provider: CoinProvider, coin_dispenser: CoinDispenser) -> None:
        self.coin_provider = coin_provider
        self.coin_dispenser = coin_dispenser

    def withdraw(self, amount: int) -> list[int]:
        result = []
        remaining_amount = amount
        for coin in self.coin_provider.get_coins():
            while remaining_amount >= coin:
                result.append(coin)
                remaining_amount -= coin
        self.coin_dispenser.dispense(result)
        return result


if __name__ == "__main__":
    atm = ATM(CoinProvider(), CoinDispenser())
    print(atm.withdraw(567))
