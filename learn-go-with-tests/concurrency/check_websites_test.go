package concurrency

import (
	"reflect"
	"testing"
)

func mockWebsitesChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsitesChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected %v, want %v", got, want)
	}
}
