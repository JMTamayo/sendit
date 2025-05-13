use lettre::{
    message::{Mailbox, Message},
    transport::smtp::{authentication::Credentials, SmtpTransport},
    Transport,
};

use crate::{EmailData, Error, CONFIG};

pub struct Mailer {
    transport: SmtpTransport,
}

impl Mailer {
    fn get_transport(&self) -> &SmtpTransport {
        &self.transport
    }

    pub fn new() -> Result<Self, Error> {
        let transport: SmtpTransport = SmtpTransport::relay(CONFIG.get_email_relay())?
            .credentials(Credentials::new(
                CONFIG.get_email_username().to_string(),
                CONFIG.get_email_password().to_string(),
            ))
            .build();

        Ok(Self { transport })
    }

    pub fn send(&self, data: &EmailData) -> Result<(), Error> {
        let from: Mailbox = match CONFIG.get_email_username().parse::<Mailbox>() {
            Ok(from) => from,
            Err(e) => return Err(Error::new(&e.to_string())),
        };

        let to: Mailbox = match data.get_recipient().parse::<Mailbox>() {
            Ok(to) => to,
            Err(e) => return Err(Error::new(&e.to_string())),
        };

        let message: Message = Message::builder()
            .from(from)
            .to(to)
            .subject(data.get_subject())
            .body(data.get_body().to_string())
            .unwrap();

        match self.get_transport().send(&message) {
            Ok(_) => Ok(()),
            Err(e) => Err(Error::new(&e.to_string())),
        }
    }
}
