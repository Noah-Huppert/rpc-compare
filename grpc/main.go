package main

import "log"
import "context"
import "github.com/Noah-Huppert/rpc-compare/grpc/server"
import "github.com/Noah-Huppert/rpc-compare/grpc/calc"
import "google.golang.org/grpc"

func main() {
	grpcAddr := "localhost:9000"
	ctx := context.Background()

	// Start server
	log.Println("starting GRPC server")

	err := server.Start(grpcAddr)
	if err != nil {
		log.Fatalf("error starting GRPC server: %s", err.Error())
	}

	// Connect via client
	log.Println("connecting via GRPC client")

	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("error connecting to GRPC server: %s", err.Error())
	}

	client := calc.NewCalculatorClient(conn)

	// Make RPC call
	numbers := &calc.Numbers{
		Numbers: []int32{1, 2, 3, 4, 5},
	}

	res, err := client.Add(ctx, numbers)
	if err != nil {
		log.Fatalf("error calling Calc.Add method: %s", err.Error())
	}

	log.Printf("Calc.Add result was: %d", res.Result)
}
