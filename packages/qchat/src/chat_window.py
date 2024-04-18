from PySide6.QtGui import QTextCursor
from PySide6.QtWidgets import QMainWindow, QVBoxLayout, QTextEdit, QLineEdit, QPushButton, QWidget, \
    QHBoxLayout


class ChatWindow(QMainWindow):
    def __init__(self):
        super().__init__()

        self.setWindowTitle("Simple Chat Interface")

        # 创建主布局
        main_layout = QVBoxLayout()

        # 消息显示区域
        self.message_display = QTextEdit(readOnly=True)
        self.message_display.setLineWrapMode(QTextEdit.NoWrap)
        main_layout.addWidget(self.message_display)

        # 输入框与发送按钮布局
        input_layout = QHBoxLayout()

        # 文本输入框
        self.input_field = QLineEdit()
        input_layout.addWidget(self.input_field)

        # 发送按钮
        send_button = QPushButton("Send")
        send_button.clicked.connect(self.send_message)
        input_layout.addWidget(send_button)

        # 将输入框布局添加到主布局
        main_layout.addLayout(input_layout)

        # 创建中央窗口小部件并应用主布局
        central_widget = QWidget()
        central_widget.setLayout(main_layout)
        self.setCentralWidget(central_widget)

    def send_message(self):
        message = self.input_field.text().strip()
        if message:
            # 在这里模拟发送消息并更新消息显示区域
            self.append_message(f"User: {message}")
            self.input_field.clear()

    def append_message(self, message: str):
        self.message_display.append(message)
        self.message_display.moveCursor(QTextCursor.End)
