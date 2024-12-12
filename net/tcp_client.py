class TCPClient:
  def __init__(self, host, port):
    self.host = host
    self.port = port
    self.connection = None

  def connect(self):
    print(f"Connecting to {self.host}:{self.port}")
    # 连接逻辑

  def send_data(self, data):
    print(f"Sending data: {data}")
    # 发送数据逻辑

class TCPClientManager:
  def __init__(self):
    self.clients = {}

  def send_data_to_client(self, data):
    pass
    # self.tcp_client.send_data(data)

  def connect_to_client(self):
    pass
    # self.tcp_client.connect()