package constant

import (
	"time"
)

var (
	VERSION         string
	CST             = time.FixedZone("CST", 8*60*60)
	CtxKeyReqParams = "request-params"
	MaxSize         = 10000
)

const (
	// AES加密密钥
	AES_KEY = "esdumpweb1997"
)

const (
	HostAgg    = "es-cn-zz11sf3n8002ft2vw.elasticsearch.aliyuncs.com:9200"
	HostOnline = "es-cn-n6w24o5er00avk3v6.elasticsearch.aliyuncs.com:9200"

	IndexClueOnline      = "clue_online_alias"
	IndexClueAgg         = "clue_agg_alias"
	IndexClueOverseasAgg = "clue_overseas_agg_alias"
	IndexCrowdOnline     = "crowd_online_alias"
	IndexCrowdAgg        = "crowd_agg_alias"
	IndexToolsOnline     = "tools_online_alias"
	IndexToolsAgg        = "tools_agg_alias"

	FieldInsertTime = "insert_time"
	FieldCreateTime = "create_time"
	FieldEventTime  = "event_time"
)

var Indices = []string{
	IndexClueOnline,
	IndexClueAgg,
	IndexClueOverseasAgg,
	IndexCrowdOnline,
	IndexCrowdAgg,
	IndexToolsOnline,
	IndexToolsAgg,
}

var TimeFields = []string{
	FieldInsertTime,
	FieldCreateTime,
	FieldEventTime,
}

const (
	FeishuFileUploadURL  = "https://open.feishu.cn/open-apis/im/v1/files"
	FeishuSendMessageURL = "https://open.feishu.cn/open-apis/im/v1/messages"
	FeishuLoginURL       = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
	FeishuAppID          = "cli_a5f513c4cbb81013"
	FeishuAppSecret      = "WGyqelxWp9MlfBPdGhS8UeOk3d5Lqzo2"
	FeishuChatID         = "oc_f1ca21e808792b06815259445943e5bd"
	FeishuTestChatID     = "oc_71abd3df82160ec57ed7801fd54f0e31"
	FeishuBoundary       = "---7MA4YWxkTrZu0gW"
)
