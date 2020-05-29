package test

import (
	"github.com/pangxieke/simple/controllers"
	"github.com/pangxieke/simple/routers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleGet(t *testing.T) {
	router := routers.New()

	reader := strings.NewReader(``)
	r, _ := http.NewRequest(http.MethodGet, "/hello", reader)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	resp := w.Result()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Response code is %v", resp.StatusCode)
	}

}

func TestHandleGet2(t *testing.T) {
	router := routers.New()
	router.AddPath("/hello", routers.Action((*controllers.HelloController).Hello))

	reader := strings.NewReader(``)
	r, _ := http.NewRequest(http.MethodGet, "/hello", reader)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Response code is %v", resp.StatusCode)
	}

}
