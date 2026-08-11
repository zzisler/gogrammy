package gogrammy

import (
	"context"
	"net/http"
	"net/url"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/zzisler/gogrammy/builder"
)

type Context struct {
	Bot    *bot.Bot
	Update *models.Update
	Ctx    context.Context
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

func New(token string, opts ...ClientOption) (*Context, error) {
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

	return &Context{Bot: b}, nil
}

func (c *Context) UserID() int64 {
	return c.Update.Message.From.ID
}

func (c *Context) AnswerCallback() *builder.AnswerCallbackBuilder {
	return builder.NewAnswerCallbackBuilder(c.Bot, c.Update.CallbackQuery.ID)
}

func (c *Context) NewInlineKeyboard() *builder.InlineKeyboardBuilder {
	return builder.NewInlineKeyboard()
}

func (c *Context) NewReplyKeyboard() *builder.ReplyKeyboardBuilder {
	return builder.NewReplyKeyboard()
}

func (c *Context) RemoveKeyboard() models.ReplyKeyboardRemove {
	return models.ReplyKeyboardRemove{RemoveKeyboard: true}
}

func (c *Context) SendChatAction(userID int64, action models.ChatAction) *builder.ChatActionBuilder {
	return builder.NewChatActionBuilder(c.Bot, userID, action)
}

func (c *Context) SendText(userID int64, text string) *builder.TextBuilder {
	return builder.NewTextBuilder(c.Bot, userID, text)
}

func (c *Context) EditText(userID int64, messageID int, text string) *builder.EditTextBuilder {
	return builder.NewEditTextBuilder(c.Bot, userID, messageID, text)
}

func (c *Context) DeleteMessage(userID int64, messageID int) *builder.DeleteBuilder {
	return builder.NewDeleteBuilder(c.Bot, userID, messageID)
}

func (c *Context) SendPhoto(userID int64) *builder.PhotoBuilder {
	return builder.NewPhotoBuilder(c.Bot, userID)
}

func (c *Context) SendVideo(userID int64) *builder.VideoBuilder {
	return builder.NewVideoBuilder(c.Bot, userID)
}
