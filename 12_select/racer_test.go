package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("check which server is fast", func(t *testing.T) {

		slowServer := makeDelayedServer(9 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		slowUrl := slowServer.URL //"http://www.facebook.com"
		fastUrl := fastServer.URL //"http://www.quii.co.uk"

		want := fastUrl

		got, _ := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

		slowServer.Close()
		fastServer.Close()
	})

	t.Run("return error if a server dosen't respond withing 10s", func(t *testing.T) {

		serverA := makeDelayedServer(20 * time.Millisecond)
		serverB := makeDelayedServer(21 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 20*time.Millisecond)

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
