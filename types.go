package gogrammy

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	HTML     = "HTML"
	Markdown = "Markdown"
)

type App struct {
	Bot      *bot.Bot
	Username string
}

type Context struct {
	Ctx    context.Context
	Bot    *bot.Bot
	Update *models.Update
}

type Handler func(c *Context)
