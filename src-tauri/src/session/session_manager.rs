pub mod session_manager {
    use crate::support_protocol::SupportProtocol;

    struct Session<S> {
        protocol:SupportProtocol,
        connection:S
    }

    struct SessionManager{
        
    }
}