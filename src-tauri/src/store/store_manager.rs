fn create_store(){
    if let Some(base_dirs) = BaseDirs::new() {
        base_dirs.executable_dir();
        // Lin: Some(/home/alice/.local/bin)
        // Win: None
        // Mac: None
    }
}