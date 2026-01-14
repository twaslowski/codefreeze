from typing import Protocol


class CoinProvider(Protocol):
    def get_coins(self) -> list[int]:
        pass
