from PySide6.QtGui import QScreen, QAction
from PySide6.QtWidgets import QLabel, QMainWindow, QHBoxLayout, QWidget, QVBoxLayout, QTextEdit, QSplitter, QToolBar, \
    QStatusBar, QApplication
from loguru import logger


class MainWindow(QMainWindow):
    def __init__(self):
        super().__init__()
        self.initUI()

    def initUI(self):
        self.setWindowTitle("Draggable Layout with Toolbar and Status Bar")

        # 创建工具栏
        toolbar = QToolBar()
        action1 = QAction("Action1", self)
        action1.triggered.connect(self.on_action)
        toolbar.addAction(action1)
        toolbar.addAction("Action 2")

        # 创建状态栏
        statusbar = QStatusBar()
        statusbar.showMessage("Ready")

        # 创建可拖动布局
        splitter = QSplitter()

        # 创建左右两个部件
        left_widget = QWidget()
        right_widget = QWidget()

        # 在左部件中添加文本编辑器
        text_edit = QTextEdit()
        left_layout = QVBoxLayout(left_widget)
        left_layout.addWidget(text_edit)

        # 将工具栏、左右部件和状态栏添加到可拖动布局中
        splitter.addWidget(toolbar)
        splitter.addWidget(left_widget)
        splitter.addWidget(right_widget)

        # 设置状态栏
        self.setStatusBar(statusbar)

        self.setCentralWidget(splitter)

    def center_window(self):
        center = QScreen.availableGeometry(QApplication.primaryScreen()).center()
        geo = self.frameGeometry()
        geo.moveCenter(center)
        self.move(geo.topLeft())

    def on_action(self):
        logger.info("Action triggered")
