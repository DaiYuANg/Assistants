// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

use tauri_plugin_log::LogTarget;
use tauri::Manager;
use window_shadows::set_shadow;
use window_vibrancy::{apply_vibrancy, NSVisualEffectMaterial};
use crate::tcp::tcp_client_connection_manager::tcp_connection_manager::ConnectionManager;

mod support_protocol;
mod util;
mod tcp;
mod udp;
mod session;

fn main() {
    builder()
}
pub fn builder(){
    tauri::Builder::default()
        .manage(ConnectionManager::new)
        .setup(|app| {
            let window = app.get_window("main").unwrap();
            #[cfg(target_os = "macos")]
            apply_vibrancy(
                &window,
                NSVisualEffectMaterial::HudWindow,
                None,
                Option::from(50.0),
            )
                .expect("Unsupported platform! 'apply_vibrancy' is only supported on macOS");
            #[cfg(any(windows, target_os = "macos"))]
            set_shadow(&window, true).unwrap();
            #[cfg(target_os = "windows")]
            apply_blur(&window, Some((18, 18, 18, 125)))
                .expect("Unsupported platform! 'apply_blur' is only supported on Windows");
            Ok(())
        })
        .setup(|app|{
            Ok(())
        })
        .plugin(
            tauri_plugin_log::Builder::default()
                .targets([LogTarget::LogDir, LogTarget::Stdout, LogTarget::Webview])
                .build(),
        )
        .plugin(tauri_plugin_store::Builder::default().build())
        .plugin(tauri_plugin_window_state::Builder::default().build())
        .invoke_handler(tauri::generate_handler![
            tcp::tcp_client_commands::tcp_client_commands::test_connection
        ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
