/**
 * @Author: liyalei
 * @Description:
 * @Version:
 * @Date: 2020/6/22 3:04 下午
 */
package controller

import (
	"ds-common-contract/common"
	"ds-common-contract/service"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

func NodeInvoke(stub shim.ChaincodeStubInterface, fn string, args []string) pb.Response {
	switch fn {
	case "ADD":
		return NodeInvokeAdd(stub, args)
	case "MOD":
		return NodeInvokeMod(stub, args)
	case "DEL":
		return NodeInvokeDel(stub, args)
	case "QUERY":
		return NodeInvokeQuery(stub, args)
	case "MULTIQUERY":
		return NodeInvokeMultiQuery(stub, args)

	default:
		return shim.Error(common.ErrChainCodeFn.Error())
	}
	return shim.Success(nil)
}

func NodeInvokeAdd(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//var node types.NodeItem

	//err := json.Unmarshal([]byte(args), &node)
	//if nil != err {
	//	return shim.Error("param unmarshal failed, err:" + err.Error())
	//}
	//if "" == node.OrgName || "" == node.NodeName {
	//	return shim.Error(common.ErrParamInvalid.Error())
	//}
	if len(args) != 2 {
		return shim.Error(common.ErrChainCodeArgsNum.Error())
	}
	key := args[0]
	value := args[1]
	err := service.SaveNode(stub, key, []byte(value))
	if nil != err {
		return shim.Error("save node failed, err:" + err.Error())
	}

	res := fmt.Sprintf("add node key %s value %s success", key, value)
	return shim.Success([]byte(res))
}

func NodeInvokeDel(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//var node types.NodeItem
	//
	//err := json.Unmarshal([]byte(args), &node)
	//if nil != err {
	//	return shim.Error("param unmarshal failed, err:" + err.Error())
	//}
	//if "" == node.OrgName || "" == node.NodeName {
	//	return shim.Error(common.ErrParamInvalid.Error())
	//}
	if len(args) != 1 {
		return shim.Error(common.ErrChainCodeArgsNum.Error())
	}
	err := service.DelNode(stub, args[0])
	if err != nil {
		return shim.Error("del node err:" + err.Error())
	}

	res := fmt.Sprintf("del key %s node success", args[0])
	return shim.Success([]byte(res))
}

func NodeInvokeMod(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//var node types.NodeItem
	//err := json.Unmarshal([]byte(args), &node)
	//if err != nil {
	//	return shim.Error("param unmarshal failed, err:" + err.Error())
	//}
	//if "" == node.OrgName || "" == node.NodeName {
	//	return shim.Error(common.ErrParamInvalid.Error())
	//}
	if len(args) != 2 {
		return shim.Error(common.ErrChainCodeArgsNum.Error())
	}
	key := args[0]
	value := args[1]
	err := service.SaveNode(stub, key, []byte(value))
	if nil != err {
		return shim.Error("save node failed, err:" + err.Error())
	}

	oldValue, err := service.GetNode(stub, key)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = service.SaveNode(stub, key, []byte(value))
	if err != nil {
		return shim.Error("save db err:" + err.Error())
	}
	res := fmt.Sprintf("modify node key %s value from %s to  %s success", key, string(oldValue), value)
	return shim.Success([]byte(res))
}

func NodeInvokeQuery(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//var node types.NodeItem
	//err := json.Unmarshal([]byte(args), &node)
	//if err != nil {
	//	return shim.Error("param unmarshal failed, err:" + err.Error())
	//}
	//if "" == node.OrgName || "" == node.NodeName {
	//	return shim.Error(common.ErrParamInvalid.Error())
	//}
	if len(args) != 1 {
		return shim.Error(common.ErrChainCodeArgsNum.Error())
	}
	key := args[0]

	res, err := service.GetNode(stub, key)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(res)
}

func NodeInvokeMultiQuery(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//var node types.NodeItem
	//err := json.Unmarshal([]byte(args), &node)
	//if err != nil {
	//	return shim.Error("param unmarshal failed, err:" + err.Error())
	//}
	//if node.OrgName == "" {
	//	return shim.Error(common.ErrParamInvalid.Error())
	//}
	//var keys []string
	//keys = append(keys, node.OrgName)
	//if node.NodeName != "" {
	//	keys = append(keys, node.NodeName)
	//}
	if len(args) != 1 {
		return shim.Error(common.ErrChainCodeArgsNum.Error())
	}
	res, err := service.GetNodes(stub, args[0])
	if nil != err {
		return shim.Error(err.Error())
	}
	return shim.Success(res)
}
