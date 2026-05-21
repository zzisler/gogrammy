# gogrammy

A simple and clean Telegram bot wrapper for Go, inspired by [Grammy](https://grammy.dev).

## Installation

```bash
go get github.com/zzisler/gogrammy
```

## Usage

```go
package main

import (
    "context"
    "github.com/zzisler/gogrammy"
)

func main() {
    app, err := gogrammy.New("YOUR_TOKEN")
    if err != nil {
        panic(err)
    }

    app.Command("/start", func(c *gogrammy.Context) {
        c.Send(c.Update.Message.Chat.ID, "Привет!", nil)
    })

    app.Start(context.Background())
}
```

## Methods

### App
- `New(token, proxy?)` — create bot instance
- `Start(ctx)` — start bot

### Commands
- `Command(cmd, handler)` — any chat command
- `PrivateCommand(cmd, handler)` — private chat only
- `GroupCommand(cmd, handler)` — group/supergroup only

### Events
- `On(eventType, handler)` — listen to events: `message`, `callback`, `join_request`, `my_chat_member`
- `OnCallback(prefix, handler)` — listen to callback with prefix

### Messages
- `Send(chatID, text, params)` — send a message
- `Edit(text, params)` — edit current message
- `EditMessage(messageID, text, params)` — edit specific message
- `Delete()` — delete current message
- `DeleteMessage(messageID)` — delete specific message
- `Forward(fromChatID, messageID)` — forward specific message
- `Copy(fromChatID, messageID, params)` — copy specific message
- `AnswerCallback(text?)` — answer callback query

### Media
- `SendPhoto(chatID, url, params)` — send photo
- `SendAudio(chatID, filename, title, artist, data, coverData)` — send audio

### Admin
- `Ban(chatID, userID, params)` — ban user
- `Unban(chatID, userID)` — unban user
- `BanChat(chatID, senderChatID)` — ban sender chat
- `UnbanChat(chatID, senderChatID)` — unban sender chat
- `GetProfilePhoto(userID, params)` — get user profile photos

### Join Requests
- `ApproveJoin(chatID, userID)` — approve join request
- `DeclineJoin(chatID, userID)` — decline join request

### Keyboards
- `Keyboard(rows...)` — create inline keyboard
- `Row(buttons...)` — create keyboard row
- `Button(text, data)` — create inline button

## License

MIT