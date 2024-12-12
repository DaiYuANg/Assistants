import dataclasses
from enum import IntEnum, auto

@dataclasses.dataclass
class SizeType(IntEnum):
    MinimumSize = auto()
    SizeHint = auto()
