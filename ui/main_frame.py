import wx


class MainFrame(wx.Frame):
    def __init__(self, *args, **kw):
        super(MainFrame, self).__init__(*args, **kw)

        self.panel = wx.Panel(self)
        self.dialog_button = wx.Button(self.panel, label="Show Dialog", pos=(10, 10))
        self.dialog_button.Bind(wx.EVT_BUTTON, self.show_dialog)

        self.form_panel = wx.Panel(self.panel)
        self.form_panel.SetBackgroundColour(wx.Colour(200, 200, 200))

        self.sizer = wx.BoxSizer(wx.HORIZONTAL)
        self.sizer.Add(self.form_panel, 1, wx.EXPAND | wx.ALL, 5)
        self.sizer.Add(self.panel, 1, wx.EXPAND | wx.ALL, 5)

        self.panel.SetSizerAndFit(self.sizer)
        self.Fit()

    def show_dialog(self):
        dialog = wx.Dialog(self, title="My Dialog", style=wx.DEFAULT_DIALOG_STYLE | wx.RESIZE_BORDER)
        dialog.ShowModal()
        dialog.Destroy()
