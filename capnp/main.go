package main

import "zombiezen.com/go/capnproto2/rpc"

//import "zombiezen.com/go/capnproto2"
import "github.com/Noah-Huppert/rpc-compare/capnp/calc"
import "github.com/Noah-Huppert/rpc-compare/capnp/server"
import "context"
import "net"
import "log"
import "fmt"

func clientExec(ctx context.Context, conn net.Conn) error {
	rpcConn := rpc.NewConn(rpc.StreamTransport(conn))
	defer rpcConn.Close()

	calcClient := calc.Calculator{Client: rpcConn.Bootstrap(ctx)}

	log.Println("making Calculator.Add call")
	res, err := calcClient.Add(ctx, func(params calc.Calculator_add_Params) error {
		numbers, err := params.Numbers()
		if err != nil {
			return fmt.Errorf("error getting initial numbers array: %s", err.Error())
		}

		for i := 0; i < 5; i++ {
			log.Printf("setting value: %d\n", i+1)
			numbers.Set(i, int32(i+1))
		}

		return params.SetNumbers(numbers)
	}).Struct()

	if err != nil {
		log.Fatalf("error calling Calculator.Add: %s", err.Error())
	}

	log.Printf("the result of Calculator.Add was: %d", res)

	return nil
}

func main() {
	ctx := context.Background()
	serverConn, clientConn := net.Pipe()

	log.Println("starting Cap'N Proto RPC server")
	go server.Start(serverConn)

	log.Println("Connecting to server via client")
	err := clientExec(ctx, clientConn)
	if err != nil {
		log.Fatalf("error executing client logic: %s", err.Error())
	}
}
