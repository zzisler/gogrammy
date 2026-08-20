package gogrammy

import (
	"context"
	"net/http"
	"net/url"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/zzisler/gogrammy/builder"
)

type Bot struct {
	api *bot.Bot
}

type Context struct {
	bot    *bot.Bot
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

func New(token string, opts ...ClientOption) (*Bot, error) {
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

	return &Bot{api: b}, nil
}

// --HELP-UPD--

func (c *Context) UserID() int64 {
	switch {
	// --message--
	case c.Update.Message != nil:
		return c.Update.Message.From.ID

	// --edited message--
	case c.Update.EditedMessage != nil:
		return c.Update.EditedMessage.From.ID

	// // --channel post--
	// case c.Update.ChannelPost != nil:
	// 	return c.Update.ChannelPost.From.ID

	// // --edited channel post--
	// case c.Update.EditedChannelPost != nil:
	// 	return c.Update.EditedChannelPost.From.ID

	// --business message--
	case c.Update.BusinessMessage != nil:
		return c.Update.BusinessMessage.From.ID

	// --edited business message--
	case c.Update.EditedBusinessMessage != nil:
		return c.Update.EditedBusinessMessage.From.ID

	// --message reaction--
	case c.Update.MessageReaction != nil:
		return c.Update.MessageReaction.User.ID

	// --inline query--
	case c.Update.InlineQuery != nil:
		return c.Update.InlineQuery.From.ID

	// --chosen inline result--
	case c.Update.ChosenInlineResult != nil:
		return c.Update.ChosenInlineResult.From.ID

	// --callback query--
	case c.Update.CallbackQuery != nil:
		return c.Update.CallbackQuery.From.ID

	// --shipping query--
	case c.Update.ShippingQuery != nil:
		return c.Update.ShippingQuery.From.ID

	// --pre checkout query--
	case c.Update.PreCheckoutQuery != nil:
		return c.Update.PreCheckoutQuery.From.ID

	// --purchased paid media--
	case c.Update.PurchasedPaidMedia != nil:
		return c.Update.PurchasedPaidMedia.From.ID

	// --managed bot--
	case c.Update.ManagedBot != nil:
		return c.Update.ManagedBot.User.ID

	// --guest message--
	case c.Update.GuestMessage != nil:
		return c.Update.GuestMessage.From.ID

	// --my chat member--
	case c.Update.MyChatMember != nil:
		return c.Update.MyChatMember.From.ID

	// --chat member--
	case c.Update.ChatMember != nil:
		return c.Update.ChatMember.From.ID

	// --chat join request--
	case c.Update.ChatJoinRequest != nil:
		return c.Update.ChatJoinRequest.From.ID

	// --subscription--
	case c.Update.Subscription != nil:
		return c.Update.Subscription.User.ID

	default:
		return 0
	}
}

func (c *Context) ChatID() int64 {
	switch {
	// --message--
	case c.Update.Message != nil:
		return c.Update.Message.Chat.ID

	// --edited message--
	case c.Update.EditedMessage != nil:
		return c.Update.EditedMessage.Chat.ID

	// --business message--
	case c.Update.BusinessMessage != nil:
		return c.Update.BusinessMessage.Chat.ID

	// --edited business message--
	case c.Update.EditedBusinessMessage != nil:
		return c.Update.EditedBusinessMessage.Chat.ID

	// --deleted business message--
	case c.Update.DeletedBusinessMessages != nil:
		return c.Update.DeletedBusinessMessages.Chat.ID

	// --business connection--
	case c.Update.BusinessConnection != nil:
		return c.Update.BusinessConnection.UserChatID

	// --message reaction--
	case c.Update.MessageReaction != nil:
		return c.Update.MessageReaction.Chat.ID

	// --message reaction count--
	case c.Update.MessageReactionCount != nil:
		return c.Update.MessageReactionCount.Chat.ID

	// --callback query--
	case c.Update.CallbackQuery != nil:
		return c.Update.CallbackQuery.Message.Message.Chat.ID

	// --guest message--
	case c.Update.GuestMessage != nil:
		return c.Update.GuestMessage.Chat.ID

	// --my chat member--
	case c.Update.MyChatMember != nil:
		return c.Update.MyChatMember.Chat.ID

	// --chat member--
	case c.Update.ChatMember != nil:
		return c.Update.ChatMember.Chat.ID

	// --chat join request--
	case c.Update.ChatJoinRequest != nil:
		return c.Update.ChatJoinRequest.Chat.ID

	default:
		return 0
	}
}

// --API--

func (c *Context) API() *bot.Bot {
	return c.bot
}

// --SEND--

func (c *Context) SendText(userID int64, text string) *builder.TextBuilder {
	return builder.NewTextBuilder(c.bot, userID, text)
}

func (c *Context) SendChatAction(userID int64, action models.ChatAction) *builder.ChatActionBuilder {
	return builder.NewChatActionBuilder(c.bot, userID, action)
}

func (c *Context) DeleteMessage(userID int64, messageID int) *builder.DeleteBuilder {
	return builder.NewDeleteBuilder(c.bot, userID, messageID)
}

func (c *Context) SendPhoto(userID int64) *builder.PhotoBuilder {
	return builder.NewPhotoBuilder(c.bot, userID)
}

func (c *Context) SendVideo(userID int64) *builder.VideoBuilder {
	return builder.NewVideoBuilder(c.bot, userID)
}

func (c *Context) SendAudio(userID int64) *builder.AudioBuilder {
	return builder.NewAudioBuilder(c.bot, userID)
}

func (c *Context) SendDocument(userID int64) *builder.DocumentBuilder {
	return builder.NewDocumentBuilder(c.bot, userID)
}

func (c *Context) SendVoice(userID int64) *builder.VoiceBuilder {
	return builder.NewVoiceBuilder(c.bot, userID)
}

func (c *Context) SendVideoNote(userID int64) *builder.VideoNoteBuilder {
	return builder.NewVideoNoteBuilder(c.bot, userID)
}

// --EDIT--

func (c *Context) EditText(userID int64, messageID int, text string) *builder.EditTextBuilder {
	return builder.NewEditTextBuilder(c.bot, userID, messageID, text)
}

func (c *Context) EditCaption(userID int64, messageID int, caption string) *builder.EditCaptionBuilder {
	return builder.NewEditCaptionBuilder(c.bot, userID, messageID, caption)
}

func (c *Context) EditMessageMedia(userID int64, messageID int) *builder.EditMediaBuilder {
	return builder.NewEditMediaBuilder(c.bot, userID, messageID)
}

// --OTHER--

func (c *Context) AnswerCallback() *builder.AnswerCallbackBuilder {
	return builder.NewAnswerCallbackBuilder(c.bot, c.Update.CallbackQuery.ID)
}

// --KEYBOARD--

func (c *Context) NewInlineKeyboard() *builder.InlineKeyboardBuilder {
	return builder.NewInlineKeyboard()
}

func (c *Context) NewReplyKeyboard() *builder.ReplyKeyboardBuilder {
	return builder.NewReplyKeyboard()
}

func (c *Context) RemoveKeyboard() models.ReplyKeyboardRemove {
	return models.ReplyKeyboardRemove{RemoveKeyboard: true}
}

// --JOIN--

func (c *Context) ApproveJoin(chatID, userID int64) *builder.ApproveJoinBuilder {
	return builder.NewApproveJoinBuilder(c.bot, chatID, userID)
}

func (c *Context) DeclineJoin(chatID, userID int64) *builder.DeclineJoinBuilder {
	return builder.NewDeclineJoinBuilder(c.bot, chatID, userID)
}
