package schema

import "encoding/json"

/* Login */

type FeishuLoginReq struct {
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

func (f FeishuLoginReq) Byte() []byte {
	b, _ := json.Marshal(f)
	return b
}

type FeishuLoginResp struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
}

/* FileUpload */
type FeishuFileUploadReq struct {
	FileType string `json:"file_type"`
	FileName string `json:"file_name"`
	Duration int    `json:"duration,omitempty"`
	File     []byte `json:"file"`
}

func (f FeishuFileUploadReq) Read(p []byte) (n int, err error) {
	b, _ := json.Marshal(f)
	copy(p, b)
	return len(p), nil
}

type FeishuFileUploadResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		FileKey string `json:"file_key"`
	} `json:"data"`
}

/* SendMessage */
type FeishuSendMessageReq struct {
	MsgType   string `json:"msg_type"`
	Content   string `json:"content"`
	ReceiveID string `json:"receive_id"`
}

func (f FeishuSendMessageReq) Byte() []byte {
	b, _ := json.Marshal(f)
	return b
}

type FeishuFileMessageContent struct {
	FileKey string `json:"file_key"`
}

func (f FeishuFileMessageContent) String() string {
	b, _ := json.Marshal(f)
	return string(b)
}
