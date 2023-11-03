pub mod tcp_client_commands{
    use std::net::{TcpStream, ToSocketAddrs};

    #[tauri::command(async)]
    pub fn test_connection(invoke_message: String)->Result<bool,String>{
        let addr = "0.0.0.0:8080";
        let mut stream = TcpStream::connect(addr);
        println!("I was invoked from JS!{}",invoke_message);
        let socket_addrs: Vec<_> = match addr.to_socket_addrs() {
            Ok(addrs) => addrs.collect(),
            Err(err) => return Err(format!("Address resolution error: {}", err)),
        };
        // 遍历解析出来的地址并尝试连接
        for socket_addr in socket_addrs {
            match TcpStream::connect(socket_addr) {
                Ok(mut stream) => {
                    println!("Connected successfully! Invoked from JS: {}", invoke_message);
                    return Ok(true);
                }
                Err(err) => {
                    println!("Connection error: {}", err);
                }
            }
        }

        // 如果没有成功连接，返回错误信息
        Err("Failed to establish a connection".to_string())
    }
}
