import wx
from injector import Module

from ui.main_frame import MainFrame


class UIModule(Module):
    def configure(self, binder):
        binder.bind(wx.App, to=wx.App())
        binder.bind(wx.Frame, to=MainFrame(None))
        # binder.bind(Database, to=Database(connection_string='sqlite:///:memory:'), scope=injector.singleton)
        # binder.bind(UserRepository)
        # binder.bind(UserService)
