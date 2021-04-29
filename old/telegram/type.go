package telegram

type SendAudioRes struct {
	Ok     bool          `json:"ok"`
	Result AudioResponse `json:"result"`
}

type SenderChat struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Username string `json:"username"`
	Type     string `json:"type"`
}

type Chat struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Username string `json:"username"`
	Type     string `json:"type"`
}

type Audio struct {
	Duration     int    `json:"duration"`
	FileName     string `json:"file_name"`
	MimeType     string `json:"mime_type"`
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
}

type AudioResponse struct {
	MessageID  int        `json:"message_id"`
	SenderChat SenderChat `json:"sender_chat"`
	Chat       Chat       `json:"chat"`
	Date       int        `json:"date"`
	Audio      Audio      `json:"audio"`
}

type SendAudioArgs struct {
	ChatId   int
	FileName string
	File     []byte
	LogoFile []byte
	Caption  string

	Duration  string
	Title     string
	Performer string
}
