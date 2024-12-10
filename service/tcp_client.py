import socket

from loguru import logger


async def tcp_client(host: str, port: int):
    client_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client_socket.connect((host, port))
    logger.debug(f"Connected to {host}:{port}")
    try:
        # 连接到指定的主机和端口
        client_socket.connect((host, port))
        logger.debug(f"Connected to {host}:{port}")

        while True:
            # 发送数据
            message = input("Enter message to send to server (or 'exit' to quit): ")
            if message.lower() == "exit":
                logger.info("Exiting...")
                break

            client_socket.sendall(message.encode("utf-8"))
            print(f"Sent message: {message}")

            # 接收服务器的响应
            data = client_socket.recv(1024)
            logger.debug(f"Received message: {data}")

    except Exception as e:
        print(f"Error: {e}")
    finally:
        # 关闭连接
        client_socket.close()
        print("Connection closed.")
