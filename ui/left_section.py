import asyncio

from PySide6.QtCore import Slot
from PySide6.QtWidgets import QVBoxLayout, QPushButton, QWidget, QHBoxLayout
from loguru import logger

from service.tcp_client import tcp_client
from ui.connect_dialog import ConnectDialog


class LeftSection(QWidget):
    def __init__(self):
        super().__init__()
        layout = QVBoxLayout()
        btn = QPushButton("Connect")
        btns = QPushButton("Server")
        btn.clicked.connect(self.on_button_click)
        layout.addWidget(btn)
        layout.addWidget(btns)
        self.setLayout(layout)

    @Slot()
    def on_button_click(self):
        dialog = ConnectDialog(self)
        result = dialog.exec_()
        logger.debug(result)
        if result == 1:
            client_param = dialog.get_data()
            logger.info(client_param)
            asyncio.run(tcp_client(host=client_param.host, port=int(client_param.port)))
