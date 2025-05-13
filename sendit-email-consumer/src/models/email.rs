use log;
use redsumer::{
    redis::{from_redis_value, ErrorKind, FromRedisValue, Value},
    results::RedsumerResult,
};
use serde::Deserialize;

#[derive(Deserialize, Debug)]
pub struct EmailData {
    recipient: String,
    subject: String,
    body: String,
}

impl EmailData {
    pub fn get_recipient(&self) -> &str {
        &self.recipient
    }

    pub fn get_subject(&self) -> &str {
        &self.subject
    }

    pub fn get_body(&self) -> &str {
        &self.body
    }
}

impl FromRedisValue for EmailData {
    fn from_redis_value(v: &Value) -> RedsumerResult<Self> {
        let json_str: String = from_redis_value(v)?;
        match serde_json::from_str(&json_str) {
            Ok(data) => Ok(data),
            Err(e) => {
                log::error!(
                    "Error parsing email data from stream message: {e}",
                    e = e.to_string()
                );
                Err((
                    ErrorKind::TypeError,
                    "Error parsing email data from stream message",
                )
                    .into())
            }
        }
    }
}
