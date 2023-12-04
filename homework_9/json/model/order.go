package model

import (
	"fmt"

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

func (o *Orders) ShowRefundOrders() {
	for _, o := range o.Orders {
		if id := o.IsRefund(); id != uuid.Nil {
			fmt.Printf("Order %q is %s.\n", id, RefundType)
		}
	}
}

type Order struct {
	ID           uuid.UUID     `json:"id"`
	Amount       float64       `json:"amount"`
	Transactions []Transaction `json:"transactions"`
}

func (o *Order) IsRefund() uuid.UUID {
	for _, t := range o.Transactions {
		if t.isRefund() {
			return o.ID
		}
	}

	return uuid.Nil
}

type Transaction struct {
	ID   uuid.UUID       `json:"id"`
	Type TransactionType `json:"type"`
}

func (t *Transaction) isRefund() bool {
	return t.Type == RefundType
}
