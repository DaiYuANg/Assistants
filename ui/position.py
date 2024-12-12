from dataclasses import dataclass
from enum import IntEnum, auto


@dataclass
class Position(IntEnum):
  West = auto()
  North = auto()
  South = auto()
  East = auto()
  Center = auto()
