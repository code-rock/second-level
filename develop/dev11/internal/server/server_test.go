package server

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestIndexFunc(t *testing.T) {
	rw := httptest.NewRecorder()

	form := url.Values{}
	form.Add("date", "09-17-9")
	form.Add("event", "Митиниг в 9:00")
	req := httptest.NewRequest(http.MethodPost, "/create_event", bytes.NewBufferString(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Вызов handler IndexFunc
	// IndexFunc(rw, req)

	if rw.Result().StatusCode != http.StatusOK {
		t.Errorf("Request code %v not equal code 200", rw.Code)
	}
	if rw.Body.String() != "Иван Петров" {
		t.Errorf(`Request body "%v" not equal body "Иван Петров"`, rw.Body.String())
	}
}
