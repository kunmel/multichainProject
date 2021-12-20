package ethereumLink

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/contract/putonchain"
	"math/big"
	"os"
)

func EthereumLink(keystorePath string, chainID string, pass string, contractAddr string, fcnName string, id string, price string, label string) (int32, string) {
	// 连接rpc
	client,err := ethclient.Dial("http://127.0.0.1:8545")

	// 定义要操作合约的账户地址
	//addr := "0x6a2e96b16eec57ba9f21b1b5c5bdc476a2d3676f"
	// 将字符串地址转为common.Address
	//common_addr := common.HexToAddress(addr)
	//user_addr := common.HexToAddress("0x6a2e96b16eec57ba9f21b1b5c5bdc476a2d3676f")
	// 此处用的是智能合约的地址
	//contract_addr := common.HexToAddress("0x120615426fbf165d6673befb179f7faaa9f08c6a")
	contract_addr := common.HexToAddress(contractAddr)
	if err != nil {
		panic("连接以太坊合约出错")
		return 404, ""
	}

	// 创建合约对象
	//contract_obj,err11 := contract.NewContract(contract_addr,client)
	contract_obj,err11 := putonchain.NewPutonchain(contract_addr,client)
	if err11 !=nil {
		panic("创建合约对象出错")
		return 404, ""
	}
	// 生成调用合约需要的用户
	//auth, err := MakeAuth("123456")
	auth, err := MakeAuth(pass, keystorePath, chainID)
	// 调用合约函数
	//_, err = contract_obj.Hold(auth, "GOD")

	idBigInt,_ := new(big.Int).SetString(id,10)
	tx, err := contract_obj.AddProject(auth, idBigInt, price, label)
	// ethereum/accounts/abi/bind/base.go line:150 显示可以直接使用nil
	//holdStr, err := contract_obj.GetHold(nil)
	//fmt.Println(holdStr)
	//_, err = contract_obj.GetHold(auth)
	if err != nil {
		fmt.Println(err)
		return 404, ""
	}
	return 200, tx.Hash().String()
}

func MakeAuth(pass string, fileName string, chainIDhex string) (*bind.TransactOpts, error) {
	// 生成相关账户的秘钥文件地址
	keystorePath := "/home/guotiezheng/go/src/github.com/ethereum/ethereum/data/keystore"
	//fileName := "UTC--2021-09-15T05-28-00.066630039Z--6a2e96b16eec57ba9f21b1b5c5bdc476a2d3676f"
	file, err := os.Open(keystorePath + "/" + fileName)
	if err != nil {
		fmt.Println("failed to open file ", err)
		return nil, err
	}

	// 初始化big.Int，此值应该在网络中使用eth.chainId()查询得到为16进制
	// chainID,flag := new(big.Int).SetString("22b8",16)
	chainID,flag := new(big.Int).SetString(chainIDhex,16)
	if !flag {
		fmt.Println("get big.int fail")
	}

	// 调用提供的生成结构体的函数，分别传入秘钥文件，密码，链ID
	auth, err := bind.NewTransactorWithChainID(file, pass, chainID)
	if err != nil {
		fmt.Println("failed to NewTransactor  ", err)
		return nil, err
	}
	// 设置参数，完成交易中最多可以使用的gas量
	auth.GasLimit = 300000

	return auth, err
}

