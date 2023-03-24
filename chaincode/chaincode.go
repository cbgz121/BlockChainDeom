package main

import (
	"chaincode/api"
	"chaincode/model"
	"chaincode/pkg/utils"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type BlockChainRealEstate struct {
}

// Init 链码初始化
func (t *BlockChainRealEstate) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("链码初始化")
	//初始化默认数据
	// var userNames = [6]string{"admin"}
	// var balances = [6]float64{0}
	//初始化账号数据
	// 在这里可以对前端传来的数据进行处理，例如将密码加密后存入数据库等等
	// 将用户名和密码合并成一个字符串
	combined := "admin" + "123456"

	// 使用SHA-256算法进行hash
	hashed := sha256.Sum256([]byte(combined))
	// 将hash结果转换为16进制字符串
    hashStr := fmt.Sprintf("%x", hashed)
	hashStr = hashStr[:16]
	account := &model.Account{
		AccountId: hashStr,
		UserName:  "admin",
		Balance:   0,
	}
	res := []string{hashStr}
		// 写入账本
	if err := utils.WriteLedger(account, stub, model.AccountKey, res); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(nil)
}

// Invoke 实现Invoke接口调用智能合约
func (t *BlockChainRealEstate) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "hello":
		return api.Hello(stub, args)
	case "queryAccountList":
		return api.QueryAccountList(stub, args)
	case "createRealEstate":
		return api.CreateRealEstate(stub, args)
	case "queryRealEstateList":
		return api.QueryRealEstateList(stub, args)
	case "createSelling":
		return api.CreateSelling(stub, args)
	case "createSellingByBuy":
		return api.CreateSellingByBuy(stub, args)
	case "querySellingList":
		return api.QuerySellingList(stub, args)
	case "querySellingListByBuyer":
		return api.QuerySellingListByBuyer(stub, args)
	case "updateSelling":
		return api.UpdateSelling(stub, args)
	case "createDonating":
		return api.CreateDonating(stub, args)
	case "queryDonatingList":
		return api.QueryDonatingList(stub, args)
	case "queryDonatingListByGrantee":
		return api.QueryDonatingListByGrantee(stub, args)
	case "updateDonating":
		return api.UpdateDonating(stub, args)
	case "register":
		return api.Register(stub, args)
	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	timeLocal, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = timeLocal
	err = shim.Start(new(BlockChainRealEstate))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
