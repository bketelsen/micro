package main

import (
	"fmt"

	"code.google.com/p/goprotobuf/proto"
	"github.com/asim/go-micro/client"
	math "github.com/bketelsen/micro/server/proto/math"
)

func main() {
	// Create new request to service go.micro.service.go-template, method Example.Call
	req := client.NewRequest("sample.math", "Math.Square", &math.Request{
		Num: proto.Int(4),
	})

	// Set arbitrary headers
	req.Headers().Set("X-User-Id", "john")
	req.Headers().Set("X-From-Id", "script")

	rsp := &math.Response{}

	// Call service
	if err := client.Call(req, rsp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.GetNum())

	// now call it using JSON RPC
	var n int32
	n = 4

	// note using standard types here, but the input protocol
	// requires a pointer to an int instead of an int
	// so we dance a little first (above)
	req1 := client.NewJsonRequest("sample.math", "Math.Square", &math.Request{
		Num: &n,
	})
	rsp = &math.Response{}
	// Call service
	if err := client.Call(req1, rsp); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.GetNum())

}
