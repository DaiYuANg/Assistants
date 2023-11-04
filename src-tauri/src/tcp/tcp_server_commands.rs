pub mod tcp_server_commands{
    use crate::util::is_port_in_use;

    #[tauri::command(async)]
    pub fn test_tcp_server() -> Result<bool,String> {
        let is_usage = is_port_in_use(8080);

        if is_usage {
            return Err("Failed to create server".to_string())
        }

        return Ok(true)
    }
}