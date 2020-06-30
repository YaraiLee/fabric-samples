/**
 * @Author: liyalei
 * @Description:
 * @Version:
 * @Date: 2020/6/22 3:04 下午
 */
package controller

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"tdchaos-basic-chaincode/common"
	"tdchaos-basic-chaincode/service"
)

func TaskInvoke(stub shim.ChaincodeStubInterface, fn string, args []string) pb.Response {
	switch fn {
	case "ADD":
		return service.InvokeAdd(stub, common.TaskPrefix, args)
	case "MOD":
		return service.InvokeMod(stub, common.TaskPrefix, args)
	case "DEL":
		return service.InvokeDel(stub, common.TaskPrefix, args)
	case "QUERY":
		return service.InvokeQuery(stub, common.TaskPrefix, args)
	default:
		return shim.Error(common.ErrChainCodeFn.Error())
	}
	return shim.Success(nil)
}
