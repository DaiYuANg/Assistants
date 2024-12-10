from dataclasses import dataclass

from PySide6.QtWidgets import QLayoutItem

from ui.position import Position


@dataclass
class ItemWrapper:
    item: QLayoutItem
    position: Position
