package processing

import (
	"bytes"
	"log"
	"os"
	"testing"
)

var (
	TEST_RESOURCES_PATH = "testdata"
)

func readFile(relativePath string) []byte {
	path := TEST_RESOURCES_PATH + relativePath
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error during reading a file from path '%s'", path)
		panic(err)
	}
	return data
}

func TestGatherHRefs(t *testing.T) {
	// make multiple html files to test with different number of hrefs and create a loop of tests for that
	t.Run("Analyze html file and return 6 hrefs", func(t *testing.T) {
		path := "/Better_Motherfucking_Website.html"
		data := readFile(path)
		reader := bytes.NewReader(data)
		hrefs, _ := GatherHRefs(reader)
		if len(hrefs) != 6 {
			t.Errorf("From '%s' returned '%d' number of refs, wanted '%d'", path, len(hrefs), 6)
		}
	})
}
