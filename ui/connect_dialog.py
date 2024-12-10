from PySide6.QtWidgets import (
    QDialog,
    QLabel,
    QDialogButtonBox,
    QVBoxLayout,
    QLineEdit,
    QFormLayout,
    QComboBox,
    QGroupBox,
    QCheckBox,
)

from model.tcp_connection_settings import TcpConnectionSettings


class ConnectDialog(QDialog):
    def __init__(self, parent=None):
        super().__init__(parent)
        self.setWindowTitle("TCP Connection Settings")

        # 主布局
        main_layout = QVBoxLayout(self)

        # 创建表单布局
        form_layout = QFormLayout()

        # 添加表单字段：TCP 地址
        self.host_input = QLineEdit(self)
        form_layout.addRow("Host Address:", self.host_input)

        # 添加表单字段：端口
        self.port_input = QLineEdit(self)
        form_layout.addRow("Port:", self.port_input)

        # 添加表单字段：选择协议
        self.protocol_combo = QComboBox(self)
        self.protocol_combo.addItems(["TCP", "UDP"])
        form_layout.addRow("Protocol:", self.protocol_combo)

        # 高级配置折叠部分
        self.advanced_group = QGroupBox("Advanced Configuration", self)
        self.advanced_group.setCheckable(True)
        self.advanced_group.setChecked(False)  # 默认收起
        self.advanced_group.toggled.connect(self.toggle_advanced)

        advanced_layout = QFormLayout()

        # 高级配置内容：TCP 超时
        self.timeout_input = QLineEdit(self)
        advanced_layout.addRow("Timeout (ms):", self.timeout_input)

        # 高级配置内容：是否启用加密
        self.encryption_check = QCheckBox("Enable Encryption", self)
        advanced_layout.addRow(self.encryption_check)

        self.advanced_group.setLayout(advanced_layout)

        # 将表单布局和高级配置组添加到主布局
        main_layout.addLayout(form_layout)
        main_layout.addWidget(self.advanced_group)

        # 创建 OK 和 Cancel 按钮
        button_box = QDialogButtonBox(
            QDialogButtonBox.Ok | QDialogButtonBox.Cancel, self
        )
        button_box.accepted.connect(self.accept)
        button_box.rejected.connect(self.reject)

        main_layout.addWidget(button_box)

        self.setLayout(main_layout)

    def toggle_advanced(self):
        # 控制高级配置部分的显示和隐藏
        if self.advanced_group.isChecked():
            self.advanced_group.setStyleSheet(
                "QGroupBox { border: 1px solid gray; margin-top: 10px; padding: 10px; }"
            )
        else:
            self.advanced_group.setStyleSheet(
                "QGroupBox { border: 1px solid gray; margin-top: 10px; padding: 10px; display:none; }"
            )

    def accept(self):
        # 调用父类的 accept 方法来接受对话框，关闭它
        super().accept()

    def get_data(self):
        # 获取表单数据
        host = self.host_input.text()
        port = self.port_input.text()
        protocol = self.protocol_combo.currentText()
        timeout = self.timeout_input.text()
        encryption_enabled = self.encryption_check.isChecked()

        # 创建并返回 TcpConnectionSettings 实例
        return TcpConnectionSettings(host, port, protocol, timeout, encryption_enabled)
