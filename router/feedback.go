package router

import (
	"bytes"
	"esdumpweb/constant"
	"esdumpweb/initial"
	"esdumpweb/schema"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/TCP404/eutil/fetch"
	"github.com/gin-gonic/gin"
)

func feedbackHandle(c *gin.Context) {
	ret, err := getAccessToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	token := ret.TenantAccessToken
	uploadRet, err := uploadFile(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	fileKey := uploadRet.Data.FileKey
	err = sendMessage(token, fileKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func getAccessToken() (schema.FeishuLoginResp, error) {
	var (
		ret schema.FeishuLoginResp
		err error
	)
	fetch.Fetch(
		http.MethodPost,
		constant.FeishuLoginURL,
		fetch.WithByteBody(schema.FeishuLoginReq{
			AppID:     constant.FeishuAppID,
			AppSecret: constant.FeishuAppSecret,
		}.Byte()),
	).
		ThenWriteTo(&ret).
		Catch(func(resp *http.Response, e error) {
			slog.Error("get access token failed", "err", e)
			err = e
		}).
		Await()
	return ret, err
}

func uploadFile(token string) (*schema.FeishuFileUploadResp, error) {
	var (
		lm       = initial.WireLoggerManager()
		filebase = filepath.Base(lm.GetLogFilePath())
		filename = fmt.Sprintf("%v_%v.json", filebase[:len(filebase)-len(filepath.Ext(filebase))], lm.GetLogTime().Format("20060102150405"))
		body     = &bytes.Buffer{}
		writer   = multipart.NewWriter(body)
	)
	writer.SetBoundary(constant.FeishuBoundary)
	writer.WriteField("file_type", "stream")
	writer.WriteField("file_name", filename)

	part, err := writer.CreateFormField("file")
	if err != nil {
		return nil, err
	}

	f, err := os.Open(lm.GetLogFilePath())
	if err != nil {
		return nil, err
	}
	defer f.Close()
	_, err = io.Copy(part, f)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	var ret schema.FeishuFileUploadResp
	fetch.
		Fetch(
			http.MethodPost,
			constant.FeishuFileUploadURL,
			fetch.WithHeader("Content-Type", writer.FormDataContentType()),
			fetch.WithHeader("Authorization", "Bearer "+token),
			fetch.WithReaderBody(body),
		).
		ThenWriteTo(&ret).
		Then(func(resp *http.Response) {
			slog.Info("upload feedback success")
		}).
		Catch(func(resp *http.Response, e error) {
			err = e
			b, _ := io.ReadAll(resp.Body)
			slog.Error("upload feedback failed", "err", e, "body", string(b))
		}).
		Await()

	return &ret, err
}

func sendMessage(token, fileKey string) error {
	var err error
	fetch.Fetch(
		http.MethodPost,
		constant.FeishuSendMessageURL,
		fetch.WithParam("receive_id_type", "chat_id"),
		fetch.WithHeader("Authorization", "Bearer "+token),
		fetch.WithHeader("Content-Type", "application/json; charset=utf-8"),
		fetch.WithByteBody(schema.FeishuSendMessageReq{
			MsgType: "file",
			// ReceiveID: constant.FeishuTestChatID,
			ReceiveID: constant.FeishuChatID,
			Content:   string(schema.FeishuFileMessageContent{FileKey: fileKey}.String()),
		}.Byte()),
	).
		Then(func(resp *http.Response) {
			slog.Info("send message success")
		}).
		Catch(func(resp *http.Response, e error) {
			err = e
			b, _ := io.ReadAll(resp.Body)
			slog.Error("send message failed", "err", e, "body", string(b))
		}).
		Await()
	return err
}
