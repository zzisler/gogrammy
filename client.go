package gogrammy

import (
	"net/http"
	"net/url"

	"github.com/go-telegram/bot"
	"github.com/zzisler/gogrammy/builder"
)

type Client struct {
	Bot      *bot.Bot
	handlers map[string]func(*Context)
}

type ClientOption func(*clientConfig)

type clientConfig struct {
	proxyURL string
}

func WithProxy(proxyURL string) ClientOption {
	return func(c *clientConfig) {
		c.proxyURL = proxyURL
	}
}

func New(token string, opts ...ClientOption) (*Client, error) {
	cfg := &clientConfig{}
	for _, opt := range opts {
		opt(cfg)
	}

	httpClient := http.DefaultClient
	if cfg.proxyURL != "" {
		proxy, err := url.Parse(cfg.proxyURL)
		if err != nil {
			return nil, err
		}
		httpClient = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxy),
			},
		}
	}

	b, err := bot.New(token, bot.WithHTTPClient(0, httpClient))
	if err != nil {
		return nil, err
	}

	return &Client{Bot: b, handlers: make(map[string]func(*Context))}, nil
}

func (c *Client) NewInlineKeyboard() *builder.InlineKeyboardBuilder {
	return builder.NewInlineKeyboard()
}

func (c *Client) NewReplyKeyboard() *builder.ReplyKeyboardBuilder {
	return builder.NewReplyKeyboard()
}

func (c *Client) SendText(userID int64, text string) *builder.TextBuilder {
	return builder.NewTextBuilder(c.Bot, userID, text)
}

func (c *Client) EditText(userID int64, messageID int, text string) *builder.EditTextBuilder {
	return builder.NewEditTextBuilder(c.Bot, userID, messageID, text)
}

func (c *Client) DeleteMessage(userID int64, messageID int) *builder.DeleteBuilder {
	return builder.NewDeleteBuilder(c.Bot, userID, messageID)
}
