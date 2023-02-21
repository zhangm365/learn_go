package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	// slowURL := "https://www.google.com"
	// fastURL := "https://www.baidu.com"
	t.Run("Compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {

		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		// 函数结束时调用
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		// fmt.Println(slowURL, fastURL)
		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got '%s', want '%s'.", got, want)
		}

	})

	t.Run("return an error if a server doesn't respond within 10s.", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Second)

		// 函数结束时调用
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}

	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
