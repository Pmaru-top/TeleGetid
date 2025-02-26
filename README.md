# TeleGetid

TeleGetid is a simple Telegram bot that allows you to retrieve user IDs. It is built using the [telebot](https://github.com/tucnak/telebot) framework.

## Features
- Retrieve Telegram user IDs quickly and easily.
- Simple setup and usage.

## Online use


## How to use

### Windows
[Download](https://github.com/pmaru-top/telegetid/releases/download/1.0.0/telegetid_windows_amd64) and run

### linux_amd64
1. Download the latest release
   ```sh
    wget https://github.com/pmaru-top/telegetid/releases/download/1.0.0/telegetid_linux_amd64
   ```
2. Initialize the bot
   ```sh
   chmod +rwx ./telegetid_linux_amd64
   ./telegetid_linux_amd64
   ```
3. On the first run, you will see the following message:
   ```
   enter your token in ./config
   ```
4. Open the `./config` file and enter your bot token.
5. Restart the bot:
   ```sh
   nohup ./telegetid_linux_amd64 &
   ```

## How to get bot token
- [@BotFather](https://t.me/BotFather)

## Build manually
```sh
go build -trimpath -ldflags "-s -w -buildid=" -o bin/telegetid ./main
```

or build scripts [build.ps1](build.ps1) (only windows)

## License
This project is licensed under the MIT License.