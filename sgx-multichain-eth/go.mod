module ethereumLink

go 1.15

require github.com/ethereum/go-ethereum v0.0.0

require (
	google.golang.org/grpc v1.40.0
	kunmel.com/sgx-multichain/multichain_proto v0.0.0
)

replace github.com/ethereum/go-ethereum v0.0.0 => ./../../github.com/ethereum/go-ethereum

replace kunmel.com/sgx-multichain/multichain_proto => ../multichain_proto

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.3

replace google.golang.org/grpc => google.golang.org/grpc v1.23.0
