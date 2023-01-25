package hertzstudy

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
)

func main() {
	c, err := client.NewClient()
	if err != nil {
		return
	}
	//send ttp get request
	status, body, err := c.Get(context.Background(), nil, "http:///www.example.com")
	if err != nil {
		return
	}
	fmt.Printf("status = %v body = %v\n", status, string(body))
	//send http post request
	var postArgs protocol.Args
	//set post args
	postArgs.Set("arg", "rainyday")
	status, body, err = c.Get(context.Background(), nil, "http:///www.example.com")
	if err != nil {
		return
	}
	fmt.Printf("status = %v body = %v\n", status, string(body))
}
