from PySide6.QtWidgets import (
    QTabWidget,
    QWidget,
    QVBoxLayout,
    QLabel,
    QPushButton,
)

class MessagePane(QTabWidget):
    def __init__(self):
        super().__init__()
        # 创建 TabWidget
        # 创建标签页 1
        tab1 = QWidget()
        layout1 = QVBoxLayout()
        layout1.addWidget(QLabel("这是第一个标签页"))
        layout1.addWidget(QPushButton("按钮 1"))
        tab1.setLayout(layout1)

        # 创建标签页 2
        tab2 = QWidget()
        layout2 = QVBoxLayout()
        layout2.addWidget(QLabel("这是第二个标签页"))
        layout2.addWidget(QPushButton("按钮 2"))
        tab2.setLayout(layout2)

        # 创建标签页 3
        tab3 = QWidget()
        layout3 = QVBoxLayout()
        layout3.addWidget(QLabel("这是第三个标签页"))
        layout3.addWidget(QPushButton("按钮 3"))
        tab3.setLayout(layout3)

        # 添加标签页到 TabWidget
        self.addTab(tab1, "标签 1")
        self.addTab(tab2, "标签 2")
        self.addTab(tab3, "标签 3")
