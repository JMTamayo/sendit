use redsumer::{
    client::{ClientArgs, ClientCredentials, CommunicationProtocol},
    consumer::{
        ClaimMessagesOptions, Consumer, ConsumerConfig, ReadNewMessagesOptions,
        ReadPendingMessagesOptions,
    },
};

use crate::{Error, CONFIG};

pub struct ConsumerBuilder;

impl ConsumerBuilder {
    pub fn build() -> Result<Consumer, Error> {
        let credentials: ClientCredentials =
            ClientCredentials::new(CONFIG.get_redis_username(), CONFIG.get_redis_password());

        let args: ClientArgs = ClientArgs::new(
            Some(credentials),
            CONFIG.get_redis_host(),
            CONFIG.get_redis_port(),
            CONFIG.get_redis_db(),
            CommunicationProtocol::RESP2,
        );

        let conf: ConsumerConfig = ConsumerConfig::new(
            CONFIG.get_stream_name_email_queue(),
            CONFIG.get_consumer_group_name_send_email(),
            CONFIG.get_service_name(),
            ReadNewMessagesOptions::new(
                CONFIG.get_consumer_options_new_messages_count(),
                CONFIG.get_consumer_options_block(),
            ),
            ReadPendingMessagesOptions::new(CONFIG.get_consumer_options_pending_messages_count()),
            ClaimMessagesOptions::new(
                CONFIG.get_consumer_options_claimed_messages_count(),
                CONFIG.get_consumer_options_min_idle_time_millisec(),
            ),
        );

        Ok(Consumer::new(
            args,
            conf,
            Some(CONFIG.get_consumer_options_since_id().to_string()),
        )?)
    }
}
