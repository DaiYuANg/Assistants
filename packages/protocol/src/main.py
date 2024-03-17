from PySide6.QtWidgets import QApplication
from injector import Injector
from loguru import logger

from local_store.store_module import StoreModule
from local_store.model import Project
from service.tcp_server_factory import TCPServerFactory
from ui_module.main_window import MainWindow
from ui_module.ui_module import UIModule
from peewee import *
injector = Injector([UIModule(), StoreModule()])
if __name__ == '__main__':
    db = injector.get(SqliteDatabase)
    db.connect()
    db.create_tables([Project])
    logger.debug("That's it, beautiful and simple logging!")
    app = injector.get(QApplication)
    window = injector.get(MainWindow)
    window.show()
    window.center_window()
    app.exec()
