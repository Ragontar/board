package testdata

import (
	"api/models"
	"encoding/json"
	"io"
	"log"
	"os"
)

var RESPONSES = Responses{}

type Responses struct {
	RegAuthResponse                              RegAuthResponse                              `json:"RegAuthResponse"`
	ChannelListUserIdGetResponse                 ChannelListUserIdGetResponse                 `json:"ChannelListUserIdGetResponse"`
	MessagesGroupGroupIdGetResponse              MessagesGroupGroupIdGetResponse              `json:"MessagesGroupGroupIdGetResponse"`
	LinkTelegramUserIdPutResponse                LinkTelegramUserIdPutResponse                `json:"LinkTelegramUserIdPutResponse"`
	LinkTelegramUserIdConfirmPutResponse         LinkTelegramUserIdConfirmPutResponse         `json"LinkTelegramUserIdConfirmPutResponse"`
	GroupListUserIdCategoryCategoryIdPutResponse GroupListUserIdCategoryCategoryIdPutResponse `json"GroupListUserIdCategoryCategoryIdPutResponse"`
	GroupListUserIdCategoryPostResponse          GroupListUserIdCategoryPostResponse          `json:"GroupListUserIdCategoryPostResponse"`
	GroupListUserIdGetResponse                   GroupListUserIdGetResponse                   `json:"GroupListUserIdGetResponse"`
	GroupListUserIdGroupGroupIdPutResponse       GroupListUserIdGroupGroupIdPutResponse       `json:"GroupListUserIdGroupGroupIdPutResponse"`
	GroupListUserIdGroupPostResponse             GroupListUserIdGroupPostResponse             `json:"GroupListUserIdGroupPostResponse"`
}

type RegAuthResponse struct {
	AuthenticatedUser models.AuthenticatedUser
}

type ChannelListUserIdGetResponse struct {
	Channels []models.TelegramChannel `json:"channels"`
}

type MessagesGroupGroupIdGetResponse struct {
	Messages []models.TelegramMessage `json:"messages"`
}

type LinkTelegramUserIdPutResponse struct {
	TelegramConfirmationCode models.TelegramConfirmationCode `json"telegramConfirmationCode"`
}

type LinkTelegramUserIdConfirmPutResponse struct {
}

type GroupListUserIdCategoryCategoryIdPutResponse struct {
	GroupCategory models.GroupCategory `json:"groupCategory"`
}

type GroupListUserIdCategoryPostResponse struct {
	GroupCategory models.GroupCategory `json:"groupCategory"`
}

type GroupListUserIdGetResponse struct {
	ChannelGroups   []models.ChannelGroup  `json:"channelGroups"`
	GroupCategories []models.GroupCategory `json:"groupCategories"`
}

type GroupListUserIdGroupGroupIdPutResponse struct {
	ChannelGroup models.ChannelGroup `json:"channelGroup"`
}

type GroupListUserIdGroupPostResponse struct {
	ChannelGroup models.ChannelGroup `json:"channelGroup"`
}

func init() {
	reader, _ := os.Open("./testdata/example_responses.json")
	body, _ := io.ReadAll(reader)
	json.Unmarshal(body, &RESPONSES)

	log.Println(RESPONSES)
}
