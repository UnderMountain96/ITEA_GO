package main

import (
	"errors"
	"fmt"
	"time"
)

type Status string

const (
	initiatedStatus  Status = "initiated"
	processingStatus        = "processing"
	successStatus           = "success"
	failStatus              = "fail"
)

var (
	errStatusIsNotProcessing = errors.New("must be \"processing\" status")
)

type Customer struct {
	email string
}

func newCustomer(email string) Customer {
	return Customer{
		email: email,
	}
}

type Time struct {
	time.Time
}

func (t Time) String() string {
	return t.Format("02.01.2006 15:04:05")
}

type Order struct {
	customer  Customer
	status    Status
	createdAt Time
	updatedAt Time
}

func newOrder(customer Customer, date Time) Order {
	return Order{
		customer:  customer,
		status:    initiatedStatus,
		createdAt: date,
		updatedAt: date,
	}
}

func (o *Order) IsProcessing() bool {
	return o.status == processingStatus
}

func (o *Order) statusChangeNotification() {
	fmt.Printf("Status changed to %q at %s\n", o.status, o.updatedAt)
}

func (o *Order) SetProcessingStatus(date Time) error {
	o.status = processingStatus
	o.updatedAt = date

	o.statusChangeNotification()

	return nil
}

func (o *Order) SetSuccessStatus(date Time) error {
	if !o.IsProcessing() {
		return fmt.Errorf(
			"SetSuccessStatus: cannot change status %q to %q, %w",
			o.status,
			successStatus,
			errStatusIsNotProcessing,
		)
	}

	o.status = successStatus
	o.updatedAt = date

	o.statusChangeNotification()

	return nil
}

func (o *Order) SetFailStatus(date Time) error {
	if !o.IsProcessing() {
		return fmt.Errorf(
			"SetFailStatus: cannot change status %q to %q, %w",
			o.status,
			failStatus,
			errStatusIsNotProcessing,
		)
	}

	o.status = failStatus
	o.updatedAt = date

	o.statusChangeNotification()

	return nil
}

func (o *Order) ShowInfoByOrder() {
	fmt.Println()
	fmt.Printf("customer email: %s\n", o.customer.email)
	fmt.Printf("order status: \t%s\n", o.status)
	fmt.Printf("created at: \t%s\n", o.createdAt)
	fmt.Printf("updated at: \t%s\n", o.updatedAt)
	fmt.Println()
}

func main() {
	timeNow := Time{time.Now()}
	orders := []Order{
		newOrder(newCustomer("Erik_Swift@example.com"), timeNow),
		newOrder(newCustomer("Olaf_Stout@example.com"), timeNow),
		newOrder(newCustomer("Baleog_Fierce@example.com"), timeNow),
	}

	for n, o := range orders {
		makeOrder(n, o)
	}
}

func makeOrder(n int, o Order) {
	defer o.ShowInfoByOrder()

	switch n {
	case 0:
		if err := o.SetSuccessStatus(Time{time.Now().Add(12 * time.Hour)}); err != nil {
			fmt.Printf("Failed change status: %s\n", err)
			return
		}
	case 1:
		if err := o.SetProcessingStatus(Time{time.Now().Add(6 * time.Hour)}); err != nil {
			fmt.Printf("Failed change status: %s\n", err)
			return
		}

		if err := o.SetFailStatus(Time{time.Now().Add(12 * time.Hour)}); err != nil {
			fmt.Printf("Failed change status: %s\n", err)
			return
		}

		if err := o.SetSuccessStatus(Time{time.Now().Add(12 * time.Hour)}); err != nil {
			fmt.Printf("Failed change status: %s\n", err)
			return
		}
	default:
		if err := o.SetProcessingStatus(Time{time.Now().Add(6 * time.Hour)}); err != nil {
			fmt.Printf("Failed change status: %s\n", err)
			return
		}

		if err := o.SetSuccessStatus(Time{time.Now().Add(12 * time.Hour)}); err != nil {
			fmt.Printf("Failed change status: %s\n", err)
			return
		}
	}
}
