package telegraph

type (
	User struct {
		ID           int64  `json:"id"`
		IsBot        bool   `json:"is_bot"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name,omitempty"`
		Username     string `json:"username,omitempty"`
		LanguageCode string `json:"language_code,omitempty"`
	}

	// Update This object represents an incoming update.
	// At most one of the optional parameters can be present in any given update.
	Update struct {
		UpdateID           int64               `json:"update_id"`
		Message            *Message            `json:"message,omitempty"`
		EditedMessage      *Message            `json:"edited_message,omitempty"`
		ChannelPost        *Message            `json:"channel_post,omitempty"`
		EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
		InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
		ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
		CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
		ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
		pre
	}

	// InlineQuery This object represents an incoming inline query. When the user sends an empty query,
	// your bot could return some default or trending results.
	InlineQuery struct {
		ID       string    `json:"id"`
		From     User      `json:"from"`
		Location *Location `json:"location,omitempty"`
		Query    string    `json:"query"`
		Offset   string    `json:"offset"`
	}

	// Location This object represents a point on the map.
	Location struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
	}

	// ChosenInlineResult Represents a result of an inline query that was chosen by the user and sent to their chat partner.
	ChosenInlineResult struct {
		ResultID        string    `json:"result_id"`
		From            User      `json:"from"`
		Location        *Location `json:"location,omitempty"`
		InlineMessageID string    `json:"inline_message_id,omitempty"`
		Query           string    `json:"query"`
	}

	// CallbackQuery This object represents an incoming callback query from a callback button in an inline keyboard.
	// If the button that originated the query was attached to a message sent by the bot,
	// the field message will be present. If the button was attached to a message sent via the bot (in inline mode),
	// the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
	CallbackQuery struct {
		ID              string   `json:"id"`
		From            User     `json:"from"`
		Message         *Message `json:"message,omitempty"`
		InlineMessageID string   `json:"inline_message_id,omitempty"`
		ChatInstance    string   `json:"chat_instance"`
		Data            string   `json:"data,omitempty"`
		GameShortName   string   `json:"game_short_name,omitempty"`
	}

	// ShippingQuery This object contains information about an incoming shipping query.
	ShippingQuery struct {
		ID              string          `json:"id"`
		From            User            `json:"from"`
		InvoicePayload  string          `json:"invoice_payload"`
		ShippingAddress ShippingAddress `json:"shipping_address"`
	}

	// ShippingAddress This object represents a shipping address.
	ShippingAddress struct {
		CountryCode string `json:"country_code"`
		State       string `json:"state"`
		City        string `json:"city"`
		StreetLine1 string `json:"street_line1"`
		StreetLine2 string `json:"street_line2"`
		PostCode    string `json:"post_code"`
	}
)