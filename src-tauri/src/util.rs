use std::net::TcpListener;

pub fn is_port_in_use(port: u16) -> bool {
    match TcpListener::bind(format!("127.0.0.1:{}", port)) {
        Ok(_) => {
            // 端口没有被占用
            true
        }
        Err(_) => {
            // 端口已被占用
            false
        }
    }
}