package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/zzisler/gogrammy"
)

// Send /keyboard to the bot after it has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := gogrammy.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	b.Command("/keyboard", keyboardHandle)
	// callback query fires when the user taps an inline button
	b.On("callback", callbackHandle)

	b.Start(ctx)
}

func keyboardHandle(c *gogrammy.Context) {
	userID := c.Update.Message.From.ID

	// build an inline keyboard: two buttons in one row, then a URL button
	// on its own row (Row() closes the current row and starts a new one)
	kb := c.NewInlineKeyboard().
		Text("Option 1", "opt1").
		Text("Option 2", "opt2").
		Row().
		URL("Go docs", "https://go.dev").
		Build()

	c.SendText(userID, "Choose an option:").ReplyMarkup(kb).Do(c.Ctx)
}

func callbackHandle(c *gogrammy.Context) {
	userID := c.Update.CallbackQuery.From.ID
	data := c.Update.CallbackQuery.Data // the string we set via .Text(label, data)

	// IMPORTANT: always answer the callback query, otherwise the button
	// keeps showing a loading spinner on the user's side
	if _, err := c.AnswerCallback().Do(c.Ctx); err != nil {
		return
	}

	switch data {
	case "opt1":
		c.SendText(userID, "You picked Option 1").Do(c.Ctx)
	case "opt2":
		c.SendText(userID, "You picked Option 2").Do(c.Ctx)
	}
}
