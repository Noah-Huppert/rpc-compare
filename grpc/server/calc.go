package server

import "github.com/Noah-Huppert/rpc-compare/grpc/calc"
import "context"
import "net"
import "google.golang.org/grpc"
import "fmt"

// CalcServer implements the Calc service
type CalcServer struct{}

// Add implements Calc.Add
func (c *CalcServer) Add(ctx context.Context, numbers *calc.Numbers) (
	*calc.Result, error) {

	var accum int32 = 0

	for _, n := range numbers.Numbers {
		accum += n
	}

	result := &calc.Result{
		Result: accum,
	}

	return result, nil
}

func Start(grpcAddr string) error {
	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to start RCP server: %s", err.Error())
	}

	calcServer := &CalcServer{}

	grpcServer := grpc.NewServer()
	calc.RegisterCalculatorServer(grpcServer, calcServer)

	go func() {
		grpcServer.Serve(lis)
	}()
	return nil
}
