import threading
from PySide6.QtCore import QTimer
from pynput import keyboard


class BackgroundListener:
    def __init__(self, window):
        self.window = window
        self.timer = QTimer()
        self.timer.timeout.connect(self.check_keyboard)
        self.timer.start(100)  # 每100毫秒检查一次键盘事件

    def check_keyboard(self):
        if keyboard.is_pressed("a"):  # 在这里更改为你想要的按键
            print("Key 'a' pressed")
            self.show_window()

    def show_window(self):
        self.window.show()
