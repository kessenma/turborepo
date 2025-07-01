# ---------- Stage 1: Builder ----------
FROM golang:1.21-alpine AS builder

# Install dependencies for the build
RUN apk add --no-cache ca-certificates git

WORKDIR /app

# Copy go.mod first
COPY go.mod ./

# Initialize go.sum (important for first-time builds)
RUN go mod tidy

# Copy source code
COPY . .

# Make sure dependencies are properly downloaded and go.sum is updated
RUN go mod tidy

# Build the application with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o bot .

# ---------- Stage 2: Minimal Runtime ----------
FROM alpine:3.19

# Install certificates for HTTPS connections
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/bot .

# Set executable permissions
RUN chmod +x /app/bot

# Run the application
CMD ["/app/bot"]
WORKDIR /usr/src/bot

# Copy Cargo.toml and Cargo.lock first for layer caching
COPY Cargo.toml ./

# Create a dummy source file to build dependencies
RUN mkdir src && echo 'fn main() { println!("Dummy"); }' > src/main.rs
RUN cargo build --release
RUN rm -rf src

# Now copy the actual source code
COPY . .

# Build the actual binary and inspect output location
RUN cargo build --release && ls -la /usr/src/bot/target/x86_64-unknown-linux-musl/release/ || ls -la /usr/src/bot/target/release/ || find /usr/src/bot -name bot -type f

# ---------- Stage 2: Minimal Runtime ----------
FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /usr/src/bot
COPY --from=builder /usr/src/bot/target/x86_64-unknown-linux-musl/release/bot .

CMD ["./bot"]
