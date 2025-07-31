package main

import (
	"fmt"
	"os"
	"testing"
)

var (
	TEST_RESOURCES_PATH = "testdata"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestMain(t *testing.T) {
	t.Run("main test case", func(t *testing.T) {
		data, err := os.ReadFile(TEST_RESOURCES_PATH + "/Better_Motherfucking_Website.html")
		check(err)

		fmt.Println(string(data))
	})
}
