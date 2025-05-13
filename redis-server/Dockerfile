FROM redis:7.4-alpine

# Copy custom Redis configuration
COPY redis.conf /usr/local/etc/redis/redis.conf

# Expose Redis port
EXPOSE 6379

# Use the custom configuration
CMD ["redis-server", "/usr/local/etc/redis/redis.conf"] 