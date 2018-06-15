.PHONY: gen

SRC=calc/calc.proto
OUT=./

# gen generates code from protobuf files
gen:
	protoc "${SRC}" --go_out="plugins=grpc:${OUT}"
