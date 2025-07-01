# Telegram Bot in Go

A simple Telegram bot built with Go that responds to messages with a dice emoji.

## Features

- Responds to `/start` with a welcome message
- Responds to any other message with a dice emoji
- Uses the official Telegram Bot API via the go-telegram/bot library

## Requirements

- Go 1.21 or later
- Docker (for containerized deployment)

## Environment Variables

- `TELOXIDE_TOKEN`: Your Telegram bot token (obtainable from BotFather)

## Running Locally

```bash
# Set your bot token
export TELOXIDE_TOKEN=your_token_here

# Run the bot
go run main.go
```

## Building and Running with Docker

```bash
# Build the Docker image
docker build -t telegram-bot .

# Run the Docker container
docker run -e TELOXIDE_TOKEN=your_token_here telegram-bot
```

## Running with Docker Compose

See the docker-compose.yml file in the parent directory.
