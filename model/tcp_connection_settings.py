class TcpConnectionSettings:
    def __init__(self, host, port, protocol, timeout, encryption_enabled):
        self.host = host
        self.port = port
        self.protocol = protocol
        self.timeout = timeout
        self.encryption_enabled = encryption_enabled

    def __repr__(self):
        return f"TcpConnectionSettings(host={self.host}, port={self.port}, protocol={self.protocol}, timeout={self.timeout}, encryption_enabled={self.encryption_enabled})"
