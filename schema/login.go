package schema

import (
	"encoding/json"

	"github.com/TCP404/eutil"
	"github.com/pkg/errors"
)

type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	AddrName string `json:"host"     validate:"required,oneof=agg online"`
	AddrHost string
}

func (r *LoginReq) String() string {
	b, _ := json.MarshalIndent(r, "", "  ")
	return string(eutil.B2S(b))
}

func LoginFixer(req *LoginReq) error {
	addrHost, err := fixAddr(req.AddrName)
	if err != nil {
		return errors.Wrap(err, "failed to parse host")
	}
	req.AddrHost = addrHost
	return nil
}

type LogoutReq struct {
	AddrName string `json:"host"`
}
