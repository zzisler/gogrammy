package gogrammy

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (a *App) handle(h Handler) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		c := &Context{Ctx: ctx, Bot: b, Update: update}
		h(c)
	}
}

func New(token string, proxy ...string) (*App, error) {
	opts := []bot.Option{}

	if len(proxy) > 0 && proxy[0] != "" {
		proxyURL, err := url.Parse(proxy[0])
		if err != nil {
			return nil, fmt.Errorf("Invalid proxy url: %w", err)
		}
		transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		httpClient := &http.Client{Transport: transport, Timeout: 30 * time.Second}
		opts = append(opts, bot.WithHTTPClient(30*time.Second, httpClient))
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		return nil, fmt.Errorf("Failed to create bot: %w", err)
	}

	me, err := b.GetMe(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Failed to get bot info: %w", err)
	}

	return &App{Bot: b, Username: me.Username}, nil
}

func (a *App) Start(ctx context.Context) {
	log.Println("Bot started!")
	a.Bot.Start(ctx)
}
