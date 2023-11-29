package main

import (
	"testing"
	"time"
)

func TestNewCustomer(t *testing.T) {
	expectedEmail := "test@example.com"

	customer, err := newCustomer(expectedEmail)
	if err != nil {
		t.Errorf("cannot create new customer: %s", err)
	}

	if customer.email != expectedEmail {
		t.Errorf("invalid email: got: %s, want: %s", customer.email, expectedEmail)
	}

	testCases := map[string]struct {
		emailValue  string
		errorReason string
	}{
		"empty email": {
			emailValue:  "",
			errorReason: "cannot set empty email",
		},
		"email not valid": {
			emailValue:  "test",
			errorReason: "email is not valid",
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			_, err := newCustomer(tc.emailValue)
			if err == nil {
				t.Error(tc.errorReason)
			}
		})
	}
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

func TestIsProcessing(t *testing.T) {
	expectedEmail := "test@example.com"

	customer, _ := newCustomer(expectedEmail)

	order := newOrder(customer, Time{time.Now()})

	if order.IsProcessing() {
		t.Errorf("order srarus must be %q", initiatedStatus)
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
			t.Errorf("order srarus must be %q", processingStatus)
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
			t.Errorf("order srarus must be %q", failStatus)
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
			t.Errorf("order srarus must be %q", successStatus)
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
