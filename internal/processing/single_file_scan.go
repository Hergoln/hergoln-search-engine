package processing

import (
	"bytes"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RunSingleFileScan(path *string) {
	fmt.Println(*path)
	data, err := os.ReadFile(*path)
	check(err)

	reader := bytes.NewReader(data)
	hrefs, _ := GatherHRefs(reader)

	fmt.Println(hrefs)
}
