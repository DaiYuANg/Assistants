import sys

from PySide6.QtWidgets import QApplication, QWidget
from dependency_injector import containers, providers
from pony.orm import set_sql_debug

from ui.window import Window


class Container(containers.DeclarativeContainer):
    app = providers.Singleton(QApplication, sys.argv)
    set_sql_debug(True)
    window = providers.Factory(Window)
