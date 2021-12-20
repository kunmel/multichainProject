module sgx-multichain/gin_demo

go 1.15

require (
	github.com/gin-gonic/gin v1.7.7
	google.golang.org/grpc v1.43.0
	kunmel.com/sgx-multichain/multichain_proto v0.0.0
)

replace kunmel.com/sgx-multichain/multichain_proto => ../multichain_proto
replace google.golang.org/grpc => google.golang.org/grpc v1.23.0