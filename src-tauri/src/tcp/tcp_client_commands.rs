pub mod tcp_client_commands{
    use std::net::{TcpStream, ToSocketAddrs};
    use log::{error, info};

    #[tauri::command(async)]
    pub fn test_connection(addr: String)->Result<bool,String>{
        let mut stream = TcpStream::connect(addr.clone());
        info!("Test connection!{}",addr);
        let socket_addrs: Vec<_> = match addr.to_socket_addrs() {
            Ok(addrs) => addrs.collect(),
            Err(err) => return Err(format!("Address resolution error: {}", err)),
        };
        for socket_addr in socket_addrs {
            if let Ok(stream) = TcpStream::connect(socket_addr) {
                info!("Connected successfully! Invoked from JS: {}", addr);
                return Ok(true);
            } else if let Err(err) = TcpStream::connect(socket_addr) {
                error!("Connection error: {}", err);
                return Err(format!("Connection error: {}", err));
            } else {
                unreachable!()
            }
        }

        // 如果没有成功连接，返回错误信息
        Err("Failed to establish a connection".to_string())
    }
}
