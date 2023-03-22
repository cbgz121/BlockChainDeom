package api

import (
	"crypto/sha256"
	"fmt"
	"encoding/json"

	"chaincode/model"
	"chaincode/pkg/utils"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func Register(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 在这里可以对前端传来的数据进行处理，例如将密码加密后存入数据库等等
	// 将用户名和密码合并成一个字符串
	combined := args[0] + args[1]

	// 使用SHA-256算法进行hash
	hashed := sha256.Sum256([]byte(combined))
	// 将hash结果转换为16进制字符串
    hashStr := fmt.Sprintf("%x", hashed)

	account := &model.Account{
		AccountId: hashStr,
		UserName:  args[0],
		Balance:   5000000,
	}
	
	// 写入账本
	if err := utils.WriteLedger(account, stub, model.AccountKey, []string{hashStr}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	registerByte, err := json.Marshal(account)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(registerByte)
}