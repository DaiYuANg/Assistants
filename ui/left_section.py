import asyncio

from PySide6.QtCore import Slot
from PySide6.QtWidgets import QVBoxLayout, QPushButton, QWidget, QHBoxLayout
from loguru import logger

from service.tcp_client import tcp_client, TcpClientWorker
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

      # 启动一个新的后台线程
      worker = TcpClientWorker(client_param.host, int(client_param.port))
      worker.status_signal.connect(self.update_status)
      worker.start()

  @staticmethod
  def update_status(status: str):
    logger.info(status)
