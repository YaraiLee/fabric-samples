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

func NodeInvoke(stub shim.ChaincodeStubInterface, fn string, args []string) pb.Response {
	switch fn {
	case "ADD":
		return service.InvokeAdd(stub, common.NodePrefix, args)
	case "MOD":
		return service.InvokeMod(stub, common.NodePrefix, args)
	case "DEL":
		return service.InvokeDel(stub, common.NodePrefix, args)
	case "QUERY":
		return service.InvokeQuery(stub, common.NodePrefix, args)
	default:
		return shim.Error(common.ErrChainCodeFn.Error())
	}
}
