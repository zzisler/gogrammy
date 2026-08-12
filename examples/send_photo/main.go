package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/zzisler/gogrammy"
)

// Send /photo to the bot after it has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	b, err := gogrammy.New(os.Getenv("EXAMPLE_TELEGRAM_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	b.Command("/photo", photoHandle)

	b.Start(ctx)
}

func photoHandle(c *gogrammy.Context) {
	userID := c.Update.Message.From.ID

	// let's look at 3 ways to send a photo

	// method 1 — via a direct URL.
	// IMPORTANT: the URL must point directly to the file.
	// we save the message in msg, to reuse its file_id in the next method.
	msg, err := c.SendPhoto(userID).
		FileURL("https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png").
		Caption("URL preloaded Go logo").
		Do(c.Ctx)
	if err != nil {
		log.Printf("send by URL failed: %v", err)
		return
	}

	// grab the file_id from the message we just sent
	var fileID string
	if len(msg.Photo) > 0 {
		fileID = msg.Photo[len(msg.Photo)-1].FileID
	}

	// method 2 — via file_id.
	// every file_id is tied to the specific bot token that uploaded it,
	// so just copying someone else's file_id won't work.
	// here we reuse the file_id we got from uploading the photo through
	// YOUR bot in the previous step.
	c.SendPhoto(userID).
		FileID(fileID).
		Caption("FileID preloaded Go logo").
		Do(c.Ctx)

	// method 3 — upload from a local file.
	c.SendPhoto(userID).
		// path is relative to wherever you run the bot from.
		// I'm running it from the repo root, so my path looks like this:
		FilePath("./examples/send_photo/photo/go.jpg").
		Caption("Local preloaded Go logo").
		Do(c.Ctx)

	// as a result, we've sent 3 photos from 3 different sources
}
