import socket


class SocketClient:
    def __init__(self):
        self.sock = socket.socket(socket.AF_INET, socket)
