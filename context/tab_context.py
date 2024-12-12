from PySide6.QtCore import QObject, Signal


class ClientContext(QObject):
  status_signal = Signal(int, str)  # tab 索引和消息内容

  def __init__(self):
    super().__init__()
    self.clients = []  # 存储 TcpClient 和其对应的标签页索引
    self.tab_index = 0  # 初始标签页索引

  def add_client(self, host: str, port: int):
    pass
  # tcp_client = TcpClient(host, port)
  # self.clients.append(tcp_client)
  # self.status_signal.emit(self.tab_index, "Connecting...")  # 显示正在连接
  # self.start_client(tcp_client)
  # self.tab_index += 1  # 为下一个连接准备新的标签页

  def start_client(self):
    pass
# 启动客户端连接（在后台线程中运行）
# asyncio.create_task(tcp_client.connect())
