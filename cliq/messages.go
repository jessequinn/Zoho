package cliq

import (
	"fmt"
	zoho "github.com/schmorrison/Zoho"
)

//func (c *API) PostMessageChat(request PostMessageChatData) (data PostMessageChatResponse, err error) {
//	endpoint := zoho.Endpoint{
//		Name:         "notes",
//		//URL:          fmt.Sprintf("https://www.zohoapis.%s/crm/v2/Notes", c.ZohoTLD),
//		URL:          fmt.Sprintf(" https://cliq.zoho.com/api/v2/channelsbyname/%s/message\n", c.ZohoTLD),
//		Method:       zoho.HTTPPost,
//		ResponseData: &PostMessageChatResponse{},
//		RequestBody:  request,
//	}
//
//	err = c.Zoho.HTTPRequest(&endpoint)
//	if err != nil {
//		return PostMessageChatResponse{}, fmt.Errorf("Failed to create notes: %s", err)
//	}
//
//	if v, ok := endpoint.ResponseData.(*PostMessageChatResponse); ok {
//		return *v, nil
//	}
//
//	return PostMessageChatResponse{}, fmt.Errorf("Data returned was not 'PostMessageChatResponse'")
//}

type PostMessageChatData struct {
	Text        string `json:"text"`
	SyncMessage bool   `json:"sync_message,omitempty"`
	ReplyTo     string `json:"reply_to,omitempty"`
}

type PostMessageChatResponse struct {
	MessageId string `json:"message_id"`
}

func (c *API) PostMessageChannel(request PostMessageChannelData) (data PostMessageChannelResponse, err error) {
	endpoint := zoho.Endpoint{
		Name:         "message",
		URL:          fmt.Sprintf("https://cliq.zoho.com/api/v2/channelsbyname/%s/message", c.channel),
		Method:       zoho.HTTPPost,
		ResponseData: &PostMessageChannelResponse{},
		RequestBody:  request,
	}

	err = c.Zoho.HTTPRequest(&endpoint)
	if err != nil {
		return PostMessageChannelResponse{}, fmt.Errorf("Failed to send message to channel: %s", err)
	}

	fmt.Println(endpoint)

	if v, ok := endpoint.ResponseData.(*PostMessageChannelResponse); ok {
		return *v, nil
	}

	return PostMessageChannelResponse{}, fmt.Errorf("Data returned was not 'PostMessageChatResponse'")
}

type PostMessageChannelData struct {
	Text        string `json:"text"`
	SyncMessage bool   `json:"sync_message,omitempty"`
	ReplyTo     string `json:"reply_to,omitempty"`
}

type PostMessageChannelResponse struct {
	MessageId string `json:"message_id"`
}
