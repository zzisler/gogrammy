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
        c.Send(c.UserID(), "Привет!", nil)
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
- `Send(chatID, text, params)` — send a message, returns `(*models.Message, error)`
- `Edit(text, params)` — edit current message
- `EditMessage(messageID, text, params)` — edit specific message by ID
- `Delete()` — delete current message
- `DeleteMessage(messageID)` — delete specific message by ID
- `Forward(fromChatID, messageID)` — forward specific message
- `Copy(fromChatID, messageID, params)` — copy specific message
- `AnswerCallback(text?)` — answer callback query

### Media
- `SendPhoto(chatID, url, params)` — send photo by URL
- `SendAudio(chatID, params)` — send audio file or cached audio by file_id

```go
// new file
c.SendAudio(c.UserID(), &gogrammy.AudioParams{
    Title:     "Track Title",
    Performer: "Artist",
    Filename:  "track.mp3",
    TrackData: data,
    CoverData: cover,
})

// cached (by Telegram file_id)
c.SendAudio(c.UserID(), &gogrammy.AudioParams{
    Title:     "Track Title",
    Performer: "Artist",
    FileID:    "BQACAgIAAxkB...",
})
```

### Context helpers
- `c.UserID()` — `c.Update.Message.From.ID`
- `c.FirstName()` — `c.Update.Message.From.FirstName`
- `c.ChatID()` — `c.Update.Message.Chat.ID`
- `c.Text()` — `c.Update.Message.Text`

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