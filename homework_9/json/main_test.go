package main

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	t.Run("read exist file", func(t *testing.T) {
		_, err := readFile("./orders.json")
		if err != nil {
			t.Errorf("file mast be exist")
		}
	})
	t.Run("read not exist file", func(t *testing.T) {
		_, err := readFile("./test.json")
		if err == nil {
			t.Errorf("file must not exist")
		}
	})
}
