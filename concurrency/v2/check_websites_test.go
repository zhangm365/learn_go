package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {

	return url != "waat://furhurterwe.geds"

}

func TestCheckWebsite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	actualResults := CheckWebsites(mockWebsiteChecker, websites)

	want := len(websites)
	got := len(actualResults)

	if got != want {
		t.Fatalf("want %v, got %v", want, got)
	}

	expectedResult := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	if !reflect.DeepEqual(expectedResult, actualResults) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
