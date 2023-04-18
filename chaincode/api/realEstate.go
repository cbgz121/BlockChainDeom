package api

import (
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/json"
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// CreateRealEstate 新建房地产
func CreateRealEstate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 10 {
		return shim.Error("参数个数不满足")
	}
	accountId := args[0] //accountId用于验证是否为管理员
	proprietor := args[1]
	totalArea := args[2]
	livingSpace := args[3]
	estateNumber := args[4]
	estateAddress := args[5]
	buildYear   := args[6]
	estateType  := args[7]
	estateStatus := args[8]
        phone        := args[9]
	if accountId == "" || proprietor == "" || totalArea == "" || livingSpace == "" || estateNumber == "" || estateAddress == "" || buildYear == "" || estateType == "" || estateStatus == "" {
		return shim.Error("参数存在空值")
	}
	if accountId != proprietor {
		return shim.Error("操作人应为本人")
	}
	// 参数数据格式转换
	var formattedTotalArea float64
	if val, err := strconv.ParseFloat(totalArea, 64); err != nil {
		return shim.Error(fmt.Sprintf("totalArea参数格式转换出错: %s", err))
	} else {
		formattedTotalArea = val
	}
	var formattedLivingSpace float64
	if val, err := strconv.ParseFloat(livingSpace, 64); err != nil {
		return shim.Error(fmt.Sprintf("livingSpace参数格式转换出错: %s", err))
	} else {
		formattedLivingSpace = val
	}
	//判断业主是否存在
	resultsProprietor, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, []string{proprietor})
	if err != nil || len(resultsProprietor) != 1 {
		return shim.Error(fmt.Sprintf("业主proprietor信息验证失败%s", err))
	}
	combined := args[1] +args[2] + args[3] + args[4] + args[5] + args[6] + args[7] + args[8]

	// 使用SHA-256算法进行hash
	hashed := sha256.Sum256([]byte(combined))
	// 将hash结果转换为16进制字符串
    hashStr := fmt.Sprintf("%x", hashed)
	hashStr = hashStr[:16]
	realEstate := &model.RealEstate{
		RealEstateID: hashStr,
		Proprietor:   proprietor,
		Encumbrance:  false,
		TotalArea:    formattedTotalArea,
		LivingSpace:  formattedLivingSpace,
		EstateNumber: estateNumber,
		EstateAddress: estateAddress,
		BuildYear: buildYear,
		EstateType: estateType,
		EstateStatus: estateStatus,
                Phone: phone,
	}
	// 写入账本
	if err := utils.WriteLedger(realEstate, stub, model.RealEstateKey, []string{realEstate.Proprietor, realEstate.RealEstateID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	realEstateByte, err := json.Marshal(realEstate)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(realEstateByte)
}

// QueryRealEstateList 查询房地产(可查询所有，也可根据所有人查询名下房产)
func QueryRealEstateList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var realEstateList []model.RealEstate
	results, err := utils.GetStateByPartialCompositeKeys2(stub, model.RealEstateKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var realEstate model.RealEstate
			err := json.Unmarshal(v, &realEstate)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryRealEstateList-反序列化出错: %s", err))
			}
			realEstateList = append(realEstateList, realEstate)
		}
	}
	realEstateListByte, err := json.Marshal(realEstateList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryRealEstateList-序列化出错: %s", err))
	}
	return shim.Success(realEstateListByte)
}
