from PySide6.QtWidgets import QWidget, QTextBrowser, QLabel, QFrame, QPushButton

from ui.border_layout import BorderLayout
from ui.left_section import LeftSection
from ui.message_pane import MessagePane
from ui.position import Position


class Window(QWidget):
    def __init__(self):
        super().__init__()
        self.resize(800, 600)
        self.central_widget = MessagePane()
        # self.central_widget.setPlainText("Central widget")

        border_layout = BorderLayout()
        border_layout.addWidget(self.central_widget, Position.Center)

        # label_north = self.create_label("North")
        # border_layout.addWidget(label_north, Position.North)

        left = LeftSection()
        border_layout.addWidget(left, Position.West)

        label_east1 = self.create_label("East 1")
        border_layout.addWidget(label_east1, Position.East)

        label_south = self.create_label("South")
        border_layout.addWidget(label_south, Position.South)

        self.setLayout(border_layout)
        self.setWindowTitle("Border Layout")
        self.center()

    @staticmethod
    def create_label(text: str):
        label = QLabel(text)
        label.setFrameStyle(QFrame.Box | QFrame.Raised)
        return label

    def center(self):
        qr = self.frameGeometry()
        cp = self.screen().availableGeometry().center()

        qr.moveCenter(cp)
        self.move(qr.topLeft())
