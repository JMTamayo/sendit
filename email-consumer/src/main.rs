use redsumer::{
    consumer::{AckMessageReply, Consumer},
    redis::{FromRedisValue, StreamId, Value},
    results::RedsumerResult,
};

mod config;
use config::conf::{Logger, CONFIG};

mod models;
use models::{email::EmailData, errors::Error};

mod services;
use services::{email::Mailer, redis::ConsumerBuilder};

#[tokio::main]
async fn main() {
    Logger::build(CONFIG.get_log_level());

    log::info!("Starting service '{s}'", s = CONFIG.get_service_name());

    let email_service: Mailer = match Mailer::new() {
        Ok(service) => service,
        Err(e) => panic!("Error creating email service: {e}", e = e.get_details()),
    };

    let mut consumer: Consumer = match ConsumerBuilder::build() {
        Ok(consumer) => consumer,
        Err(e) => panic!("Error creating consumer: {e}", e = e.get_details()),
    };

    loop {
        log::info!(
            "Preparing to consume stream '{s}' by the consumer '{c}' subscribing to the group '{g}'",
            s = CONFIG.get_stream_name_email_queue(),
            c = CONFIG.get_service_name(),
            g = CONFIG.get_consumer_group_name_send_email(),
        );

        let messages: Vec<StreamId> = match consumer.consume().await {
            Ok(reply) => reply.get_messages().to_owned(),
            Err(e) => {
                log::error!("Error consuming messages: {e}");
                continue;
            }
        };

        log::info!("Total messages to process: {n}", n = messages.len());

        for m in messages {
            log::info!("Processing message: {m:?}");

            match consumer.is_still_mine(&m.id) {
                Ok(reply) => match reply.belongs_to_me() {
                    true => {
                        log::info!("Message is ready to be processed: {m:?}", m = m.id);
                    }
                    false => {
                        log::warn!("Message is still in consumer pending list: {m:?}", m = m.id);
                        continue;
                    }
                },
                Err(e) => {
                    log::error!("Error checking if message is still in consumer pending list: {e}",);
                    continue;
                }
            }

            let value: &Value = match m.map.get("data") {
                Some(value) => value,
                None => {
                    log::error!("Message does not contain 'data' field");
                    continue;
                }
            };

            let email_data: EmailData = match EmailData::from_redis_value(value) {
                Ok(data) => data,
                Err(e) => {
                    log::error!("Error parsing email data from stream message: {e}",);
                    continue;
                }
            };

            log::debug!("Email data: {data:?}", data = email_data);

            let result: Result<(), Error> = email_service.send(&email_data);
            match result {
                Ok(_) => log::info!("Email sent successfully: {m:?}", m = m.id),
                Err(e) => {
                    log::error!("Error sending email: {e}", e = e.get_details());
                    continue;
                }
            }

            let ack: RedsumerResult<AckMessageReply> = consumer.ack(&m.id).await;
            match ack {
                Ok(reply) => match reply.was_acked() {
                    true => log::info!("Message acknowledged: {m:?}", m = m.id),
                    false => log::error!("Error acknowledging message: {m:?}", m = m.id),
                },
                Err(e) => {
                    log::error!("Error acknowledging message: {e}");
                }
            }
        }
    }
}
