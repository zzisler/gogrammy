<div align="center">

[English](../../README.md) · **Русский**

**Docs:** [API-справочник](API.ru.md) · [API Reference (EN)](../../API.md) · [Примеры](../../examples)

</div>

---

# gogrammy

Простая и понятная обёртка для Telegram-ботов на Go, вдохновлённая [Grammy](https://grammy.dev).

Обёртка над [go-telegram/bot](https://github.com/go-telegram/bot) с fluent builder-API вместо громоздких `&Param{}` структур.

<p>
    <svg width="15" height="15" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" style="vertical-align: middle; margin-right: 0px;">
        <path d="M9 12L11 14L15 10M12 3L13.9101 4.87147L16.5 4.20577L17.2184 6.78155L19.7942 7.5L19.1285 10.0899L21 12L19.1285 13.9101L19.7942 16.5L17.2184 17.2184L16.5 19.7942L13.9101 19.1285L12 21L10.0899 19.1285L7.5 19.7942L6.78155 17.2184L4.20577 16.5L4.87147 13.9101L3 12L4.87147 10.0899L4.20577 7.5L6.78155 6.78155L7.5 4.20577L10.0899 4.87147L12 3Z" stroke="#26A5E4" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
    </svg>
    <b>gogrammy</b> в официальном списке библиотек Telegram —
    <a href="https://core.telegram.org/bots/samples#go">core.telegram.org/bots/samples#go</a>
</p>

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