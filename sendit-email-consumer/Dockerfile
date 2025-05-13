FROM rust:1.86.0

WORKDIR /app

COPY ./src ./src
COPY ./Cargo.toml ./Cargo.toml

RUN cargo build --release

FROM debian:12.10

COPY --from=builder /app/target/release/sendit-email-consumer .

CMD ["./sendit-email-consumer"]
