package models

type SocialType string

const (
	WhatsApp  SocialType = "whatsapp"
	Telegram  SocialType = "telegram"
	Instagram SocialType = "instagram"
	Facebook  SocialType = "facebook"
	Twitter   SocialType = "twitter"
	LinkedIn  SocialType = "linkedin"
)

type Social struct {
	FarmerID   int        `json:"farmer_id" db:"farmer_id"`
	Platfrom   SocialType `json:"platform" db:"platform"`
	AccountUrl string     `json:"account_url" db:"account_url"`
}
