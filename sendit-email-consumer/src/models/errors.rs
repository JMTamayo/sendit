use lettre::transport::smtp::Error as SmtpError;
use redsumer::results::RedsumerError;

#[derive(Debug)]
pub struct Error {
    details: String,
}

impl Error {
    pub fn get_details(&self) -> &str {
        &self.details
    }

    pub fn new(details: &str) -> Self {
        Self {
            details: details.to_string(),
        }
    }
}

impl From<RedsumerError> for Error {
    fn from(error: RedsumerError) -> Self {
        Self::new(&error.to_string())
    }
}

impl From<SmtpError> for Error {
    fn from(error: SmtpError) -> Self {
        Self::new(&error.to_string())
    }
}
