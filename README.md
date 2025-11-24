# Simple Slack Bot

A simple Slack bot built in Go using [Slacker v2](https://github.com/shomali11/slacker) that responds to mentions and direct messages with fun commands, age calculation, and more.

---
![Go](https://img.shields.io/badge/Go-1.21-blue?logo=go&logoColor=white)
![Stars](https://img.shields.io/github/stars/JOSIAHTHEPROGRAMMER/simple-slack-bot?style=social)
![Forks](https://img.shields.io/github/forks/JOSIAHTHEPROGRAMMER/simple-slack-bot?style=social)
![Open Issues](https://img.shields.io/github/issues/JOSIAHTHEPROGRAMMER/simple-slack-bot)
![Last Commit](https://img.shields.io/github/last-commit/JOSIAHTHEPROGRAMMER/simple-slack-bot)
![Go Report Card](https://goreportcard.com/badge/github.com/JOSIAHTHEPROGRAMMER/simple-slack-bot)
![Slack](https://img.shields.io/badge/Platform-Slack-blue?logo=slack&logoColor=white)

## Features

- **Hello** – Greets the user.
- **Compliment Me** – Gives a random compliment.
- **Roll a Dice** – Rolls a 6-sided dice.
- **What Time Is It** – Shows the current time.
- **I Was Born In <year>** – Calculates your age.
- **Help** – Lists all available commands.

The bot responds only when mentioned or in direct messages.

---

## Setup

1. **Clone the repository**

```bash
git clone https://github.com/JOSIAHTHEPROGRAMMER/simple-slack-bot.git
cd Slack-age-bot
```


### 2. Create a .env file in the project root:

```env
Copy code
SLACK_BOT_TOKEN=xoxb-your-bot-token
SLACK_APP_TOKEN=xapp-your-app-level-token
```
### 3. Install dependencies and run:

```bash
go mod tidy
go run main.go
```

## Slack App Permissions
Your Slack app needs the following OAuth scopes:

- app_mentions:read – Read messages that mention the bot

- chat:write – Send messages as the bot

- chat:write.public – Send messages in channels the bot isn’t a member of

- channels:read – View public channel info

- im:history & im:write – Read and send messages in direct messages

- groups:history – Read messages in private channels the bot is added to

- mpim:history – Read messages in group DMs

## Usage
Mention the bot in a channel or send a direct message.

Use any of the commands:

```less

@Simple Bot hello
@Simple Bot compliment me
@Simple Bot roll a dice
@Simple Bot what time is it
@Simple Bot I was born in 2004
@Simple Bot help
The bot will respond with the appropriate reply.
```
