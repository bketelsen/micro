package handler

import (
	"code.google.com/p/go.net/context"
	"code.google.com/p/goprotobuf/proto"

	math "github.com/bketelsen/micro/server/proto/math"
	log "github.com/golang/glog"
)

type Math struct{}

func (e *Math) Square(ctx context.Context, req *math.Request, rsp *math.Response) error {
	log.Info("Received Math.Square request")

	num := req.GetNum()
	resp := num * num

	rsp.Num = proto.Int(int(resp))

	return nil
}
