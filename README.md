# Discord Bot

### How to install (Ubuntu)

```
sudo apt update
sudo apt install snapd
sudo snap install go --classic

git clone <url> && cd DiscordBot
go mod tidy
```

### How to run (Ubuntu)

```
go build -o bot
./bot -t <token>
```