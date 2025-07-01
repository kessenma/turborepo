package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	fmt.Println("🧪 Entered main function...")

	// Get token from environment variable
	token := os.Getenv("TELOXIDE_TOKEN")
	if token == "" {
		log.Fatal("❌ TELOXIDE_TOKEN environment variable not set")
	}

	fmt.Println("✅ Bot token loaded.")

	// Create bot instance
	b, err := bot.New(token)
	if err != nil {
		log.Fatalf("❌ Failed to create bot: %v", err)
	}

	fmt.Println("🟢 Starting bot...")

	// Register message handler
	b.RegisterHandler(bot.HandlerTypeMessage, "/start", bot.MatchTypeExact, handleStart)
	// Handle all messages
	b.RegisterHandler(bot.HandlerTypeMessage, "", bot.MatchTypePrefix, handleMessage)

	// Start the bot
	b.Start(context.Background())
}

// handleStart handles the /start command
func handleStart(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	fmt.Printf("📩 Received /start command from %d\n", update.Message.Chat.ID)

	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Hello! I'm a simple Telegram bot. Send me a message and I'll respond with a dice.",
	})

	if err != nil {
		fmt.Printf("❌ Error sending message: %v\n", err)
	}
}

// handleMessage handles all other messages
func handleMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	fmt.Printf("📩 Received message from %d: %s\n", update.Message.Chat.ID, update.Message.Text)

	// Send a dice in response
	_, err := b.SendDice(ctx, &bot.SendDiceParams{
		ChatID: update.Message.Chat.ID,
	})

	if err != nil {
		fmt.Printf("🎲 Failed to send dice: %v\n", err)
	}
}
