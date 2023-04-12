package api

import (
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/json"
	"crypto/sha256"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// QueryAccountList 查询账户列表
func QueryAccountList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var accountList []model.Account
	// 将用户名和密码合并成一个字符串
	combined := args[0] + args[1]
	res := make([]string,0)
	if combined != "" {
		// 使用SHA-256算法进行hash
		hashed := sha256.Sum256([]byte(combined))
		// 将hash结果转换为16进制字符串
		hashStr := fmt.Sprintf("%x", hashed)
		hashStr = hashStr[:16]
		res = append(res,hashStr)
	}
	results, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, res)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var account model.Account
			err := json.Unmarshal(v, &account)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
			}
			accountList = append(accountList, account)
		}
	}
	accountListByte, err := json.Marshal(accountList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
	}
	return shim.Success(accountListByte)
}
