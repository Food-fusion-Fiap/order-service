package dtos

type ChangeOrderStatusDto struct {
	OrderId        string `json:"orderId"`
	ChangeToStatus string `json:"changeToStatus"`
}
