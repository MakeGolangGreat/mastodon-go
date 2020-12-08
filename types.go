package mastodon

const (
	public   Visibility = "public"
	unlisted Visibility = "unlisted"
	private  Visibility = "private"
	direct   Visibility = "direct"
)

type (
	// Status 嘟文类型
	// See https://docs.joinmastodon.org/entities/status/#example
	Status struct {
		ID        string  `json:"id"`
		CreatedAt string  `json:"created_at"`
		Content   string  `json:"content"`
		Account   Account `json:"account"`
	}
	// StatusParams 是发嘟时的参数
	StatusParams struct {
		Status      string `json:"status" validate:"required,max=500,min=1"`
		MediaIds    string `json:"media_ids"  validate:"required"`
		Poll        string `json:"poll"  validate:"required"`
		InReplyToID string `json:"in_reply_to_id,omitempty"`
		Visibility  string `json:"visibility,omitempty"`
		Sensitive   bool   `json:"sensitive"`
		SpoilerText string `json:"spoiler_text"`
	}
	// StatusRes 发嘟完的响应体
	StatusRes struct {
		ID string `json:"id"`
		Account   Account `json:"account"`
	}
	// Account 账号
	Account struct {
		UserName string `json:"username"`
		URL      string `json:"url"`
	}
	// Tag 标签
	Tag struct {
	}
	// Emoji 表情包
	Emoji struct {
	}
	// Media 媒体文件
	Media struct {
	}
	// Mention 提及
	Mention struct {
	}
	// Visibility 可见行
	Visibility string
	// HomeReq 首页列表请求体
	HomeReq struct {
		maxID   string
		sinceID string
		minID   string
		limit   string
		local   string
	}
	// HomeResp 首页列表的返回体
	HomeResp struct {
		account            Account
		application        interface{}
		bookmarked         bool
		card               interface{}
		Content            string
		createdAt          string
		emojis             []Emoji
		favourited         bool
		favouritesCount    int8
		ID                 string
		inReplyToAccountID string
		inReplyToID        string
		language           string
		mediaAttachments   []Media
		mentions           []Mention
		muted              bool
		pinned             bool
		poll               interface{}
		reblog             interface{}
		reblogged          bool
		reblogsCount       int8
		repliesCount       int8
		sensitive          bool
		spoilerText        string
		tags               []Tag
		uri                string
		url                string
		Visibility         Visibility
	}
)
