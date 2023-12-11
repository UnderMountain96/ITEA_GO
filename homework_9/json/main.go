package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/UnderMountain96/ITEA_GO/json/model"
)

func main() {
	byteValue, err := readFile("./orders.json")
	if err != nil {
		panic(err)
	}

	var orders model.Orders

	if err := json.Unmarshal(byteValue, &orders.Orders); err != nil {
		panic(err)
	}

	for _, id := range orders.GetRefundOrders() {
		fmt.Printf("Order %q is %s.\n", id, model.RefundType)
	}
}

func readFile(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	return io.ReadAll(jsonFile)
}
