package telegraph_test

import (
	"fmt"
	"net/http"
	"telegraph"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestSendMessageForwardSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointForwardMessage, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 100,
			"from": {
				"id": 1234567890,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 1234567890,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510125931,
			"text": "test via server"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewForwardMessage("1233456", "1234567890", 1234567890)
	model, res, err := client.ForwardMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageForwardWithDisable(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointForwardMessage, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 100,
			"from": {
				"id": 1234567890,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 1234567890,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510125931,
			"text": "*test via server*"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewForwardMessage("1233456", "1232435435", 12323434).SetDisableNotification(true)
	model, res, err := client.ForwardMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageForwardError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointForwardMessage, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewForwardMessage("1233456", "test", 123234234)
	model, res, err := client.ForwardMessage(*message).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestSendMessageForwardFailedUnmarshal(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointForwardMessage, "token")).Reply(http.StatusBadRequest).XML("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewForwardMessage("1233456", "test", 123123213)
	model, res, err := client.ForwardMessage(*message).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendMessageForwardFailed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointForwardMessage, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewForwardMessage("1233456", "test", 234234234)
	model, res, err := client.ForwardMessage(*message).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendMessageTextSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 100,
			"from": {
				"id": 1234567890,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 1234567890,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510125931,
			"text": "test via server"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewTextMessage("1233456", "test")
	model, res, err := client.SendMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageTextWithMarkdown(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 100,
			"from": {
				"id": 1234567890,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 1234567890,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510125931,
			"text": "*test via server*"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewTextMessage("1233456", "test").SetParseMode(telegraph.ParseModeMarkdown)
	model, res, err := client.SendMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageTextWithHTML(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 100,
			"from": {
				"id": 1234567890,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 1234567890,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510125931,
			"text": "test via server"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewTextMessage("1233456", "test").SetParseMode(telegraph.ParseModeHTML)
	model, res, err := client.SendMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageTextDisableWebHook(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 100,
			"from": {
				"id": 1234567890,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 1234567890,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510125931,
			"text": "test via server"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewTextMessage("1233456", "test").SetDisableWebPagePreview(true)
	model, res, err := client.SendMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageTextDisableNotification(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 100,
			"from": {
				"id": 1234567890,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 1234567890,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510125931,
			"text": "test via server"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewTextMessage("1233456", "test").SetDisableNotification(true)
	model, res, err := client.SendMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageTextReplyId(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 100,
			"from": {
				"id": 1234567890,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 1234567890,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510125931,
			"text": "test via server"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewTextMessage("1233456", "test").SetReplyMessageToId(1234567890)
	model, res, err := client.SendMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageTextReplyMarkup(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 100,
			"from": {
				"id": 1234567890,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 1234567890,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510125931,
			"text": "test via server"
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	reply := telegraph.ForceReply{
		ForceReply: true,
	}
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewTextMessage("1233456", "test").SetForceReply(reply).
		SetInlineKeyboardMarkup(inline).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{})
	model, res, err := client.SendMessage(*message).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendMessageTextError(t *testing.T) {
	gock.New(telegraph.BaseURL).Head(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusInternalServerError).JSON("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewTextMessage("1233456", "test")
	model, res, err := client.SendMessage(*message).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Error(t, err)
}

func TestSendMessageTextFailedUnmarshal(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusBadRequest).XML("")
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewTextMessage("1233456", "test")
	model, res, err := client.SendMessage(*message).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendMessageTextFailed(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendMessage, "token")).Reply(http.StatusBadRequest).JSON(`{
		"ok": false,
		"error_code": 400,
		"description": "Bad Request: chat not found"
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewTextMessage("1233456", "test")
	model, res, err := client.SendMessage(*message).Commit()

	assert.Nil(t, model)
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	assert.Error(t, err)
}

func TestSendPhotoSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendPhoto, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 247,
			"from": {
				"id": 34234234,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 75092216,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510135752,
			"photo": [
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPLbkZTzbYSVvmAAAgI",
					"file_size": 1652,
					"width": 90,
					"height": 90
				},
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPGoi7BSr0V_vWAAAgI",
					"file_size": 3926,
					"width": 128,
					"height": 128
				}
			]
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewPhotoMessage("1233456", "http://www.cubesoft.com/image/test.jpg").SetCaption("test")
	model, res, err := client.SendPhoto(*message, false).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendPhotoDisableNotification(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendPhoto, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 247,
			"from": {
				"id": 34234234,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 75092216,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510135752,
			"photo": [
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPLbkZTzbYSVvmAAAgI",
					"file_size": 1652,
					"width": 90,
					"height": 90
				},
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPGoi7BSr0V_vWAAAgI",
					"file_size": 3926,
					"width": 128,
					"height": 128
				}
			]
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewPhotoMessage("1233456", "./LICENSE").SetDisableNotification(true)
	model, res, err := client.SendPhoto(*message, true).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendPhotoSetReplyToMessageId(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendPhoto, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 247,
			"from": {
				"id": 34234234,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 75092216,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510135752,
			"photo": [
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPLbkZTzbYSVvmAAAgI",
					"file_size": 1652,
					"width": 90,
					"height": 90
				},
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPGoi7BSr0V_vWAAAgI",
					"file_size": 3926,
					"width": 128,
					"height": 128
				}
			]
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewPhotoMessage("1233456", "./LICENSE").SetReplyToMessageId(342412342)
	model, res, err := client.SendPhoto(*message, true).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendPhotoReplyMarkup(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendPhoto, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 247,
			"from": {
				"id": 34234234,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 75092216,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510135752,
			"photo": [
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPLbkZTzbYSVvmAAAgI",
					"file_size": 1652,
					"width": 90,
					"height": 90
				},
				{
					"file_id": "AgADBAADcLs4G4AXZAdV7i1aVL3gsfjz4RkABPGoi7BSr0V_vWAAAgI",
					"file_size": 3926,
					"width": 128,
					"height": 128
				}
			]
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	reply := telegraph.ForceReply{
		ForceReply: true,
	}
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewPhotoMessage("1233456", "./LICENSE").SetForceReply(reply).
		SetInlineKeyboardMarkup(inline).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{})
	model, res, err := client.SendPhoto(*message, true).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendAudioSuccess(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendAudio, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 322343,
			"from": {
				"id": 234234324,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 34234234,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510279759,
			"audio": {
				"duration": 162,
				"mime_type": "audio/mpeg",
				"title": "test",
				"performer": "cube",
				"file_id": "NDNDJF949388JF30",
				"file_size": 2668544
			}
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	message := telegraph.NewAudioMessage("1233456", "http://www.cubesoft.com/audio/test.mp3").SetCaption("ok").
		SetDuration(1000).SetPerformer("Cube").SetTitle("soft").SetDisableNotification(true).
		SetReplyToMessageId(123332)
	model, res, err := client.SendAudio(*message, false).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}

func TestSendAudioReplyMarkup(t *testing.T) {
	gock.New(telegraph.BaseURL).Post(fmt.Sprintf(telegraph.EndpointSendAudio, "token")).Reply(http.StatusOK).JSON(`{
		"ok": true,
		"result": {
			"message_id": 322343,
			"from": {
				"id": 234234324,
				"is_bot": true,
				"first_name": "cube",
				"username": "cubesoft"
			},
			"chat": {
				"id": 34234234,
				"first_name": "cube",
				"last_name": "soft",
				"username": "cubesoft",
				"type": "private"
			},
			"date": 1510279759,
			"audio": {
				"duration": 162,
				"mime_type": "audio/mpeg",
				"title": "test",
				"performer": "cube",
				"file_id": "NDNDJF949388JF30",
				"file_size": 2668544
			}
		}
	}`)
	defer gock.Off()

	client := telegraph.NewClient("token")
	reply := telegraph.ForceReply{
		ForceReply: true,
	}
	inline := [][]telegraph.InlineKeyboardButton{}
	message := telegraph.NewAudioMessage("1233456", "./LICENSE").SetForceReply(reply).
		SetInlineKeyboardMarkup(inline).SetReplyKeyboardMarkup(telegraph.ReplyKeyboardMarkup{}).
		SetReplyKeyboardRemove(telegraph.ReplyKeyboardRemove{})
	model, res, err := client.SendAudio(*message, true).Commit()

	assert.NotNil(t, model)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NoError(t, err)
}