/**
 * @Author: liyalei
 * @Description:
 * @Version:
 * @Date: 2020/6/24 4:51 下午
 */
package service

import (
	"encoding/json"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"tdchaos-basic-chaincode/common"
)

func InvokeAdd(stub shim.ChaincodeStubInterface, prefix string, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error(prefix + "ADD " + common.ErrChainCodeArgsNum.Error())
	}
	value := []byte(args[0])
	keys := args[1:]

	uniqKey, err := stub.CreateCompositeKey(prefix, keys)
	if nil != err {
		return shim.Error(err.Error())
	}

	err = stub.PutState(uniqKey, value)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func InvokeMod(stub shim.ChaincodeStubInterface, prefix string, args []string) pb.Response {
	if len(args) < 2 {
		return shim.Error(prefix + "MOD " + common.ErrChainCodeArgsNum.Error())
	}
	value := []byte(args[0])
	keys := args[1:]
	uniqKey, err := stub.CreateCompositeKey(prefix, keys)
	if nil != err {
		return shim.Error(err.Error())
	}
	oldValue, err := stub.GetState(uniqKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	if oldValue == nil || len(oldValue) == 0 {
		return shim.Error(common.ErrNotFound.Error())
	}

	err = stub.PutState(uniqKey, value)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func InvokeDel(stub shim.ChaincodeStubInterface, prefix string, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error(prefix + "DEL " + common.ErrChainCodeArgsNum.Error())
	}
	uniqKey, err := stub.CreateCompositeKey(prefix, args)
	if nil != err {
		return shim.Error(err.Error())
	}
	err = stub.DelState(uniqKey)
	if nil != err {
		shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func InvokeQuery(stub shim.ChaincodeStubInterface, prefix string, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error(prefix + "QUERY " + common.ErrChainCodeArgsNum.Error())
	}

	var result []string
	iterator, err := stub.GetStateByPartialCompositeKey(prefix, args)
	if nil != err {
		return shim.Error(err.Error())
	}

	for iterator.HasNext() {
		res, err := iterator.Next()
		if nil != err {
			return shim.Error(err.Error())
		}
		result = append(result, string(res.Value))
	}
	res, err := json.Marshal(result)
	if nil != err {
		return shim.Error(err.Error())
	}
	return shim.Success(res)
}
