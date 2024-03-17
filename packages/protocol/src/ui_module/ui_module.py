import sys

from PySide6.QtWidgets import QApplication
from injector import Module, singleton, provider

from .main_window import MainWindow
from qt_material import apply_stylesheet


class UIModule(Module):
    def configure(self, binder):
        binder.bind(MainWindow)

    @singleton
    @provider
    def application_provide(self) -> QApplication:
        app = QApplication(sys.argv)
        apply_stylesheet(app, theme='dark_teal.xml')
        return app
