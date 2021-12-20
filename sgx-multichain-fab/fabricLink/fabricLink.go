package fabricLink

import (
	"log"

	"fabricLink/cli"
)

const (
	org1CfgPath = "/home/guotiezheng/go/src/sgx-multichain/sgx-multichain-fab/config/org1sdk-config.yaml"
	org2CfgPath = "/home/guotiezheng/go/src/sgx-multichain/sgx-multichain-fab/config/org2sdk-config.yaml"
)

var (
	peer0Org1 = "peer0.org1.example.com"
	peer0Org2 = "peer0.org2.example.com"
)

func FabricLink(ccName string, ccArgs []string, fcnName string) (int32, string) {
	org1Client := cli.New(org1CfgPath, "Org1", "Admin", "User1")
	org2Client := cli.New(org2CfgPath, "Org2", "Admin", "User1")

	defer org1Client.Close()
	defer org2Client.Close()

	// Install, instantiate, invoke, query
	return DealTx(org1Client, org2Client, ccName, ccArgs, fcnName)
}


func DealTx(cli1, cli2 *cli.Client, ccName string, ccArgs []string, fcnName string) (int32, string) {
	log.Println("=================== DealTx begin ===================")
	defer log.Println("=================== DealTx end ===================")
	txID, err := cli1.InvokeCC([]string{peer0Org1, peer0Org2}, ccName, ccArgs, fcnName)
	if err != nil {
		log.Panicf("Invoke chaincode error: %v", err)
		return 404, ""
	}
	log.Println("Invoke chaincode success")
	_, err = cli1.InvokeCC([]string{peer0Org1, peer0Org2}, ccName, ccArgs, "readLedger")
	if err != nil {
		log.Panicf("Invoke chaincode error: %v", err)
		return 404, ""
	}
	log.Println("Invoke chaincode success")
	//if err := cli1.InvokeCC("peer0.org1.example.com", "a"); err != nil {
	//	log.Panicf("Query chaincode error: %v", err)
	//	return 404, ""
	//}
	//log.Println("Query chaincode success on peer0.org1")
	return 200, string(txID)
}

