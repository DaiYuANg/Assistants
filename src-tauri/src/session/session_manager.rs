pub mod session_manager {
    use std::collections::HashMap;
    use std::sync::Arc;
    use di::injectable;
    use tokio::sync::RwLock;
    use crate::session::session::Session;

    pub struct SessionManager{
        // sm_self: Mutex<Option<SessionManager>>,
       sessions: Arc<RwLock<HashMap<String, Session<String>>>>,
    }

    // lazy_static! {
    //     static ref INSTANCE: Mutex<Option<SessionManager>> = {
    //         Mutex::new(Some(SessionManager::new()))
        // };
    // }
    impl SessionManager {
        // fn new() -> Option<_> {
        //     // SessionManager {
        //     //     // sm_self: Mutex::new(None),
        //     //     // sessions: Arc::new(Default::default()),
        //     // }
        //     return None
        // }

        // fn get_instance() -> &'static Mutex<Option<SessionManager>> {
        // }
    }

    impl SessionManager {
        pub fn list_session(){

        }
    }
}