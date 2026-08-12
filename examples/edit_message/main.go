package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/zzisler/gogrammy"
)

// Send /edit to the bot after it has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := gogrammy.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	b.Command("/edit", editHandle)

	b.Start(ctx)
}

func editHandle(c *gogrammy.Context) {
	userID := c.Update.Message.From.ID

	// send a message first — editing always needs a message_id to target
	msg, err := c.SendText(userID, "Original text").Do(c.Ctx)
	if err != nil {
		log.Printf("send failed: %v", err)
		return
	}

	time.Sleep(2 * time.Second) // just so the change is visible in this demo

	// edit that same message by its message_id
	if _, err := c.EditText(userID, msg.ID, "Edited text!").Do(c.Ctx); err != nil {
		log.Printf("edit failed: %v", err)
		return
	}

	// EditCaption works the same way, but only on messages that already
	// have media (photo/video/etc) — you can't turn text into a caption
	// or vice versa, Telegram doesn't allow that
}
