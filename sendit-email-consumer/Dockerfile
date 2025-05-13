FROM rust:1.86.0 AS builder

WORKDIR /app

COPY ./src ./src
COPY ./Cargo.toml ./Cargo.toml

RUN cargo build --release

FROM rust:1.86.0-slim

COPY --from=builder /app/target/release/sendit-email-consumer .

CMD ["./sendit-email-consumer"]
