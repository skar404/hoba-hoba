package bitly

type CreateLinkRes struct {
	CreatedAt      string        `json:"created_at"`
	ID             string        `json:"id"`
	Link           string        `json:"link"`
	CustomBitLinks []interface{} `json:"custom_bitlinks"`
	LongURL        string        `json:"long_url"`
	Archived       bool          `json:"archived"`
	Tags           []interface{} `json:"tags"`
	DeepLinks      []interface{} `json:"deeplinks"`
	References     References    `json:"references"`
}

type References struct {
	Group string `json:"group"`
}

type CreateLinkReq struct {
	GroupGuid string `json:"group_guid"`
	Domain    string `json:"domain"`
	LongUrl   string `json:"long_url"`
}
