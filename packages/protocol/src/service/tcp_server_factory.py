import socket
from loguru import logger


class TCPServerFactory:
    def __init__(self):
        server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        server_socket.bind(("localhost", 8090))
        server_socket.listen(5)
        logger.info(f"Listening on port 8090")
