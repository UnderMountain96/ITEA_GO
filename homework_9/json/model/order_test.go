package model

import (
	"testing"

	"github.com/google/uuid"
)

func TestOrderIsRefund(t *testing.T) {
	t.Run("transaction type refund", func(t *testing.T) {
		expectedId := uuid.MustParse("a628de4c-f967-4eb9-a2e6-1e8afa7020ff")
		order := Order{
			ID:     expectedId,
			Amount: 500.00,
			Transactions: []Transaction{
				{ID: uuid.MustParse("ca3796e1-716d-40b2-8a2f-4e09ca64edbd"), Type: AuthType},
				{ID: uuid.MustParse("6bc9a6bd-5eed-4f45-bbf9-e95c68e214bf"), Type: RefundType},
			},
		}

		if !order.IsRefund() {
			t.Errorf("order type must be %q", RefundType)
		}
	})
	t.Run("transaction type not refund", func(t *testing.T) {
		expectedId := uuid.MustParse("a628de4c-f967-4eb9-a2e6-1e8afa7020ff")
		order := Order{
			ID:     expectedId,
			Amount: 500.00,
			Transactions: []Transaction{
				{ID: uuid.MustParse("ca3796e1-716d-40b2-8a2f-4e09ca64edbd"), Type: AuthType},
				{ID: uuid.MustParse("6bc9a6bd-5eed-4f45-bbf9-e95c68e214bf"), Type: SettleType},
			},
		}

		if order.IsRefund() {
			t.Errorf("order type not must be %q", RefundType)
		}
	})
	t.Run("orders refund", func(t *testing.T) {
		orders := Orders{Orders: []Order{
			{
				ID:     uuid.MustParse("a628de4c-f967-4eb9-a2e6-1e8afa7020ff"),
				Amount: 500.00,
				Transactions: []Transaction{
					{ID: uuid.MustParse("ca3796e1-716d-40b2-8a2f-4e09ca64edbd"), Type: AuthType},
					{ID: uuid.MustParse("6bc9a6bd-5eed-4f45-bbf9-e95c68e214bf"), Type: RefundType},
				},
			},
			{
				ID:     uuid.MustParse("dcf61566-66e9-4a8b-a26d-f5549cd4e00d"),
				Amount: 500.00,
				Transactions: []Transaction{
					{ID: uuid.MustParse("69f343ca-a8d7-4f9b-bb7d-e8106e49d002"), Type: AuthType},
					{ID: uuid.MustParse("a25637ec-ebed-4dc5-881b-99c3c4ce4b4e"), Type: RefundType},
				},
			},
		}}

		expectedRefundOrders := 2
		refundOrders := len(orders.GetRefundOrders())

		if refundOrders != expectedRefundOrders {
			t.Errorf("invalid refund orders: got: %d, want: %d", refundOrders, refundOrders)
		}
	})
	t.Run("transaction type settle", func(t *testing.T) {
		expectedId := uuid.MustParse("a628de4c-f967-4eb9-a2e6-1e8afa7020ff")
		expectedType := SettleType
		transaction := Transaction{ID: expectedId, Type: expectedType}

		if transaction.isRefund() {
			t.Errorf("transaction type must be %q", expectedType)
		}
	})
	t.Run("transaction type refund", func(t *testing.T) {
		expectedId := uuid.MustParse("a628de4c-f967-4eb9-a2e6-1e8afa7020ff")
		expectedType := RefundType
		transaction := Transaction{ID: expectedId, Type: expectedType}

		if !transaction.isRefund() {
			t.Errorf("transaction type must be %q", expectedType)
		}
	})
}

func TestTransactionIsRefund(t *testing.T) {
	t.Run("transaction type auth", func(t *testing.T) {
		expectedId := uuid.MustParse("a628de4c-f967-4eb9-a2e6-1e8afa7020ff")
		expectedType := AuthType
		transaction := Transaction{ID: expectedId, Type: expectedType}

		if transaction.isRefund() {
			t.Errorf("transaction type must be %q", expectedType)
		}
	})
	t.Run("transaction type settle", func(t *testing.T) {
		expectedId := uuid.MustParse("a628de4c-f967-4eb9-a2e6-1e8afa7020ff")
		expectedType := SettleType
		transaction := Transaction{ID: expectedId, Type: expectedType}

		if transaction.isRefund() {
			t.Errorf("transaction type must be %q", expectedType)
		}
	})
	t.Run("transaction type refund", func(t *testing.T) {
		expectedId := uuid.MustParse("a628de4c-f967-4eb9-a2e6-1e8afa7020ff")
		expectedType := RefundType
		transaction := Transaction{ID: expectedId, Type: expectedType}

		if !transaction.isRefund() {
			t.Errorf("transaction type must be %q", expectedType)
		}
	})
}
