package model

import (
	"github.com/google/uuid"
)

type TransactionType string

const (
	AuthType   TransactionType = "auth"
	SettleType TransactionType = "settle"
	RefundType TransactionType = "refund"
)

type Orders struct {
	Orders []Order
}

func (o *Orders) GetRefundOrders() []uuid.UUID {
	refundOrders := []uuid.UUID{}
	for _, o := range o.Orders {
		if o.IsRefund() {
			refundOrders = append(refundOrders, o.ID)
		}
	}
	return refundOrders
}

type Order struct {
	ID           uuid.UUID     `json:"id"`
	Amount       float64       `json:"amount"`
	Transactions []Transaction `json:"transactions"`
}

func (o *Order) IsRefund() bool {
	for _, t := range o.Transactions {
		if t.isRefund() {
			return true
		}
	}

	return false
}

type Transaction struct {
	ID   uuid.UUID       `json:"id"`
	Type TransactionType `json:"type"`
}

func (t *Transaction) isRefund() bool {
	return t.Type == RefundType
}
