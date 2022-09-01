package models

type Charge struct {
	ID           uint   `json:"id" gorm:"primary_key;auto_increment"`
	UserID       string `json:"user_id"`
	Amount       int64  `json:"amount"`
	ReceiptEmail string `json:"receipt_email"`
	CardToken    string `json:"card_token"`
	Name         string `json:"name"`
}

func (c *Charge) TableName() string {
	return "charges"
}
