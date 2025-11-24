package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker/v2"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Initialize bot with mention-only mode
	bot := slacker.NewClient(
		os.Getenv("SLACK_BOT_TOKEN"),
		os.Getenv("SLACK_APP_TOKEN"),
		slacker.WithBotMode(slacker.BotModeIgnoreAll),
	)

	// Hello command
	bot.AddCommand(&slacker.CommandDefinition{
		Command:     "hello",
		Description: "Says hello",
		Handler: func(ctx *slacker.CommandContext) {
			ctx.Response().Reply("Hello, world!")
		},
	})

	// Compliment command
	bot.AddCommand(&slacker.CommandDefinition{
		Command:     "compliment me",
		Description: "Gives you a random compliment",
		Handler: func(ctx *slacker.CommandContext) {
			compliments := []string{
				"You're awesome! 😎",
				"You're doing great! 💪",
				"Keep up the amazing work! 🌟",
			}
			ctx.Response().Reply(compliments[rng.Intn(len(compliments))])
		},
	})

	// Dice roll command
	bot.AddCommand(&slacker.CommandDefinition{
		Command:     "roll a dice",
		Description: "Rolls a 6-sided dice",
		Handler: func(ctx *slacker.CommandContext) {
			roll := rng.Intn(6) + 1
			ctx.Response().Reply(fmt.Sprintf("You rolled a %d 🎲", roll))
		},
	})

	// Current time in Trinidad command
	bot.AddCommand(&slacker.CommandDefinition{
		Command:     "what time is it",
		Description: "Tells the current time in Trinidad",
		Handler: func(ctx *slacker.CommandContext) {

			loc, err := time.LoadLocation("America/Port_of_Spain")
			if err != nil {

				loc = time.FixedZone("AST", -4*60*60)
			}

			now := time.Now().In(loc).Format("15:04:05 MST on 02 Jan 2006")
			ctx.Response().Reply(fmt.Sprintf("The current time in Trinidad is %s 🕒", now))
		},
	})

	// Age calculator command
	bot.AddCommand(&slacker.CommandDefinition{
		Command:     "I was born in <year>",
		Description: "Calculates your age",
		Examples:    []string{"I was born in 2004"},
		Handler: func(ctx *slacker.CommandContext) {
			yearParam := ctx.Request().Param("year")
			birthYear, err := strconv.Atoi(yearParam)
			if err != nil {
				ctx.Response().Reply("Please provide a valid year.")
				return
			}

			currentYear := time.Now().Year()

			// Validate year
			if birthYear < 1900 || birthYear > currentYear {
				ctx.Response().Reply("Please provide a realistic birth year between 1900 and now.")
				return
			}

			age := currentYear - birthYear
			ctx.Response().Reply(fmt.Sprintf("You are %d years old!", age))
		},
	})

	// Help command
	bot.AddCommand(&slacker.CommandDefinition{
		Command:     "help",
		Description: "Lists all available commands",
		Handler: func(ctx *slacker.CommandContext) {
			helpText := `*Available commands:*
• *hello* - Says hello
• *compliment me* - Gives you a random compliment
• *roll a dice* - Rolls a 6-sided dice
• *what time is it* - Tells the current time
• *I was born in <year>* - Calculates your age (example: I was born in 2004)
• *help* - Shows this message

Just mention me with any of these commands!`
			ctx.Response().Reply(helpText)
		},
	})

	// Start the bot

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Bot is starting...")
	err = bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
