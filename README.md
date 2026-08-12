<div align="center">

**English** · [Русский](docs/ru/README.ru.md)

**Docs:** [API Reference](./API.md) · [API-справочник (RU)](docs/ru/API.ru.md) · [Examples](./examples)

</div>

---

# gogrammy

A simple and clean Telegram bot wrapper for Go, inspired by [Grammy](https://grammy.dev).

A fluent builder-API wrapper over [go-telegram/bot](https://github.com/go-telegram/bot) — no more bulky `&Param{}` structs.

```go
// before (go-telegram/bot directly)
b.SendMessage(ctx, &bot.SendMessageParams{
    ChatID:    userID,
    Text:      "Hello!",
    ParseMode: models.ParseModeHTML,
})

// with gogrammy
ctx.SendText(userID, "Hello!").
    ParseMode(models.ParseModeHTML).
    Do(ctx.Ctx)
```

## Why gogrammy

- **Builders instead of param structs.** Every content type (text, photo, video, document, audio, voice, video note) has its own builder with a chainable API. Your IDE's autocomplete only shows what's actually valid for that type.
- **One `Context` type for everything.** No juggling between `Client` and `Context` — sending, editing, and keyboards all live in one place.
- **Event-based routing.** `Command`, `On`, `OnCallback` — register handlers without writing your own dispatcher, built on top of `go-telegram/bot`.
- **Nothing reinvented.** gogrammy doesn't replace the HTTP client or the Telegram API layer — it uses the stable `go-telegram/bot` under the hood and adds a convenience layer on top.

## Install

```bash
go get github.com/zzisler/gogrammy@latest
```

## Quick start

```go
package main

import (
    "context"
    "os"

    "github.com/zzisler/gogrammy"
)

func main() {
    b, err := gogrammy.New(os.Getenv("BOT_TOKEN"))
    if err != nil {
        panic(err)
    }

    b.Command("/start", func(c *gogrammy.Context) {
        userID := c.Update.Message.From.ID
        c.SendText(userID, "Hi! I'm a gogrammy bot 👋").Do(c.Ctx)
    })

    b.Start(context.Background())
}
```

More examples in [`examples/`](./examples).

## What's included

| Category | Methods |
|---|---|
| Sending | `SendText`, `SendPhoto`, `SendVideo`, `SendDocument`, `SendAudio`, `SendVoice`, `SendVideoNote` |
| Editing | `EditText`, `EditCaption`, `EditMedia` |
| Deleting | `DeleteMessage` |
| Keyboards | `NewInlineKeyboard`, `NewReplyKeyboard`, `RemoveKeyboard` |
| Join requests | `ApproveJoin`, `DeclineJoin` |
| Other | `SendChatAction`, `AnswerCallback` |
| Routing | `Command`, `On`, `OnCallback`, `Start` |

Full method reference with all optional parameters and examples: [`API.md`](./API.md).

## Working with files

Every media builder accepts a source in three ways:

```go
ctx.SendPhoto(userID).FileID("AAA...")        // file already on Telegram's servers
ctx.SendPhoto(userID).FileURL("https://...")  // direct link
ctx.SendPhoto(userID).FilePath("./photo.png") // upload from disk
```

## Keyboards

Inline keyboard — attached to a specific message, sends a `callback_query` on tap:

```go
kb := ctx.NewInlineKeyboard().
    Text("Button 1", "btn1").
    Row().
    URL("Website", "https://example.com").
    Build()

ctx.SendText(userID, "Choose:").ReplyMarkup(kb).Do(ctx.Ctx)
```

Reply keyboard — replaces the user's system keyboard, taps come back as regular text messages:

```go
kb := ctx.NewReplyKeyboard().
    Text("Menu").
    Row().
    Text("Help").
    Resize().
    Build()

ctx.SendText(userID, "Choose:").ReplyMarkup(kb).Do(ctx.Ctx)
```

Removing a reply keyboard:

```go
ctx.SendText(userID, "Keyboard removed").ReplyMarkup(ctx.RemoveKeyboard()).Do(ctx.Ctx)
```

## Status

Actively developed. Core functionality (sending/editing/deleting all content types, keyboards, routing, join requests) is implemented and tested against a live bot. Feedback and issues are welcome.

## License

MIT