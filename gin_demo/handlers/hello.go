package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"io"
	pb "kunmel.com/sgx-multichain/multichain_proto"
	"log"
	"net/http"
	"strconv"
)


func PutOnChain(c *gin.Context) {
	fabIP := ":2222"
	fabPem := "/home/guotiezheng/go/src/sgx-multichain/key/test.pem"
	ethIP := ":3333"
	ethPem := "/home/guotiezheng/go/src/sgx-multichain/key/test.pem"
	serverName := "www.kunmellh.com"
	var client pb.MultichainServiceClient
	var response int32
	var blockID string
	var txID string
	id := c.Query("id")
	price := c.Query("price")
	label := c.Query("label")
	chainID := c.Query("chainID")
	priceInt,_  := strconv.Atoi(price)
	if  chainID == "hyperledger fabric" || priceInt < 10000 {
		client = setupConnect(fabIP, fabPem, serverName)
		response, blockID, txID = sendFab(client, "mycc", []string{id, price, label}, "writeLedger")
	} else if chainID == "ethereum" || priceInt >=10000 {
		client = setupConnect(ethIP, ethPem, serverName)
		response, blockID, txID= sendEth(client, "UTC--2021-09-15T05-28-00.066630039Z--6a2e96b16eec57ba9f21b1b5c5bdc476a2d3676f", "22b8", "123456",  "0x120615426fbf165d6673befb179f7faaa9f08c6a", "addProject", id, price, label)
	}
	c.JSON(http.StatusOK, gin.H{
		"status" : decodeRsponse(response),
		"txID" : txID,
		"blockID" :  blockID,
	})
}


func setupConnect(ip string, pemPath string, serverName string)  pb.MultichainServiceClient {
	creds, err := credentials.NewClientTLSFromFile(pemPath, serverName)
	if err != nil {
		grpclog.Fatalf("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(ip, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln("peer grpc client cannot dial ", ip)
	}
	fmt.Println("gRPC connection created")
	client := pb.NewMultichainServiceClient(conn)
	return client
}

func sendFab(client pb.MultichainServiceClient, ccName string, ccArg []string, fcnName string) (int32, string, string) {
	serverStream, err := client.NewFabTx(context.Background(), &pb.NewFabRequest{
		CCName: ccName,
		CCArg: ccArg,
		FcnName: fcnName,
	})
	if err != nil {
		log.Fatalln(err)
	}
	var result *pb.CommitResponse
	var statuscode int32
	var blockID string
	var txID string
	for {
		result, err = serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		statuscode = result.StatusCode
		blockID = result.BlockID
		txID = result.TxID
	}
	return statuscode, blockID, txID
}

func sendEth(client pb.MultichainServiceClient, keyPath string, chainID string, pass string, contractAddress string, fcnName string, id string, price string, label string) (int32,string,string) {
	serverStream, err := client.NewEthTx(context.Background(), &pb.NewEthRequest{
		KeystorePath: keyPath,
		ChainID: chainID,
		Pass: pass,
		ContractAddress: contractAddress,
		FcnName: fcnName,
		Id: id,
		Price: price,
		Label: label,
	})
	if err != nil {
		log.Fatalln(err)
	}
	var response *pb.CommitResponse
	var statuscode int32
	var blockID string
	var txID string
	for {
		response, err = serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		statuscode = response.StatusCode
		blockID = response.BlockID
		txID = response.TxID
	}
	return statuscode, blockID, txID
}

func decodeRsponse(resp int32) string{
	if resp == 200 {
		return "success "
	} else if resp == 404 {
		return "error"
	}
	return "error"
}
