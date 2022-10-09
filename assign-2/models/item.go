package models

type Item struct {
	ID          uint   `gorm:"primaryKey;column:item_id" json:"lineItemId"`
	ItemCode    string `gorm:"column:item_code" json:"itemCode"`
	Description string `gorm:"column:description" json:"description"`
	Quantity    uint   `gorm:"column:quantity" json:"quantity"`
	OrderID     uint   `gorm:"column:order_id"`
}
