# 使用 pyside6 制作 一个 聊天界面
import sys

from PySide6.QtWidgets import QApplication
from pynput import keyboard

from packages.qchat.src.chat_window import ChatWindow


def on_press(key):
    if key == keyboard.Key.esc:
        return False  # stop listener
    try:
        k = key.char  # single-char keys
    except:
        k = key.name  # other keys
    if k in ['1', '2', 'left', 'right']:  # keys of interest
        # self.keys.append(k)  # store it in global-like variable
        print('Key pressed: ' + k)
        return False  # stop listener; remove this if want more keys


if __name__ == '__main__':
    app = QApplication(sys.argv)
    window = ChatWindow()
    listener = keyboard.Listener(on_press=on_press)
    listener.start()  # start to listen on a separate thread
    # listener.join()  # remove if main thread is polling self.keys
    window.show()  # listener = BackgroundListener(window)
    sys.exit(app.exec())
