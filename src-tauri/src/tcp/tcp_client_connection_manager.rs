pub mod tcp_connection_manager{
    use std::collections::HashMap;
    use std::net::TcpStream;
    use std::sync::Arc;
    use tokio::sync::Mutex;
    pub struct ConnectionManager {
        connections: Arc<Mutex<HashMap<String, TcpStream>>>,
    }

    impl ConnectionManager {
        pub fn new() -> ConnectionManager {
            ConnectionManager {
                connections: Arc::new(Mutex::new(HashMap::new())),
            }
        }
        pub fn add_connection(&self, client_id: String, stream: TcpStream) {
            // let mut connections = self.connections.lock().unwrap();
            // connections.insert(client_id, stream);
        }

        pub fn remove_connection(&self, client_id: &str) {
            // let mut connections = self.connections.lock().unwrap();
            // connections.remove(client_id);
        }

        pub fn get_connection(&self, client_id: &str) -> Option<TcpStream> {
            return None;
            // let connections = self.connections.lock().unwrap();
            // connections.get(client_id).cloned()
        }
    }
}