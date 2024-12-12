package schema

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"time"

	"esdumpweb/constant"
	"esdumpweb/initial"

	"github.com/TCP404/esdumpcore/core"
	"github.com/TCP404/esdumpcore/outputer"
	"github.com/TCP404/eutil"
	"github.com/pkg/errors"
)

type (
	E = core.Hit
	L = core.Hit
)

type DumpReq struct {
	AddrName      string    `json:"host"         validate:"required,oneof=agg online"`
	Index         string    `json:"index"        validate:"required"`
	TimeField     string    `json:"timeField"    validate:"required,oneof=insert_time event_time"`
	StartTime     time.Time `json:"startTime"    validate:"required,gtfield=EndTime"`
	EndTime       time.Time `json:"endTime"      validate:"required,ltfield=StartTime,ltfield=180d"`
	Product       string    `json:"product"      validate:"required,gt=0"`
	Condition     string    `json:"condition"`
	SaveLocation  string
	AddrHost      string
	OutputHandler outputer.Outputer[E]
	BodyBool      *core.ESBodyBool
}

func (r *DumpReq) String() string {
	// b, _ := json.Marshal(r)
	// return string(eutil.B2S(b))
	return fmt.Sprintf("%+v", *r)
}

func fixAddr(hostName string) (string, error) {
	switch hostName {
	case "agg":
		return constant.HostAgg, nil
	case "online":
		return constant.HostOnline, nil
	default:
		return "", errors.New("invalid host")
	}
}

func fixIndex(index string) error {
	if !eutil.In(index, constant.Indices...) {
		return errors.New("invalid index")
	}
	return nil
}

func fixCondition(cond string) (*core.ESBodyBool, error) {
	if len(cond) == 0 {
		return new(core.ESBodyBool), nil
	}
	var condition core.ESBodyBool
	err := json.Unmarshal([]byte(cond), &condition)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse condition")
	}
	return &condition, nil
}

func fixOutputHandler(location string) (outputer.Outputer[E], error) {
	if len(location) == 0 {
		return nil, errors.New("save location is required")
	}
	switch filepath.Ext(location) {
	case ".csv":
		return outputer.NewCSV[E](location), nil
	default:
		return nil, nil
	}
}

func DumpFixer(req *DumpReq) error {
	host, err := fixAddr(req.AddrName)
	if err != nil {
		return errors.Wrap(err, "failed to parse addr")
	}
	req.AddrHost = "http://" + host

	if err := fixIndex(req.Index); err != nil {
		return errors.Wrap(err, "failed to parse index")
	}

	dir, filename := initial.WireConfig().SaveDir, fmt.Sprintf("%s_%s_%s_%s_%s.csv", req.Index, req.Product, req.TimeField, req.StartTime.Format("20060102150405"), req.EndTime.Format("20060102150405"))
	req.SaveLocation = filepath.Join(dir, filename)

	outputerHandler, err := fixOutputHandler(req.SaveLocation)
	if err != nil {
		return errors.Wrap(err, "failed to parse save location")
	}
	req.OutputHandler = outputerHandler

	condition, err := fixCondition(req.Condition)
	if err != nil {
		return errors.Wrap(err, "failed to parse condition")
	}
	condition.Filter = append(condition.Filter, core.M{"term": core.M{"product.keyword": req.Product}})
	req.BodyBool = condition
	return nil
}
