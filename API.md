<div align="center">

[English](./README.md) · [Русский](./README.ru.md)

**API Reference** · [Examples](./examples)

</div>

---

# API Reference

## Содержание
- [Отправка сообщений](#отправка-сообщений)
  - [SendText](#sendtext)
  - [SendPhoto](#sendphoto)
  - [SendVideo](#sendvideo)
  - [SendAudio](#sendaudio)
  - [SendDocument](#senddocument)
  - [SendVoice](#sendvoice)
  - [SendVideoNote](#sendvideonote)
- [Редактирование](#редактирование)
  - [EditText](#edittext)
  - [EditCaption](#editcaption)
  - [EditMessageMedia](#editmessagemedia)
- [Удаление](#удаление)
  - [DeleteMessage](#deletemessage)
- [Клавиатуры](#клавиатуры)
  - [NewInlineKeyboard](#newinlinekeyboard)
  - [NewReplyKeyboard](#newreplykeyboard)
  - [RemoveKeyboard](#removekeyboard)
- [Заявки на вступление](#заявки-на-вступление)
  - [ApproveJoin](#approvejoin)
  - [DeclineJoin](#declinejoin)
- [Прочее](#прочее)
  - [AnswerCallback](#answercallback)
  - [SendChatAction](#sendchataction)
- [Роутинг](#роутинг)
  - [Command](#command)
  - [On](#on)
  - [OnCallback](#oncallback)
  - [Start](#start)

---

## Отправка сообщений

### SendText

`ctx.SendText(userID int64, text string)` -> `*TextBuilder`

Создаёт билдер для отправки текстового сообщения.

| Метод | Обязательный | Описание |
|---|---|---|
| `.ParseMode(mode models.ParseMode)` | нет | Устанавливает режим парсинга (HTML, Markdown и т.п.) |
| `.ReplyTo(messageID int)` | нет | Указывает ID сообщения, на которое нужно ответить |
| `.ReplyMarkup(kb models.ReplyMarkup)` | нет | Добавляет клавиатуру или inline-разметку |
| `.Do(ctx context.Context)` | — | Отправляет сообщение. Возвращает `(*models.Message, error)` |

**Пример:**
```go
msg, err := ctx.SendText(userID, "Привет!").
    ParseMode(models.ModeHTML).
    ReplyTo(123).
    ReplyMarkup(inlineKeyboard).
    Do(ctx.Ctx)
```

[Наверх](#api-reference)

---

### SendPhoto

`ctx.SendPhoto(userID int64)` -> `*PhotoBuilder`

Создаёт билдер для отправки фотографии.

| Метод | Обязательный | Описание |
|---|---|---|
| `.FileID(id string)` | да* | Задаёт фото по file_id |
| `.FileURL(url string)` | да* | Задаёт фото по URL |
| `.FilePath(path string)` | да* | Задаёт фото по локальному пути (читает файл) |
| `.Caption(caption string)` | нет | Устанавливает подпись к фото |
| `.ParseMode(mode models.ParseMode)` | нет | Устанавливает режим парсинга для подписи |
| `.ReplyMarkup(kb models.ReplyMarkup)` | нет | Добавляет клавиатуру или inline-разметку |
| `.ReplyTo(messageID int)` | нет | Указывает ID сообщения, на которое нужно ответить |
| `.Do(ctx context.Context)` | — | Отправляет фото. Возвращает `(*models.Message, error)` |

> *Обязательно вызвать ровно один из трёх способов задания источника фото, иначе `.Do()` вернёт ошибку.

**Пример:**
```go
// через FileID
msg, err := ctx.SendPhoto(userID).
    FileID("AgACAgIAAxkD...").
    Caption("Подпись").
    ParseMode(models.ModeHTML).
    ReplyTo(123).
    ReplyMarkup(inlineKeyboard).
    Do(ctx.Ctx)

// через FileURL
msg, err := ctx.SendPhoto(userID).
    FileURL("https://example.com/photo.jpg").
    Do(ctx.Ctx)

// через FilePath
msg, err := ctx.SendPhoto(userID).
    FilePath("/tmp/photo.jpg").
    Do(ctx.Ctx)
```

[Наверх](#api-reference)

---

### SendVideo

`ctx.SendVideo(userID int64)` -> `*VideoBuilder`

Создаёт билдер для отправки видео.

| Метод | Обязательный | Описание |
|---|---|---|
| `.FileID(id string)` | да* | Задаёт видео по file_id |
| `.FileURL(url string)` | да* | Задаёт видео по URL |
| `.FilePath(path string)` | да* | Задаёт видео по локальному пути (читает файл) |
| `.Caption(caption string)` | нет | Устанавливает подпись к видео |
| `.ParseMode(mode models.ParseMode)` | нет | Устанавливает режим парсинга для подписи |
| `.ReplyMarkup(kb models.ReplyMarkup)` | нет | Добавляет клавиатуру или inline-разметку |
| `.ReplyTo(messageID int)` | нет | Указывает ID сообщения, на которое нужно ответить |
| `.Duration(seconds int)` | нет | Устанавливает длительность видео в секундах |
| `.Size(width, height int)` | нет | Устанавливает размеры видео (ширина, высота) |
| `.Do(ctx context.Context)` | — | Отправляет видео. Возвращает `(*models.Message, error)` |

> *Обязательно вызвать ровно один из трёх способов задания источника видео, иначе `.Do()` вернёт ошибку.

**Пример:**
```go
// через FileID
msg, err := ctx.SendVideo(userID).
    FileID("BAACAgIAAxkD...").
    Caption("Подпись").
    ParseMode(models.ModeHTML).
    Duration(60).
    Size(640, 480).
    ReplyTo(123).
    ReplyMarkup(inlineKeyboard).
    Do(ctx.Ctx)

// через FileURL
msg, err := ctx.SendVideo(userID).
    FileURL("https://example.com/video.mp4").
    Do(ctx.Ctx)

// через FilePath
msg, err := ctx.SendVideo(userID).
    FilePath("/tmp/video.mp4").
    Do(ctx.Ctx)
```

[Наверх](#api-reference)

---

### SendAudio

`ctx.SendAudio(userID int64)` -> `*AudioBuilder`

Создаёт билдер для отправки аудио.

| Метод | Обязательный | Описание |
|---|---|---|
| `.FileID(id string)` | да* | Задаёт аудио по file_id |
| `.FileURL(url string)` | да* | Задаёт аудио по URL |
| `.FilePath(path string)` | да* | Задаёт аудио по локальному пути (читает файл) |
| `.Caption(caption string)` | нет | Устанавливает подпись к аудио |
| `.ParseMode(mode models.ParseMode)` | нет | Устанавливает режим парсинга для подписи |
| `.ReplyMarkup(kb models.ReplyMarkup)` | нет | Добавляет клавиатуру или inline-разметку |
| `.ReplyTo(messageID int)` | нет | Указывает ID сообщения, на которое нужно ответить |
| `.Duration(seconds int)` | нет | Устанавливает длительность аудио в секундах |
| `.Performer(performer string)` | нет | Устанавливает имя исполнителя |
| `.Title(title string)` | нет | Устанавливает название трека |
| `.Do(ctx context.Context)` | — | Отправляет аудио. Возвращает `(*models.Message, error)` |

> *Обязательно вызвать ровно один из трёх способов задания источника аудио, иначе `.Do()` вернёт ошибку.

**Пример:**
```go
// через FileID
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

// через FileURL
msg, err := ctx.SendAudio(userID).
    FileURL("https://example.com/audio.mp3").
    Do(ctx.Ctx)

// через FilePath
msg, err := ctx.SendAudio(userID).
    FilePath("/tmp/audio.mp3").
    Do(ctx.Ctx)
```

[Наверх](#api-reference)

---

### SendDocument

`ctx.SendDocument(userID int64)` -> `*DocumentBuilder`

Создаёт билдер для отправки документа.

| Метод | Обязательный | Описание |
|---|---|---|
| `.FileID(id string)` | да* | Задаёт документ по file_id |
| `.FileURL(url string)` | да* | Задаёт документ по URL |
| `.FilePath(path string)` | да* | Задаёт документ по локальному пути (читает файл) |
| `.Caption(caption string)` | нет | Устанавливает подпись к документу |
| `.ParseMode(mode models.ParseMode)` | нет | Устанавливает режим парсинга для подписи |
| `.ReplyMarkup(kb models.ReplyMarkup)` | нет | Добавляет клавиатуру или inline-разметку |
| `.ReplyTo(messageID int)` | нет | Указывает ID сообщения, на которое нужно ответить |
| `.Do(ctx context.Context)` | — | Отправляет документ. Возвращает `(*models.Message, error)` |

> *Обязательно вызвать ровно один из трёх способов задания источника документа, иначе `.Do()` вернёт ошибку.

**Пример:**
```go
// через FileID
msg, err := ctx.SendDocument(userID).
    FileID("BQACAgIAAxkD...").
    Caption("Подпись").
    ParseMode(models.ModeHTML).
    ReplyTo(123).
    ReplyMarkup(inlineKeyboard).
    Do(ctx.Ctx)

// через FileURL
msg, err := ctx.SendDocument(userID).
    FileURL("https://example.com/doc.pdf").
    Do(ctx.Ctx)

// через FilePath
msg, err := ctx.SendDocument(userID).
    FilePath("/tmp/doc.pdf").
    Do(ctx.Ctx)
```

[Наверх](#api-reference)

---

### SendVoice

`ctx.SendVoice(userID int64)` -> `*VoiceBuilder`

Создаёт билдер для отправки голосового сообщения.

| Метод | Обязательный | Описание |
|---|---|---|
| `.FileID(id string)` | да* | Задаёт голосовое по file_id |
| `.FileURL(url string)` | да* | Задаёт голосовое по URL |
| `.FilePath(path string)` | да* | Задаёт голосовое по локальному пути (читает файл) |
| `.Caption(caption string)` | нет | Устанавливает подпись к голосовому |
| `.ParseMode(mode models.ParseMode)` | нет | Устанавливает режим парсинга для подписи |
| `.ReplyMarkup(kb models.ReplyMarkup)` | нет | Добавляет клавиатуру или inline-разметку |
| `.ReplyTo(messageID int)` | нет | Указывает ID сообщения, на которое нужно ответить |
| `.Duration(seconds int)` | нет | Устанавливает длительность голосового в секундах |
| `.Do(ctx context.Context)` | — | Отправляет голосовое. Возвращает `(*models.Message, error)` |

> *Обязательно вызвать ровно один из трёх способов задания источника голосового, иначе `.Do()` вернёт ошибку.

**Пример:**
```go
// через FileID
msg, err := ctx.SendVoice(userID).
    FileID("AwACAgIAAxkD...").
    Caption("Подпись").
    ParseMode(models.ModeHTML).
    Duration(30).
    ReplyTo(123).
    ReplyMarkup(inlineKeyboard).
    Do(ctx.Ctx)

// через FileURL
msg, err := ctx.SendVoice(userID).
    FileURL("https://example.com/voice.ogg").
    Do(ctx.Ctx)

// через FilePath
msg, err := ctx.SendVoice(userID).
    FilePath("/tmp/voice.ogg").
    Do(ctx.Ctx)
```

[Наверх](#api-reference)

---

### SendVideoNote

`ctx.SendVideoNote(userID int64)` -> `*VideoNoteBuilder`

Создаёт билдер для отправки видео-заметки (кружок).

| Метод | Обязательный | Описание |
|---|---|---|
| `.FileID(id string)` | да* | Задаёт видео-заметку по file_id |
| `.FileURL(url string)` | да* | Задаёт видео-заметку по URL |
| `.FilePath(path string)` | да* | Задаёт видео-заметку по локальному пути (читает файл) |
| `.ReplyMarkup(kb models.ReplyMarkup)` | нет | Добавляет клавиатуру или inline-разметку |
| `.ReplyTo(messageID int)` | нет | Указывает ID сообщения, на которое нужно ответить |
| `.Duration(seconds int)` | нет | Устанавливает длительность видео-заметки в секундах |
| `.Length(pixels int)` | нет | Устанавливает размер стороны квадрата в пикселях |
| `.Do(ctx context.Context)` | — | Отправляет видео-заметку. Возвращает `(*models.Message, error)` |

> *Обязательно вызвать ровно один из трёх способов задания источника видео-заметки, иначе `.Do()` вернёт ошибку.

**Пример:**
```go
// через FileID
msg, err := ctx.SendVideoNote(userID).
    FileID("DQACAgIAAxkD...").
    Duration(5).
    Length(240).
    ReplyTo(123).
    ReplyMarkup(inlineKeyboard).
    Do(ctx.Ctx)

// через FileURL
msg, err := ctx.SendVideoNote(userID).
    FileURL("https://example.com/videonote.mp4").
    Do(ctx.Ctx)

// через FilePath
msg, err := ctx.SendVideoNote(userID).
    FilePath("/tmp/videonote.mp4").
    Do(ctx.Ctx)
```

[Наверх](#api-reference)

</details>

---

## Редактирование

### EditText

`ctx.EditText(userID int64, messageID int, text string)` -> `*EditTextBuilder`

Создаёт билдер для редактирования существующего текстового сообщения.

| Метод | Обязательный | Описание |
|---|---|---|
| `.ParseMode(mode models.ParseMode)` | нет | Устанавливает режим парсинга (HTML, Markdown и т.п.) |
| `.ReplyMarkup(kb models.ReplyMarkup)` | нет | Обновляет клавиатуру или inline-разметку |
| `.Do(ctx context.Context)` | — | Отправляет запрос на редактирование. Возвращает `(*models.Message, error)` |

**Пример:**
```go
msg, err := ctx.EditText(userID, messageID, "Новый текст").
    ParseMode(models.ModeHTML).
    ReplyMarkup(newKeyboard).
    Do(ctx.Ctx)
```

[Наверх](#api-reference)

---

### EditCaption

`ctx.EditCaption(userID int64, messageID int, caption string)` -> `*EditCaptionBuilder`

Создаёт билдер для редактирования подписи у медиа-сообщения.

| Метод | Обязательный | Описание |
|---|---|---|
| `.ParseMode(mode models.ParseMode)` | нет | Устанавливает режим парсинга для подписи |
| `.ReplyMarkup(kb models.ReplyMarkup)` | нет | Обновляет клавиатуру или inline-разметку |
| `.Do(ctx context.Context)` | — | Отправляет запрос на редактирование подписи. Возвращает `(*models.Message, error)` |

**Пример:**
```go
msg, err := ctx.EditCaption(userID, messageID, "Новая подпись").
    ParseMode(models.ModeHTML).
    ReplyMarkup(newKeyboard).
    Do(ctx.Ctx)
```

[Наверх](#api-reference)

---

### EditMessageMedia

`ctx.EditMessageMedia(userID int64, messageID int)` -> `*EditMediaBuilder`

Создаёт билдер для замены медиа в существующем сообщении.

| Метод | Обязательный | Описание |
|---|---|---|
| `.PhotoFileID(id, caption string)` | да* | Задаёт фото по file_id |
| `.PhotoURL(url, caption string)` | да* | Задаёт фото по URL |
| `.PhotoFromPath(path, caption string)` | да* | Задаёт фото по локальному пути (читает файл) |
| `.VideoFileID(id, caption string)` | да* | Задаёт видео по file_id |
| `.VideoURL(url, caption string)` | да* | Задаёт видео по URL |
| `.VideoFromPath(path, caption string)` | да* | Задаёт видео по локальному пути (читает файл) |
| `.DocumentFileID(id, caption string)` | да* | Задаёт документ по file_id |
| `.DocumentURL(url, caption string)` | да* | Задаёт документ по URL |
| `.DocumentFromPath(path, caption string)` | да* | Задаёт документ по локальному пути (читает файл) |
| `.AudioFileID(id, caption string)` | да* | Задаёт аудио по file_id |
| `.AudioURL(url, caption string)` | да* | Задаёт аудио по URL |
| `.AudioFromPath(path, caption string)` | да* | Задаёт аудио по локальному пути (читает файл) |
| `.ReplyMarkup(kb models.ReplyMarkup)` | нет | Обновляет клавиатуру или inline-разметку |
| `.Do(ctx context.Context)` | — | Отправляет запрос на замену медиа. Возвращает `(*models.Message, error)` |

> *Обязательно вызвать ровно один метод установки медиа любого типа с одним из трёх способов задания источника, иначе `.Do()` вернёт ошибку.

**Пример:**
```go
// через FileID (фото)
msg, err := ctx.EditMessageMedia(userID, messageID).
    PhotoFileID("AgACAgIAAxkD...", "Новая подпись").
    ReplyMarkup(newKeyboard).
    Do(ctx.Ctx)

// через URL (видео)
msg, err := ctx.EditMessageMedia(userID, messageID).
    VideoURL("https://example.com/video.mp4", "Новое видео").
    Do(ctx.Ctx)

// через локальный путь (документ)
msg, err := ctx.EditMessageMedia(userID, messageID).
    DocumentFromPath("/tmp/doc.pdf", "Новый документ").
    Do(ctx.Ctx)
```

[Наверх](#api-reference)

---

## Удаление

### DeleteMessage

`ctx.DeleteMessage(userID int64, messageID int)` -> `*DeleteBuilder`

Создаёт билдер для удаления сообщения.

| Метод | Обязательный | Описание |
|---|---|---|
| `.Do(ctx context.Context)` | — | Удаляет сообщение. Возвращает `(bool, error)` |

**Пример:**
```go
ok, err := ctx.DeleteMessage(userID, messageID).
    Do(ctx.Ctx)
```

[Наверх](#api-reference)

---

## Клавиатуры

### NewInlineKeyboard

`ctx.NewInlineKeyboard()` -> `*InlineKeyboardBuilder`

Создаёт билдер для построения inline-клавиатуры.

| Метод | Обязательный | Описание |
|---|---|---|
| `.Text(text, callbackData string)` | нет | Добавляет кнопку с текстом и callback-данными |
| `.URL(text, url string)` | нет | Добавляет кнопку с текстом и HTTP-ссылкой |
| `.Row()` | нет | Завершает текущий ряд кнопок и начинает новый |
| `.Build()` | — | Собирает и возвращает `models.InlineKeyboardMarkup` |

**Пример:**
```go
kb := ctx.NewInlineKeyboard().
    Text("Кнопка 1", "data1").
    Text("Кнопка 2", "data2").
    Row().
    URL("Перейти", "https://example.com").
    Build()

// Затем kb можно передать в .ReplyMarkup(kb) любого билдера
```

[Наверх](#api-reference)

---

### NewReplyKeyboard

`ctx.NewReplyKeyboard()` -> `*ReplyKeyboardBuilder`

Создаёт билдер для построения reply-клавиатуры.

| Метод | Обязательный | Описание |
|---|---|---|
| `.Text(text string)` | нет | Добавляет кнопку с текстом |
| `.Row()` | нет | Завершает текущий ряд кнопок и начинает новый |
| `.Resize()` | нет | Включает автоматическое изменение размера клавиатуры |
| `.OneTime()` | нет | Делает клавиатуру одноразовой (скрывается после нажатия) |
| `.Build()` | — | Собирает и возвращает `models.ReplyKeyboardMarkup` |

**Пример:**
```go
kb := ctx.NewReplyKeyboard().
    Text("Кнопка 1").
    Text("Кнопка 2").
    Row().
    Text("Кнопка 3").
    Resize().
    OneTime().
    Build()

// Затем kb можно передать в .ReplyMarkup(kb) любого билдера
```

[Наверх](#api-reference)

---

### RemoveKeyboard

`ctx.RemoveKeyboard()` -> `models.ReplyKeyboardRemove`

Возвращает объект для удаления текущей reply-клавиатуры.

**Пример:**
```go
// Передаётся в .ReplyMarkup для скрытия клавиатуры
msg, err := ctx.SendText(userID, "Клавиатура скрыта").
    ReplyMarkup(ctx.RemoveKeyboard()).
    Do(ctx.Ctx)
```

[Наверх](#api-reference)

---

## Заявки на вступление

### ApproveJoin

`ctx.ApproveJoin(chatID, userID int64)` -> `*ApproveJoinBuilder`

Одобряет заявку на вступление в чат.

| Метод | Обязательный | Описание |
|---|---|---|
| `.Do(ctx context.Context)` | — | Отправить. Возвращает `(bool, error)` |

**Пример:**
```go
ctx.ApproveJoin(chatID, userID).Do(ctx.Ctx)
```

[Наверх](#api-reference)

---

### DeclineJoin

`ctx.DeclineJoin(chatID, userID int64)` -> `*DeclineJoinBuilder`

Отклоняет заявку на вступление в чат.

| Метод | Обязательный | Описание |
|---|---|---|
| `.Do(ctx context.Context)` | — | Отправить. Возвращает `(bool, error)` |

**Пример:**
```go
ctx.DeclineJoin(chatID, userID).Do(ctx.Ctx)
```

[Наверх](#api-reference)

---

## Прочее

### SendChatAction

`ctx.SendChatAction(userID int64, action models.ChatAction)` -> `*ChatActionBuilder`

Отправляет статус действия бота в чате — например «печатает…» или «отправляет фото…».

| Метод | Обязательный | Описание |
|---|---|---|
| `.Do(ctx context.Context)` | — | Отправить. Возвращает `(bool, error)` |

**Пример:**
```go
ctx.SendChatAction(userID, models.ChatActionTyping).Do(ctx.Ctx)
```

Статус действия виден пользователю ограниченное время (около 5 секунд), затем гаснет сам.

[Наверх](#api-reference)

---

### AnswerCallback

`ctx.AnswerCallback()` -> `*AnswerCallbackBuilder`

Отвечает на нажатие inline-кнопки — убирает индикатор загрузки у кнопки, опционально показывает всплывающее уведомление.

| Метод | Обязательный | Описание |
|---|---|---|
| `.Text(text string)` | нет | Текст всплывающего уведомления |
| `.ShowAlert()` | нет | Показать текст как модальное окно, а не мелкое уведомление сверху |
| `.Do(ctx context.Context)` | — | Отправить. Возвращает `(bool, error)` |

**Пример:**
```go
// просто убрать индикатор загрузки
ctx.AnswerCallback().Do(ctx.Ctx)

// с текстом-уведомлением
ctx.AnswerCallback().Text("Готово!").Do(ctx.Ctx)

// с модальным окном
ctx.AnswerCallback().Text("Ошибка!").ShowAlert().Do(ctx.Ctx)
```

[Наверх](#api-reference)

---

## Роутинг

### Command

`ctx.Command(cmd string, h Handler)`

Регистрирует хендлер на точное совпадение текста сообщения с командой.

**Пример:**
```go
b.Command("/start", func(ctx *gogrammy.Context) {
    userID := c.Update.Message.From.ID
    ctx.SendText(userID, "Привет!").Do(ctx.Ctx)
})
```

[Наверх](#api-reference)

---

### On

`ctx.On(eventType string, h Handler)`

Регистрирует хендлер по типу входящего события (апдейта). Поддерживаемые `eventType`: `message`, `edited_message`, `channel_post`, `edited_channel_post`, `callback`, `inline_query`, `chosen_inline_result`, `poll`, `poll_answer`, `my_chat_member`, `chat_member`, `join_request`, `chat_boost`, `removed_chat_boost` и другие.

**Пример:**
```go
b.On("callback", func(ctx *gogrammy.Context) {
    data := ctx.Update.CallbackQuery.Data
    // ...
})
```

[Наверх](#api-reference)

---

### OnCallback

`ctx.OnCallback(prefix string, h Handler)`

Регистрирует хендлер на callback-запросы, чей `callback_data` начинается с указанного префикса.

**Пример:**
```go
b.OnCallback("product_", func(ctx *gogrammy.Context) {
    data := ctx.Update.CallbackQuery.Data // например "product_42"
    // ...
})
```

[Наверх](#api-reference)

---

### Start

`ctx.Start(ctx context.Context)`

Запускает бота — начинает получать и обрабатывать входящие апдейты. Блокирующий вызов, выполняется до отмены переданного `context.Context`.

**Пример:**
```go
b.Start(context.Background())
```

[Наверх](#api-reference)