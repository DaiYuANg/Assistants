import logging
import sys

import wx
from injector import Injector

from ui.main_panel import ExamplePanel
from ui.ui_module import UIModule
injector = Injector(modules=[UIModule()])
if __name__ == "__main__":
    logging.info("test")
    app = injector.get(wx.App)
    frame = injector.get(wx.Frame)
    panel = ExamplePanel(frame)
    frame.Center()
    frame.Show()
    frame.SetFocus()
    app.MainLoop()
