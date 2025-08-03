package main

import (
	"fmt"
	"os"
	"testing"
)

var (
	TEST_RESOURCES_PATH = "testdata"
)

func TestMain(t *testing.T) {
	t.Run("main test case", func(t *testing.T) {
		data, _ := os.ReadFile(TEST_RESOURCES_PATH + "/Better_Motherfucking_Website.html")

		fmt.Println(data)
	})
}
