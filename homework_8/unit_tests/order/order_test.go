package main

import (
	"errors"
	"testing"
	"time"
)

func TestNewCustomer(t *testing.T) {
	t.Run("successfully make new customer", func(t *testing.T) {
		expectedEmail := "test@example.com"

		customer, err := newCustomer(expectedEmail)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		if customer.email != expectedEmail {
			t.Errorf("invalid email: got: %s, want: %s", customer.email, expectedEmail)
		}
	})

	t.Run("failed make new customer", func(t *testing.T) {
		testCases := map[string]struct {
			emailValue  string
			errorReason error
		}{
			"empty email": {
				emailValue:  "",
				errorReason: errEmailIsEmpty,
			},
			"email not valid": {
				emailValue:  "test",
				errorReason: errEmailIsNotValid,
			},
		}

		for tn, tc := range testCases {
			t.Run(tn, func(t *testing.T) {
				_, err := newCustomer(tc.emailValue)

				if !errors.Is(err, tc.errorReason) {
					t.Error(err)
				}
			})
		}
	})
}

func TestNewOrder(t *testing.T) {
	expectedEmail := "test@example.com"

	customer, _ := newCustomer(expectedEmail)

	order := newOrder(customer, Time{time.Now()})

	if !order.updatedAt.Equal(order.createdAt.Time) {
		t.Error("updatedAt must be equal to createdAt")
	}

	if !order.IsInitiated() {
		t.Errorf("order status must be %s", initiatedStatus)
	}

	if order.customer != customer {
		t.Errorf("invalid customer: got: %#v, want: %#v", order.customer, customer)
	}
}

func TestStatusChange(t *testing.T) {
	t.Run("successfully set processing status", func(t *testing.T) {
		expectedEmail := "test@example.com"

		customer, _ := newCustomer(expectedEmail)

		order := newOrder(customer, Time{time.Now()})

		err := order.SetProcessingStatus(Time{time.Now()})
		if err != nil {
			t.Errorf("cannot set processing status: %s", err)
		}

		if !order.IsProcessing() {
			t.Errorf("order status must be %q", processingStatus)
		}
	})
	t.Run("failed set processing status", func(t *testing.T) {
		expectedEmail := "test@example.com"

		customer, _ := newCustomer(expectedEmail)

		order := newOrder(customer, Time{time.Now()})

		order.status = failStatus

		err := order.SetProcessingStatus(Time{time.Now()})
		if err == nil {
			t.Errorf("cannot set processing status: %s", err)
		}
	})
	t.Run("successfully set fail status", func(t *testing.T) {
		expectedEmail := "test@example.com"

		customer, _ := newCustomer(expectedEmail)

		order := newOrder(customer, Time{time.Now()})

		order.SetProcessingStatus(Time{time.Now()})

		order.SetFailStatus(Time{time.Now()})

		if order.status != failStatus {
			t.Errorf("order status must be %q", failStatus)
		}
	})
	t.Run("failed set fail status", func(t *testing.T) {
		expectedEmail := "test@example.com"

		customer, _ := newCustomer(expectedEmail)

		order := newOrder(customer, Time{time.Now()})

		err := order.SetFailStatus(Time{time.Now()})
		if err == nil {
			t.Errorf("cannot set fail status: %s", err)
		}
	})
	t.Run("successfully set success status", func(t *testing.T) {
		expectedEmail := "test@example.com"

		customer, _ := newCustomer(expectedEmail)

		order := newOrder(customer, Time{time.Now()})

		order.SetProcessingStatus(Time{time.Now()})

		order.SetSuccessStatus(Time{time.Now()})

		if order.status != successStatus {
			t.Errorf("order status must be %q", successStatus)
		}
	})
	t.Run("failed set success status", func(t *testing.T) {
		expectedEmail := "test@example.com"

		customer, _ := newCustomer(expectedEmail)

		order := newOrder(customer, Time{time.Now()})

		err := order.SetSuccessStatus(Time{time.Now()})
		if err == nil {
			t.Errorf("cannot set success status: %s", err)
		}
	})
}
