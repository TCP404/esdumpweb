package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"esdumpweb/initial"
	"esdumpweb/router"
	"esdumpweb/test/testkit"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	app *gin.Engine
	lm  = initial.WireLoggerManager()
)

func setup() error {
	initial.WireConfigManager()
	app = router.App()
	return nil
}

func TestLoginRouter(t *testing.T) {
	tests := []*testkit.PostReq{
		testkit.NewPostReq("/login",
			map[string]any{
				"username": "admin",
				"password": "admin",
				"host":     "agg",
			},
		),
	}
	w := httptest.NewRecorder()
	for _, tc := range tests {
		req, err := tc.Req()
		if err != nil {
			t.Errorf("generate request failed: %v", err)
		}

		app.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		t.Log(w.Body.String())
	}
}

func TestProcessRouter(t *testing.T) {
	req, err := testkit.NewGetReq("/process", nil).Req()
	if err != nil {
		t.Errorf("generate request failed: %v", err)
	}
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	t.Log(w.Body.String())
}

func TestDumpRouter(t *testing.T) {
	req, err := testkit.NewPostReq("/dump", map[string]any{
		"host":      "online",
		"index":     "clue_online_alias",
		"timeField": "insert_time",
		"startTime": "2024-11-01T00:00:00+08:00",
		"endTime":   "2024-11-06T01:00:00+08:00",
		"product":   "小红书",
	}).Req()

	if err != nil {
		t.Errorf("generate request failed: %v", err)
	}
	w := httptest.NewRecorder()
	app.ServeHTTP(w, req)
	t.Logf("response body: %v", w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		os.Exit(2)
	}
	code := m.Run()
	lm.Close()
	os.Exit(code)
}
