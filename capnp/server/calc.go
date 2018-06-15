package server

import "zombiezen.com/go/capnproto2/rpc"
import "github.com/Noah-Huppert/rpc-compare/capnp/calc"
import "fmt"
import "net"

// CalcServer implements the Calculator service
type CalcServer struct{}

// Add implements Calculator.Add
func (c CalcServer) Add(call calc.Calculator_add) error {
	numbers, err := call.Params.Numbers()
	if err != nil {
		return fmt.Errorf("error reading data: %s", err.Error())
	}

	var accum int32 = 0
	for i := 0; i < numbers.Len(); i++ {
		accum += numbers.At(i)
	}

	call.Results.SetResult(accum)

	return nil
}

func Start(conn net.Conn) error {
	calcServer := CalcServer{}
	server := calc.Calculator_ServerToClient(calcServer)

	rpcConn := rpc.NewConn(rpc.StreamTransport(conn), rpc.MainInterface(server.Client))

	return rpcConn.Wait()
}
