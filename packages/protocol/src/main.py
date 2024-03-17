from PySide6.QtWidgets import QApplication
from injector import Injector
from loguru import logger

from service.tcp_server_factory import TCPServerFactory
from ui_module.main_window import MainWindow
from ui_module.ui_module import UIModule

injector = Injector([UIModule()])
TCPServerFactory()
if __name__ == '__main__':
    logger.debug("That's it, beautiful and simple logging!")
    app = injector.get(QApplication)
    window = injector.get(MainWindow)
    window.show()
    window.center_window()
    app.exec()
