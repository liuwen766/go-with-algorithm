package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

func main() {
	fmt.Println("rpc服务端启动......")
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("serve error:", err)
	}
}

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by 0")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}
