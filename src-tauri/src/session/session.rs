use chrono::{NaiveDateTime, Utc};
use crate::support_protocol::SupportProtocol;

pub struct Session<S> {
    protocol:SupportProtocol,
    connection:S,
    timestamp: NaiveDateTime,
    active: bool
}

impl<S> Session<S> {
    pub fn new(protocol: SupportProtocol, connection: S,active:bool) -> Self {
        Session {
            protocol,
            connection,
            timestamp: Utc::now().naive_utc(),
            active
        }
    }
}