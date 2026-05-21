package gogrammy

import (
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type BanParams struct {
	UntilDate int
	RevokeMsg bool
}

type ProfilePhotoParams struct {
	Offset int
	Limit  int
}

// ---------------------------BAN------------------------------

func (c *Context) Ban(chatID any, userID int64, params *BanParams) (bool, error) {

	p := &bot.BanChatMemberParams{
		ChatID: chatID,
		UserID: userID,
	}

	if params != nil {
		p.UntilDate = params.UntilDate
		p.RevokeMessages = params.RevokeMsg
	}

	isBan, err := c.Bot.BanChatMember(c.Ctx, p)
	if err != nil {
		log.Println("BanMember error:", err)
		return false, err
	}
	return isBan, nil
}

func (c *Context) BanChat(chatID any, senderChatID int64) (bool, error) {

	isBan, err := c.Bot.BanChatSenderChat(c.Ctx, &bot.BanChatSenderChatParams{
		ChatID:       chatID,
		SenderChatID: int(senderChatID),
	})
	if err != nil {
		log.Println("BanChat error:", err)
		return false, err
	}
	return isBan, nil
}

// ---------------------------UNBAN------------------------------

func (c *Context) Unban(chatID any, userID int64) (bool, error) {

	isUnban, err := c.Bot.UnbanChatMember(c.Ctx, &bot.UnbanChatMemberParams{
		ChatID:       chatID,
		UserID:       userID,
		OnlyIfBanned: true,
	})
	if err != nil {
		log.Println("Unban error:", err)
		return false, err
	}
	return isUnban, nil
}

func (c *Context) UnbanChat(chatID any, senderChatID int64) (bool, error) {

	isUnban, err := c.Bot.UnbanChatSenderChat(c.Ctx, &bot.UnbanChatSenderChatParams{
		ChatID:       chatID,
		SenderChatID: int(senderChatID),
	})
	if err != nil {
		log.Println("UnbanChat error:", err)
		return false, err
	}
	return isUnban, nil
}

// ---------------------------GET-USER-PHOTO------------------------------

func (c *Context) GetProfilePhoto(userID int64, params *ProfilePhotoParams) (*models.UserProfilePhotos, error) {

	p := &bot.GetUserProfilePhotosParams{
		UserID: userID,
	}

	if params != nil {
		p.Offset = params.Offset
		p.Limit = params.Limit
	}

	photos, err := c.Bot.GetUserProfilePhotos(c.Ctx, p)
	if err != nil {
		log.Println("GetProfilePhoto error:", err)
		return nil, err
	}
	return photos, nil
}

// ---------------------------APPROVE-JOIN-CHAT------------------------------

func (c *Context) ApproveJoin(chatID any, userID int64) (bool, error) {
	isApprove, err := c.Bot.ApproveChatJoinRequest(c.Ctx, &bot.ApproveChatJoinRequestParams{
		ChatID: chatID,
		UserID: userID,
	})
	if err != nil {
		log.Println("ApproveJoin error:", err)
		return false, err
	}
	return isApprove, nil
}

// ---------------------------DICLINE-JOIN-CHAT------------------------------

func (c *Context) DeclineJoin(chatID any, userID int64) (bool, error) {
	isDecline, err := c.Bot.DeclineChatJoinRequest(c.Ctx, &bot.DeclineChatJoinRequestParams{
		ChatID: chatID,
		UserID: userID,
	})
	if err != nil {
		log.Println("DeclineJoin error:", err)
		return false, err
	}
	return isDecline, nil
}
