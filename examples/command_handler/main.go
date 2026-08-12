package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/zzisler/gogrammy"
)

// Send /hi to the bot after it has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := gogrammy.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	b.Command("/hi", startHandle)

	b.Start(ctx)
}

func startHandle(c *gogrammy.Context) {
	userID := c.Update.Message.From.ID
	c.SendText(userID, "Hello!").Do(c.Ctx)
}
