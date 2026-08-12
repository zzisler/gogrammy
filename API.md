<div align="center">

**English** · [Русский](docs/ru/API.ru.md)

**Docs:** [README](./README.md) · [README (RU)](docs/ru/README.ru.md) · [Examples](./examples)

</div>

---

# API Reference

## Contents
- [Sending messages](#sending-messages)
  - [SendText](#sendtext)
  - [SendPhoto](#sendphoto)
  - [SendVideo](#sendvideo)
  - [SendAudio](#sendaudio)
  - [SendDocument](#senddocument)
  - [SendVoice](#sendvoice)
  - [SendVideoNote](#sendvideonote)
- [Editing](#editing)
  - [EditText](#edittext)
  - [EditCaption](#editcaption)
  - [EditMessageMedia](#editmessagemedia)
- [Deleting](#deleting)
  - [DeleteMessage](#deletemessage)
- [Keyboards](#keyboards)
  - [NewInlineKeyboard](#newinlinekeyboard)
  - [NewReplyKeyboard](#newreplykeyboard)
  - [RemoveKeyboard](#removekeyboard)
- [Join requests](#join-requests)
  - [ApproveJoin](#approvejoin)
  - [DeclineJoin](#declinejoin)
- [Miscellaneous](#miscellaneous)
  - [AnswerCallback](#answercallback)
  - [SendChatAction](#sendchataction)
- [Routing](#routing)
  - [Command](#command)
  - [On](#on)
  - [OnCallback](#oncallback)
  - [Start](#start)

---

## Sending messages

### SendText

`ctx.SendText(userID int64, text string)` -> `*TextBuilder`

Creates a builder for sending a text message.

| Method | Required | Description |
|---|---|---|
| `.ParseMode(mode models.ParseMode)` | no | Sets the parsing mode (HTML, Markdown, etc.) |
| `.ReplyTo(messageID int)` | no | Specifies the ID of the message to reply to |
| `.ReplyMarkup(kb models.ReplyMarkup)` | no | Adds a keyboard or inline markup |
| `.Do(ctx context.Context)` | — | Sends the message. Returns `(*models.Message, error)` |

**Example:**
```go
msg, err := ctx.SendText(userID, "Привет!").
    ParseMode(models.ModeHTML).
    ReplyTo(123).
    ReplyMarkup(inlineKeyboard).
    Do(ctx.Ctx)
```

[Back to top](#api-reference)

---

### SendPhoto

`ctx.SendPhoto(userID int64)` -> `*PhotoBuilder`

Creates a builder for sending a photo.

| Method | Required | Description |
|---|---|---|
| `.FileID(id string)` | yes* | Sets the photo by file_id |
| `.FileURL(url string)` | yes* | Sets the photo by URL |
| `.FilePath(path string)` | yes* | Sets the photo by local path (reads the file) |
| `.Caption(caption string)` | no | Sets the photo caption |
| `.ParseMode(mode models.ParseMode)` | no | Sets the parsing mode for the caption |
| `.ReplyMarkup(kb models.ReplyMarkup)` | no | Adds a keyboard or inline markup |
| `.ReplyTo(messageID int)` | no | Specifies the ID of the message to reply to |
| `.Do(ctx context.Context)` | — | Sends the photo. Returns `(*models.Message, error)` |

> *Exactly one of the three methods for specifying the photo source must be called, otherwise `.Do()` will return an error.

**Example:**
```go
// via FileID
msg, err := ctx.SendPhoto(userID).
    FileID("AgACAgIAAxkD...").
    Caption("Подпись").
    ParseMode(models.ModeHTML).
    ReplyTo(123).
    ReplyMarkup(inlineKeyboard).
    Do(ctx.Ctx)

// via FileURL
msg, err := ctx.SendPhoto(userID).
    FileURL("https://example.com/photo.jpg").
    Do(ctx.Ctx)

// via FilePath
msg, err := ctx.SendPhoto(userID).
    FilePath("/tmp/photo.jpg").
    Do(ctx.Ctx)
```

[Back to top](#api-reference)

---

### SendVideo

`ctx.SendVideo(userID int64)` -> `*VideoBuilder`

Creates a builder for sending a video.

| Method | Required | Description |
|---|---|---|
| `.FileID(id string)` | yes* | Sets the video by file_id |
| `.FileURL(url string)` | yes* | Sets the video by URL |
| `.FilePath(path string)` | yes* | Sets the video by local path (reads the file) |
| `.Caption(caption string)` | no | Sets the video caption |
| `.ParseMode(mode models.ParseMode)` | no | Sets the parsing mode for the caption |
| `.ReplyMarkup(kb models.ReplyMarkup)` | no | Adds a keyboard or inline markup |
| `.ReplyTo(messageID int)` | no | Specifies the ID of the message to reply to |
| `.Duration(seconds int)` | no | Sets the video duration in seconds |
| `.Size(width, height int)` | no | Sets the video dimensions (width, height) |
| `.Do(ctx context.Context)` | — | Sends the video. Returns `(*models.Message, error)` |

> *Exactly one of the three methods for specifying the video source must be called, otherwise `.Do()` will return an error.

**Example:**
```go
// via FileID
msg, err := ctx.SendVideo(userID).
    FileID("BAACAgIAAxkD...").
    Caption("Подпись").
    ParseMode(models.ModeHTML).
    Duration(60).
    Size(640, 480).
    ReplyTo(123).
    ReplyMarkup(inlineKeyboard).
    Do(ctx.Ctx)

// via FileURL
msg, err := ctx.SendVideo(userID).
    FileURL("https://example.com/video.mp4").
    Do(ctx.Ctx)

// via FilePath
msg, err := ctx.SendVideo(userID).
    FilePath("/tmp/video.mp4").
    Do(ctx.Ctx)
```

[Back to top](#api-reference)

---

### SendAudio

`ctx.SendAudio(userID int64)` -> `*AudioBuilder`

Creates a builder for sending audio.

| Method | Required | Description |
|---|---|---|
| `.FileID(id string)` | yes* | Sets the audio by file_id |
| `.FileURL(url string)` | yes* | Sets the audio by URL |
| `.FilePath(path string)` | yes* | Sets the audio by local path (reads the file) |
| `.Caption(caption string)` | no | Sets the audio caption |
| `.ParseMode(mode models.ParseMode)` | no | Sets the parsing mode for the caption |
| `.ReplyMarkup(kb models.ReplyMarkup)` | no | Adds a keyboard or inline markup |
| `.ReplyTo(messageID int)` | no | Specifies the ID of the message to reply to |
| `.Duration(seconds int)` | no | Sets the audio duration in seconds |
| `.Performer(performer string)` | no | Sets the performer name |
| `.Title(title string)` | no | Sets the track title |
| `.Do(ctx context.Context)` | — | Sends the audio. Returns `(*models.Message, error)` |

> *Exactly one of the three methods for specifying the audio source must be called, otherwise `.Do()` will return an error.

**Example:**
```go
// via FileID
msg, err := ctx.SendAudio(userID).
    FileID("CQACAgIAAxkD...").
    Caption("Подпись").
    ParseMode(models.ModeHTML).
    Duration(180).
    Performer("Artist").
    Title("Song").
    ReplyTo(123).
    ReplyMarkup(inlineKeyboard).
    Do(ctx.Ctx)

// via FileURL
msg, err := ctx.SendAudio(userID).
    FileURL("https://example.com/audio.mp3").
    Do(ctx.Ctx)

// via FilePath
msg, err := ctx.SendAudio(userID).
    FilePath("/tmp/audio.mp3").
    Do(ctx.Ctx)
```

[Back to top](#api-reference)

---

### SendDocument

`ctx.SendDocument(userID int64)` -> `*DocumentBuilder`

Creates a builder for sending a document.

| Method | Required | Description |
|---|---|---|
| `.FileID(id string)` | yes* | Sets the document by file_id |
| `.FileURL(url string)` | yes* | Sets the document by URL |
| `.FilePath(path string)` | yes* | Sets the document by local path (reads the file) |
| `.Caption(caption string)` | no | Sets the document caption |
| `.ParseMode(mode models.ParseMode)` | no | Sets the parsing mode for the caption |
| `.ReplyMarkup(kb models.ReplyMarkup)` | no | Adds a keyboard or inline markup |
| `.ReplyTo(messageID int)` | no | Specifies the ID of the message to reply to |
| `.Do(ctx context.Context)` | — | Sends the document. Returns `(*models.Message, error)` |

> *Exactly one of the three methods for specifying the document source must be called, otherwise `.Do()` will return an error.

**Example:**
```go
// via FileID
msg, err := ctx.SendDocument(userID).
    FileID("BQACAgIAAxkD...").
    Caption("Подпись").
    ParseMode(models.ModeHTML).
    ReplyTo(123).
    ReplyMarkup(inlineKeyboard).
    Do(ctx.Ctx)

// via FileURL
msg, err := ctx.SendDocument(userID).
    FileURL("https://example.com/doc.pdf").
    Do(ctx.Ctx)

// via FilePath
msg, err := ctx.SendDocument(userID).
    FilePath("/tmp/doc.pdf").
    Do(ctx.Ctx)
```

[Back to top](#api-reference)

---

### SendVoice

`ctx.SendVoice(userID int64)` -> `*VoiceBuilder`

Creates a builder for sending a voice message.

| Method | Required | Description |
|---|---|---|
| `.FileID(id string)` | yes* | Sets the voice message by file_id |
| `.FileURL(url string)` | yes* | Sets the voice message by URL |
| `.FilePath(path string)` | yes* | Sets the voice message by local path (reads the file) |
| `.Caption(caption string)` | no | Sets the voice message caption |
| `.ParseMode(mode models.ParseMode)` | no | Sets the parsing mode for the caption |
| `.ReplyMarkup(kb models.ReplyMarkup)` | no | Adds a keyboard or inline markup |
| `.ReplyTo(messageID int)` | no | Specifies the ID of the message to reply to |
| `.Duration(seconds int)` | no | Sets the voice message duration in seconds |
| `.Do(ctx context.Context)` | — | Sends the voice message. Returns `(*models.Message, error)` |

> *Exactly one of the three methods for specifying the voice source must be called, otherwise `.Do()` will return an error.

**Example:**
```go
// via FileID
msg, err := ctx.SendVoice(userID).
    FileID("AwACAgIAAxkD...").
    Caption("Подпись").
    ParseMode(models.ModeHTML).
    Duration(30).
    ReplyTo(123).
    ReplyMarkup(inlineKeyboard).
    Do(ctx.Ctx)

// via FileURL
msg, err := ctx.SendVoice(userID).
    FileURL("https://example.com/voice.ogg").
    Do(ctx.Ctx)

// via FilePath
msg, err := ctx.SendVoice(userID).
    FilePath("/tmp/voice.ogg").
    Do(ctx.Ctx)
```

[Back to top](#api-reference)

---

### SendVideoNote

`ctx.SendVideoNote(userID int64)` -> `*VideoNoteBuilder`

Creates a builder for sending a video note (round video).

| Method | Required | Description |
|---|---|---|
| `.FileID(id string)` | yes* | Sets the video note by file_id |
| `.FileURL(url string)` | yes* | Sets the video note by URL |
| `.FilePath(path string)` | yes* | Sets the video note by local path (reads the file) |
| `.ReplyMarkup(kb models.ReplyMarkup)` | no | Adds a keyboard or inline markup |
| `.ReplyTo(messageID int)` | no | Specifies the ID of the message to reply to |
| `.Duration(seconds int)` | no | Sets the video note duration in seconds |
| `.Length(pixels int)` | no | Sets the square side length in pixels |
| `.Do(ctx context.Context)` | — | Sends the video note. Returns `(*models.Message, error)` |

> *Exactly one of the three methods for specifying the video note source must be called, otherwise `.Do()` will return an error.

**Example:**
```go
// via FileID
msg, err := ctx.SendVideoNote(userID).
    FileID("DQACAgIAAxkD...").
    Duration(5).
    Length(240).
    ReplyTo(123).
    ReplyMarkup(inlineKeyboard).
    Do(ctx.Ctx)

// via FileURL
msg, err := ctx.SendVideoNote(userID).
    FileURL("https://example.com/videonote.mp4").
    Do(ctx.Ctx)

// via FilePath
msg, err := ctx.SendVideoNote(userID).
    FilePath("/tmp/videonote.mp4").
    Do(ctx.Ctx)
```

[Back to top](#api-reference)

</details>

---

## Editing

### EditText

`ctx.EditText(userID int64, messageID int, text string)` -> `*EditTextBuilder`

Creates a builder for editing an existing text message.

| Method | Required | Description |
|---|---|---|
| `.ParseMode(mode models.ParseMode)` | no | Sets the parsing mode (HTML, Markdown, etc.) |
| `.ReplyMarkup(kb models.ReplyMarkup)` | no | Updates the keyboard or inline markup |
| `.Do(ctx context.Context)` | — | Sends an edit request. Returns `(*models.Message, error)` |

**Example:**
```go
msg, err := ctx.EditText(userID, messageID, "Новый текст").
    ParseMode(models.ModeHTML).
    ReplyMarkup(newKeyboard).
    Do(ctx.Ctx)
```

[Back to top](#api-reference)

---

### EditCaption

`ctx.EditCaption(userID int64, messageID int, caption string)` -> `*EditCaptionBuilder`

Creates a builder for editing the caption of a media message.

| Method | Required | Description |
|---|---|---|
| `.ParseMode(mode models.ParseMode)` | no | Sets the parsing mode for the caption |
| `.ReplyMarkup(kb models.ReplyMarkup)` | no | Updates the keyboard or inline markup |
| `.Do(ctx context.Context)` | — | Sends an edit request for the caption. Returns `(*models.Message, error)` |

**Example:**
```go
msg, err := ctx.EditCaption(userID, messageID, "Новая подпись").
    ParseMode(models.ModeHTML).
    ReplyMarkup(newKeyboard).
    Do(ctx.Ctx)
```

[Back to top](#api-reference)

---

### EditMessageMedia

`ctx.EditMessageMedia(userID int64, messageID int)` -> `*EditMediaBuilder`

Creates a builder for replacing media in an existing message.

| Method | Required | Description |
|---|---|---|
| `.PhotoFileID(id, caption string)` | yes* | Sets the photo by file_id |
| `.PhotoURL(url, caption string)` | yes* | Sets the photo by URL |
| `.PhotoFromPath(path, caption string)` | yes* | Sets the photo by local path (reads the file) |
| `.VideoFileID(id, caption string)` | yes* | Sets the video by file_id |
| `.VideoURL(url, caption string)` | yes* | Sets the video by URL |
| `.VideoFromPath(path, caption string)` | yes* | Sets the video by local path (reads the file) |
| `.DocumentFileID(id, caption string)` | yes* | Sets the document by file_id |
| `.DocumentURL(url, caption string)` | yes* | Sets the document by URL |
| `.DocumentFromPath(path, caption string)` | yes* | Sets the document by local path (reads the file) |
| `.AudioFileID(id, caption string)` | yes* | Sets the audio by file_id |
| `.AudioURL(url, caption string)` | yes* | Sets the audio by URL |
| `.AudioFromPath(path, caption string)` | yes* | Sets the audio by local path (reads the file) |
| `.ReplyMarkup(kb models.ReplyMarkup)` | no | Updates the keyboard or inline markup |
| `.Do(ctx context.Context)` | — | Sends a request to replace the media. Returns `(*models.Message, error)` |

> *Exactly one media-setting method of any type with one of the three source methods must be called, otherwise `.Do()` will return an error.

**Example:**
```go
// via FileID (фото)
msg, err := ctx.EditMessageMedia(userID, messageID).
    PhotoFileID("AgACAgIAAxkD...", "Новая подпись").
    ReplyMarkup(newKeyboard).
    Do(ctx.Ctx)

// via URL (video)
msg, err := ctx.EditMessageMedia(userID, messageID).
    VideoURL("https://example.com/video.mp4", "Новое видео").
    Do(ctx.Ctx)

// via local path (document)
msg, err := ctx.EditMessageMedia(userID, messageID).
    DocumentFromPath("/tmp/doc.pdf", "Новый документ").
    Do(ctx.Ctx)
```

[Back to top](#api-reference)

---

## Deleting

### DeleteMessage

`ctx.DeleteMessage(userID int64, messageID int)` -> `*DeleteBuilder`

Creates a builder for deleting a message.

| Method | Required | Description |
|---|---|---|
| `.Do(ctx context.Context)` | — | Deletes the message. Returns `(bool, error)` |

**Example:**
```go
ok, err := ctx.DeleteMessage(userID, messageID).
    Do(ctx.Ctx)
```

[Back to top](#api-reference)

---

## Keyboards

### NewInlineKeyboard

`ctx.NewInlineKeyboard()` -> `*InlineKeyboardBuilder`

Creates a builder for constructing an inline keyboard.

| Method | Required | Description |
|---|---|---|
| `.Text(text, callbackData string)` | no | Adds a button with text and callback data |
| `.URL(text, url string)` | no | Adds a button with text and an HTTP link |
| `.Row()` | no | Finishes the current row of buttons and starts a new one |
| `.Build()` | — | Builds and returns `models.InlineKeyboardMarkup` |

**Example:**
```go
kb := ctx.NewInlineKeyboard().
    Text("Кнопка 1", "data1").
    Text("Кнопка 2", "data2").
    Row().
    URL("Перейти", "https://example.com").
    Build()

// Then kb can be passed to .ReplyMarkup(kb) of any builder
```

[Back to top](#api-reference)

---

### NewReplyKeyboard

`ctx.NewReplyKeyboard()` -> `*ReplyKeyboardBuilder`

Creates a builder for constructing a reply keyboard.

| Method | Required | Description |
|---|---|---|
| `.Text(text string)` | no | Adds a button with text |
| `.Row()` | no | Finishes the current row of buttons and starts a new one |
| `.Resize()` | no | Enables automatic keyboard resizing |
| `.OneTime()` | no | Makes the keyboard one-time (hidden after pressing) |
| `.Build()` | — | Builds and returns `models.ReplyKeyboardMarkup` |

**Example:**
```go
kb := ctx.NewReplyKeyboard().
    Text("Кнопка 1").
    Text("Кнопка 2").
    Row().
    Text("Кнопка 3").
    Resize().
    OneTime().
    Build()

// Then kb can be passed to .ReplyMarkup(kb) of any builder
```

[Back to top](#api-reference)

---

### RemoveKeyboard

`ctx.RemoveKeyboard()` -> `models.ReplyKeyboardRemove`

Returns an object for removing the current reply keyboard.

**Example:**
```go
// Passed to .ReplyMarkup to hide the keyboard
msg, err := ctx.SendText(userID, "Клавиатура скрыта").
    ReplyMarkup(ctx.RemoveKeyboard()).
    Do(ctx.Ctx)
```

[Back to top](#api-reference)

---

## Join requests

### ApproveJoin

`ctx.ApproveJoin(chatID, userID int64)` -> `*ApproveJoinBuilder`

Approves a chat join request.

| Method | Required | Description |
|---|---|---|
| `.Do(ctx context.Context)` | — | Sends the request. Returns `(bool, error)` |

**Example:**
```go
ctx.ApproveJoin(chatID, userID).Do(ctx.Ctx)
```

[Back to top](#api-reference)

---

### DeclineJoin

`ctx.DeclineJoin(chatID, userID int64)` -> `*DeclineJoinBuilder`

Declines a chat join request.

| Method | Required | Description |
|---|---|---|
| `.Do(ctx context.Context)` | — | Sends the request. Returns `(bool, error)` |

**Example:**
```go
ctx.DeclineJoin(chatID, userID).Do(ctx.Ctx)
```

[Back to top](#api-reference)

---

## Miscellaneous

### SendChatAction

`ctx.SendChatAction(userID int64, action models.ChatAction)` -> `*ChatActionBuilder`

Sends the bot's action status in the chat — for example, “typing…” or “sending a photo…”.

| Method | Required | Description |
|---|---|---|
| `.Do(ctx context.Context)` | — | Sends the request. Returns `(bool, error)` |

**Example:**
```go
ctx.SendChatAction(userID, models.ChatActionTyping).Do(ctx.Ctx)
```

The action status is visible to the user for a limited time (about 5 seconds), then disappears automatically.

[Back to top](#api-reference)

---

### AnswerCallback

`ctx.AnswerCallback()` -> `*AnswerCallbackBuilder`

Responds to an inline button press — removes the loading indicator from the button and optionally shows a popup notification.

| Method | Required | Description |
|---|---|---|
| `.Text(text string)` | no | Popup notification text |
| `.ShowAlert()` | no | Shows the text as a modal window instead of a small notification at the top |
| `.Do(ctx context.Context)` | — | Sends the request. Returns `(bool, error)` |

**Example:**
```go
// simply remove the loading indicator
ctx.AnswerCallback().Do(ctx.Ctx)

// with a notification text
ctx.AnswerCallback().Text("Готово!").Do(ctx.Ctx)

// with a modal window
ctx.AnswerCallback().Text("Ошибка!").ShowAlert().Do(ctx.Ctx)
```

[Back to top](#api-reference)

---

## Routing

### Command

`ctx.Command(cmd string, h Handler)`

Registers a handler for an exact match between the message text and the command.

**Example:**
```go
b.Command("/start", func(ctx *gogrammy.Context) {
    userID := c.Update.Message.From.ID
    ctx.SendText(userID, "Привет!").Do(ctx.Ctx)
})
```

[Back to top](#api-reference)

---

### On

`ctx.On(eventType string, h Handler)`

Registers a handler by incoming event (update) type. Supported `eventType` values: `message`, `edited_message`, `channel_post`, `edited_channel_post`, `callback`, `inline_query`, `chosen_inline_result`, `poll`, `poll_answer`, `my_chat_member`, `chat_member`, `join_request`, `chat_boost`, `removed_chat_boost` и другие.

**Example:**
```go
b.On("callback", func(ctx *gogrammy.Context) {
    data := ctx.Update.CallbackQuery.Data
    // ...
})
```

[Back to top](#api-reference)

---

### OnCallback

`ctx.OnCallback(prefix string, h Handler)`

Registers a handler for callback queries whose `callback_data` starts with the specified prefix.

**Example:**
```go
b.OnCallback("product_", func(ctx *gogrammy.Context) {
    data := ctx.Update.CallbackQuery.Data // например "product_42"
    // ...
})
```

[Back to top](#api-reference)

---

### Start

`ctx.Start(ctx context.Context)`

Starts the bot — begins receiving and processing incoming updates. This is a blocking call that runs until the provided `context.Context` is cancelled.

**Example:**
```go
b.Start(context.Background())
```

[Back to top](#api-reference)