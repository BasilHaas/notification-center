# notification-center
Application to generate messages by Webhooks.

## Build and Run

### 1. Export environment variables TG_CHAT_ID and TG_BOT_TOKEN
```bash
export TG_CHAT_ID=<chat_id_with_your_bot>
export TG_BOT_TOKEN=<your_bot_secret_token>
```

or write variables to environment file:
```bash
echo TG_CHAT_ID=<chat_id_with_your_bot> >> /etc/environment
echo TG_BOT_TOKEN=<your_bot_secret_token> >> /etc/environment
source /etc/environment
```

### 2. Compile and run application
```bash
go build main.go
chmod +x main.go
./main.go
```

Now you can send POST request with JSON data to :8000 port on your machine.
