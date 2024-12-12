import sys

from PySide6.QtWidgets import QApplication, QWidget, QVBoxLayout
from dependency_injector import containers, providers

from ui.border_layout import BorderLayout
from ui.message_pane import MessagePane
from ui.window import Window


class UIModule(containers.DeclarativeContainer):
  app = providers.Singleton(QApplication, sys.argv)
  # message_pane = providers.Singleton(MessagePane)
  # border_layout = providers.Singleton(BorderLayout)
  window = providers.Factory(Window)
