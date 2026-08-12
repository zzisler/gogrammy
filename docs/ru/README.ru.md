<div align="center">

[English](../../README.md) · **Русский**

**Docs:** [API-справочник](API.ru.md) · [API Reference (EN)](../../API.md) · [Примеры](../../examples)

</div>

---

# gogrammy

Простая и понятная обёртка для Telegram-ботов на Go, вдохновлённая [Grammy](https://grammy.dev).

Обёртка над [go-telegram/bot](https://github.com/go-telegram/bot) с fluent builder-API вместо громоздких `&Param{}` структур.

```go
// было (go-telegram/bot напрямую)
b.SendMessage(ctx, &bot.SendMessageParams{
    ChatID:    userID,
    Text:      "Привет!",
    ParseMode: models.ParseModeHTML,
})

// стало (gogrammy)
ctx.SendText(userID, "Привет!").
    ParseMode(models.ParseModeHTML).
    Do(ctx.Ctx)
```

## Почему gogrammy

- **Билдеры вместо структур параметров.** Каждый тип контента (текст, фото, видео, документ, аудио, голосовое, кружочек) — свой билдер с цепочкой методов. Автокомплит в IDE сразу показывает, что применимо именно к этому типу.
- **Один тип `Context` на всё.** Не нужно помнить, где `Client`, а где `Context` — все методы отправки, редактирования, клавиатур висят в одном месте.
- **Роутинг по типу события.** `Command`, `On`, `OnCallback` — регистрация хендлеров без своей диспетчеризации, поверх `go-telegram/bot`.
- **Ничего лишнего.** gogrammy не переизобретает HTTP-клиент и работу с Telegram API — она использует стабильный `go-telegram/bot` под капотом и добавляет только слой удобства поверх.

## Установка

```bash
go get github.com/zzisler/gogrammy
```

## Быстрый старт

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
        c.SendText(userID, "Привет! Я бот на gogrammy 👋").Do(c.Ctx)
    })

    b.Start(context.Background())
}
```

Больше примеров — в папке [`examples/`](./../../examples).

## Что умеет

| Категория | Методы |
|---|---|
| Отправка | `SendText`, `SendPhoto`, `SendVideo`, `SendDocument`, `SendAudio`, `SendVoice`, `SendVideoNote` |
| Редактирование | `EditText`, `EditCaption`, `EditMedia` |
| Удаление | `DeleteMessage` |
| Клавиатуры | `NewInlineKeyboard`, `NewReplyKeyboard`, `RemoveKeyboard` |
| Заявки на вступление | `ApproveJoin`, `DeclineJoin` |
| Прочее | `SendChatAction`, `AnswerCallback` |
| Роутинг | `Command`, `On`, `OnCallback`, `Start` |

Полное описание каждого метода, всех опциональных параметров и примеры — в [`API.md`](./API.md).

## Работа с файлами

Каждый билдер медиа принимает источник тремя способами:

```go
ctx.SendPhoto(userID).FileID("AAA...")        // файл уже на серверах Telegram
ctx.SendPhoto(userID).FileURL("https://...")  // прямая ссылка
ctx.SendPhoto(userID).FilePath("./photo.png") // загрузка с диска
```

## Клавиатуры

Inline-клавиатура — привязана к конкретному сообщению, при нажатии присылает `callback_query`:

```go
kb := ctx.NewInlineKeyboard().
    Text("Кнопка 1", "btn1").
    Row().
    URL("Сайт", "https://example.com").
    Build()

ctx.SendText(userID, "Выбери:").ReplyMarkup(kb).Do(ctx.Ctx)
```

Reply-клавиатура — заменяет системную клавиатуру пользователя, нажатия приходят как обычные текстовые сообщения:

```go
kb := ctx.NewReplyKeyboard().
    Text("Меню").
    Row().
    Text("Помощь").
    Resize().
    Build()

ctx.SendText(userID, "Выбери:").ReplyMarkup(kb).Do(ctx.Ctx)
```

Снять reply-клавиатуру:

```go
ctx.SendText(userID, "Клавиатура убрана").ReplyMarkup(ctx.RemoveKeyboard()).Do(ctx.Ctx)
```

## Статус проекта

Библиотека активно разрабатывается. Основной функционал (отправка/редактирование/удаление всех типов контента, клавиатуры, роутинг, заявки на вступление) реализован и протестирован на живом боте. Обратная связь и issues приветствуются.

## Лицензия

MIT