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

func TaskInvoke(stub shim.ChaincodeStubInterface, fn string, args []string) pb.Response {
	switch fn {
	case "ADD":
		return TaskInvokeAdd(stub, args)
	case "MOD":
		return TaskInvokeMod(stub, args)
	case "DEL":
		return TaskInvokeDel(stub, args)
	case "QUERY":
		return TaskInvokeQuery(stub, args)
	default:
		return shim.Error(common.ErrChainCodeFn.Error())
	}
	return shim.Success(nil)
}

func TaskInvokeAdd(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//var task types.Task
	//err := json.Unmarshal([]byte(args), &task)
	//if nil != err {
	//	return shim.Error("param unmarshal failed, err:" + err.Error())
	//}
	//
	//if task.Owner == "" || task.Name == "" || task.Version == "" || task.Data == "" {
	//	return shim.Error(common.ErrParamInvalid.Error())
	//}
	if len(args) != 2 {
		return shim.Error(common.ErrChainCodeArgsNum.Error())
	}
	key := args[0]
	value := args[1]

	err := service.SaveTask(stub, key, []byte(value))
	if err != nil {
		return shim.Error("save task failed, err:" + err.Error())
	}
	res := fmt.Sprintf("add task key %s value %s success", key, value)
	return shim.Success([]byte(res))
}
func TaskInvokeMod(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//var task types.Task
	//err := json.Unmarshal([]byte(args), &task)
	//if nil != err {
	//	return shim.Error("param unmarshal failed, err:" + err.Error())
	//}
	//if task.Owner == "" || task.Name == "" || task.Version == "" {
	//	return shim.Error(common.ErrParamInvalid.Error())
	//}
	if len(args) != 2 {
		return shim.Error(common.ErrChainCodeArgsNum.Error())
	}
	key := args[0]
	value := args[1]

	oldValue, err := service.GetTask(stub, key)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = service.SaveTask(stub, key, []byte(value))
	if err != nil {
		return shim.Error("save task failed, err:" + err.Error())
	}
	res := fmt.Sprintf("mod task key %s value from %s to %s success", key, string(oldValue), value)

	return shim.Success([]byte(res))
}
func TaskInvokeDel(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//var task types.Task
	//err := json.Unmarshal([]byte(args), &task)
	//if nil != err {
	//	return shim.Error("param unmarshal failed, err:" + err.Error())
	//}
	//if task.Owner == "" || task.Name == "" || task.Version == "" {
	//	return shim.Error(common.ErrParamInvalid.Error())
	//}
	if len(args) != 1 {
		return shim.Error(common.ErrChainCodeArgsNum.Error())
	}
	key := args[0]
	err := service.DelTask(stub, key)
	if nil != err {
		shim.Error("del task err:" + err.Error())
	}

	res := fmt.Sprintf("del task %s success", key)
	return shim.Success([]byte(res))
}
func TaskInvokeQuery(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//var task types.Task
	//err := json.Unmarshal([]byte(args), &task)
	//if nil != err {
	//	return shim.Error("param unmarshal failed, err:" + err.Error())
	//}
	//if task.Owner == "" || task.Name == "" || task.Version == "" {
	//	return shim.Error(common.ErrParamInvalid.Error())
	//}
	if len(args) != 1 {
		return shim.Error(common.ErrChainCodeArgsNum.Error())
	}
	key := args[0]
	res, err := service.GetTask(stub, key)
	if err != nil {
		return shim.Error("get task err:" + err.Error())
	}
	return shim.Success(res)
}
