use std::{env, sync::Arc};

use chrono::Utc;
use env_logger::{fmt, Builder, Env};

pub struct Logger;

impl Logger {
    pub fn build(log_level: &str) {
        Builder::from_env(Env::default().default_filter_or(log_level))
            .format_timestamp(Some(fmt::TimestampPrecision::Millis))
            .format_module_path(false)
            .init();
    }
}

pub struct Config {
    service_name: String,
    log_level: String,

    redis_username: String,
    redis_password: String,

    redis_host: String,
    redis_port: u16,
    redis_db: i64,

    stream_name_email_queue: String,
    consumer_group_name_send_email: String,

    consumer_options_since_id: String,
    consumer_options_new_messages_count: usize,
    consumer_options_pending_messages_count: usize,
    consumer_options_claimed_messages_count: usize,
    consumer_options_block: usize,
    consumer_options_min_idle_time_millisec: usize,

    email_relay: String,
    email_username: String,
    email_password: String,
}

impl Config {
    pub fn get_service_name(&self) -> &str {
        &self.service_name
    }

    pub fn get_log_level(&self) -> &str {
        &self.log_level
    }

    pub fn get_redis_username(&self) -> &str {
        &self.redis_username
    }

    pub fn get_redis_password(&self) -> &str {
        &self.redis_password
    }

    pub fn get_redis_host(&self) -> &str {
        &self.redis_host
    }

    pub fn get_redis_port(&self) -> u16 {
        self.redis_port
    }

    pub fn get_redis_db(&self) -> i64 {
        self.redis_db
    }

    pub fn get_stream_name_email_queue(&self) -> &str {
        &self.stream_name_email_queue
    }

    pub fn get_consumer_group_name_send_email(&self) -> &str {
        &self.consumer_group_name_send_email
    }

    pub fn get_consumer_options_since_id(&self) -> &str {
        &self.consumer_options_since_id
    }

    pub fn get_consumer_options_new_messages_count(&self) -> usize {
        self.consumer_options_new_messages_count
    }

    pub fn get_consumer_options_pending_messages_count(&self) -> usize {
        self.consumer_options_pending_messages_count
    }

    pub fn get_consumer_options_claimed_messages_count(&self) -> usize {
        self.consumer_options_claimed_messages_count
    }

    pub fn get_consumer_options_block(&self) -> usize {
        self.consumer_options_block
    }

    pub fn get_consumer_options_min_idle_time_millisec(&self) -> usize {
        self.consumer_options_min_idle_time_millisec
    }

    pub fn get_email_relay(&self) -> &str {
        &self.email_relay
    }

    pub fn get_email_username(&self) -> &str {
        &self.email_username
    }

    pub fn get_email_password(&self) -> &str {
        &self.email_password
    }

    fn new() -> Self {
        let id: i64 = Utc::now().timestamp_millis();

        let service_name: String = format!("email-consumer-{id}");
        let log_level: String = env::var("LOG_LEVEL").unwrap_or("info".to_string());

        let redis_username: String = env::var("REDIS_USERNAME").expect("REDIS_USERNAME not found");
        let redis_password: String = env::var("REDIS_PASSWORD").expect("REDIS_PASSWORD not found");

        let redis_host: String = env::var("REDIS_HOST").expect("REDIS_HOST not found");
        let redis_port: u16 = env::var("REDIS_PORT")
            .expect("REDIS_PORT not found")
            .parse::<u16>()
            .expect("REDIS_PORT is not a valid number");

        let redis_db: i64 = env::var("REDIS_DB")
            .expect("REDIS_DB not found")
            .parse::<i64>()
            .expect("REDIS_DB is not a valid number");

        let stream_name_email_queue: String =
            env::var("STREAM_NAME_EMAIL_QUEUE").expect("STREAM_NAME_EMAIL_QUEUE not found");
        let consumer_group_name_send_email: String = env::var("CONSUMER_GROUP_NAME_SEND_EMAIL")
            .expect("CONSUMER_GROUP_NAME_SEND_EMAIL not found");
        let consumer_options_since_id: String =
            env::var("CONSUMER_OPTIONS_SINCE_ID").expect("CONSUMER_OPTIONS_SINCE_ID not found");
        let consumer_options_new_messages_count: usize =
            env::var("CONSUMER_OPTIONS_NEW_MESSAGES_COUNT")
                .expect("CONSUMER_OPTIONS_NEW_MESSAGES_COUNT not found")
                .parse::<usize>()
                .expect("CONSUMER_OPTIONS_NEW_MESSAGES_COUNT is not a valid number");
        let consumer_options_pending_messages_count: usize =
            env::var("CONSUMER_OPTIONS_PENDING_MESSAGES_COUNT")
                .expect("CONSUMER_OPTIONS_PENDING_MESSAGES_COUNT not found")
                .parse::<usize>()
                .expect("CONSUMER_OPTIONS_PENDING_MESSAGES_COUNT is not a valid number");
        let consumer_options_claimed_messages_count: usize =
            env::var("CONSUMER_OPTIONS_CLAIMED_MESSAGES_COUNT")
                .expect("CONSUMER_OPTIONS_CLAIMED_MESSAGES_COUNT not found")
                .parse::<usize>()
                .expect("CONSUMER_OPTIONS_CLAIMED_MESSAGES_COUNT is not a valid number");
        let consumer_options_block: usize = env::var("CONSUMER_OPTIONS_BLOCK")
            .expect("CONSUMER_OPTIONS_BLOCK not found")
            .parse::<usize>()
            .expect("CONSUMER_OPTIONS_BLOCK is not a valid number");
        let consumer_options_min_idle_time_millisec: usize =
            env::var("CONSUMER_OPTIONS_MIN_IDLE_TIME_MILLISEC")
                .expect("CONSUMER_OPTIONS_MIN_IDLE_TIME_MILLISEC not found")
                .parse::<usize>()
                .expect("CONSUMER_OPTIONS_MIN_IDLE_TIME_MILLISEC is not a valid number");

        let email_relay: String = env::var("EMAIL_RELAY").expect("EMAIL_RELAY not found");
        let email_username: String = env::var("EMAIL_USERNAME").expect("EMAIL_USERNAME not found");
        let email_password: String = env::var("EMAIL_PASSWORD").expect("EMAIL_PASSWORD not found");

        Self {
            log_level,
            service_name,
            redis_username,
            redis_password,
            redis_host,
            redis_port,
            redis_db,
            stream_name_email_queue,
            consumer_group_name_send_email,
            consumer_options_since_id,
            consumer_options_new_messages_count,
            consumer_options_pending_messages_count,
            consumer_options_claimed_messages_count,
            consumer_options_block,
            consumer_options_min_idle_time_millisec,
            email_relay,
            email_username,
            email_password,
        }
    }
}

lazy_static::lazy_static! {
    pub static ref CONFIG: Arc<Config> = Arc::new(Config::new());
}
