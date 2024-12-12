package router

import (
	"esdumpweb/initial"
	"log/slog"
	"os"
	"testing"
)

var (
	lm = initial.WireLoggerManager()
)

func testSetup() error {
	for range 10 {
		slog.Info("test", "test", "test")
	}
	return nil
}

func Test_getAccessToken(t *testing.T) {
	ret, err := getAccessToken()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ret.TenantAccessToken)
}

func Test_uploadFile(t *testing.T) {
	ret, err := getAccessToken()
	if err != nil {
		t.Fatal(err)
	}
	token := ret.TenantAccessToken

	upRet, err := uploadFile(token)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(upRet.Data.FileKey)
}

func Test_sendMessage(t *testing.T) {
	ret, err := getAccessToken()
	if err != nil {
		t.Fatal(err)
	}
	token := ret.TenantAccessToken
	fileKey := "file_v3_00hg_c5a49f4c-769c-4c97-b541-2656d26aad7g"
	sendMessage(token, fileKey)
}

func TestMain(m *testing.M) {
	if err := testSetup(); err != nil {
		os.Exit(2)
	}
	code := m.Run()
	lm.Close()
	os.Exit(code)
}
