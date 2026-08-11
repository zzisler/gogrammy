package builder

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// ==================== VOICE ======================

type VoiceBuilder struct {
	bot         *bot.Bot
	userID      int64
	voice       models.InputFile
	caption     string
	parseMode   models.ParseMode
	replyMarkup models.ReplyMarkup
	replyTo     int
	duration    int
	err         error
}

func NewVoiceBuilder(b *bot.Bot, userID int64) *VoiceBuilder {
	return &VoiceBuilder{bot: b, userID: userID}
}

func (b *VoiceBuilder) FileID(id string) *VoiceBuilder {
	b.voice = &models.InputFileString{Data: id}
	return b
}

func (b *VoiceBuilder) FileURL(url string) *VoiceBuilder {
	b.voice = &models.InputFileString{Data: url}
	return b
}

func (b *VoiceBuilder) FilePath(path string) *VoiceBuilder {
	data, err := os.ReadFile(path)
	if err != nil {
		b.err = err
		return b
	}
	b.voice = &models.InputFileUpload{Filename: filepath.Base(path), Data: bytes.NewReader(data)}
	return b
}

func (b *VoiceBuilder) Caption(caption string) *VoiceBuilder {
	b.caption = caption
	return b
}

func (b *VoiceBuilder) ParseMode(mode models.ParseMode) *VoiceBuilder {
	b.parseMode = mode
	return b
}

func (b *VoiceBuilder) ReplyMarkup(kb models.ReplyMarkup) *VoiceBuilder {
	b.replyMarkup = kb
	return b
}

func (b *VoiceBuilder) ReplyTo(messageID int) *VoiceBuilder {
	b.replyTo = messageID
	return b
}

func (b *VoiceBuilder) Duration(second int) *VoiceBuilder {
	b.duration = second
	return b
}

func (b *VoiceBuilder) Do(ctx context.Context) (*models.Message, error) {
	if b.err != nil {
		return nil, b.err
	}
	if b.voice == nil {
		return nil, fmt.Errorf("voice source is not set: call FileID, URL or FromPath")
	}

	params := &bot.SendVoiceParams{
		ChatID:      b.userID,
		Voice:       b.voice,
		Caption:     b.caption,
		ParseMode:   b.parseMode,
		ReplyMarkup: b.replyMarkup,
		Duration:    b.duration,
	}
	if b.replyTo != 0 {
		params.ReplyParameters = &models.ReplyParameters{MessageID: b.replyTo}
	}

	return b.bot.SendVoice(ctx, params)
}

// ==================== VIDEO NOTE ======================

type VideoNoteBuilder struct {
	bot         *bot.Bot
	userID      int64
	videoNote   models.InputFile
	replyMarkup models.ReplyMarkup
	replyTo     int
	duration    int
	length      int
	err         error
}

func NewVideoNoteBuilder(b *bot.Bot, userID int64) *VideoNoteBuilder {
	return &VideoNoteBuilder{bot: b, userID: userID}
}

func (b *VideoNoteBuilder) FileID(id string) *VideoNoteBuilder {
	b.videoNote = &models.InputFileString{Data: id}
	return b
}

func (b *VideoNoteBuilder) FileURL(url string) *VideoNoteBuilder {
	b.videoNote = &models.InputFileString{Data: url}
	return b
}

func (b *VideoNoteBuilder) FilePath(path string) *VideoNoteBuilder {
	data, err := os.ReadFile(path)
	if err != nil {
		b.err = err
		return b
	}
	b.videoNote = &models.InputFileUpload{Filename: filepath.Base(path), Data: bytes.NewReader(data)}
	return b
}

func (b *VideoNoteBuilder) ReplyMarkup(kb models.ReplyMarkup) *VideoNoteBuilder {
	b.replyMarkup = kb
	return b
}

func (b *VideoNoteBuilder) ReplyTo(messageID int) *VideoNoteBuilder {
	b.replyTo = messageID
	return b
}

func (b *VideoNoteBuilder) Duration(second int) *VideoNoteBuilder {
	b.duration = second
	return b
}

func (b *VideoNoteBuilder) Length(length int) *VideoNoteBuilder {
	b.length = length
	return b
}

func (b *VideoNoteBuilder) Do(ctx context.Context) (*models.Message, error) {
	if b.err != nil {
		return nil, b.err
	}
	if b.videoNote == nil {
		return nil, fmt.Errorf("videoNote source is not set: call FileID, URL or FromPath")
	}

	params := &bot.SendVideoNoteParams{
		ChatID:      b.userID,
		VideoNote:   b.videoNote,
		ReplyMarkup: b.replyMarkup,
		Duration:    b.duration,
		Length:      b.length,
	}
	if b.replyTo != 0 {
		params.ReplyParameters = &models.ReplyParameters{MessageID: b.replyTo}
	}

	return b.bot.SendVideoNote(ctx, params)
}
