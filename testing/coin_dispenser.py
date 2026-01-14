from typing import Protocol

class CoinDispenser(Protocol):
    def dispense(self, coins: list[int]):
        pass

    def total_number_transactions(self) -> int:
        pass