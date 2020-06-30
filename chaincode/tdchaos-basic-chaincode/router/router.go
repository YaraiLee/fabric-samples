/**
 * @Author: liyalei
 * @Description:
 * @Version:
 * @Date: 2020/6/22 2:42 下午
 */
package router

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"tdchaos-basic-chaincode/controller"
)

type BasicChainCode struct {
}

func (b *BasicChainCode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success([]byte("Init Success"))
}

func (b *BasicChainCode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fn, args := stub.GetFunctionAndParameters()
	//if len(args) != 2 {
	//	return shim.Error(common.ErrChainCodeArgsNum.Error() + ", len(args):" + strconv.Itoa(len(args)))
	//}
	switch fn {
	case "Node":
		return controller.NodeInvoke(stub, args[0], args[1:])
	case "Task":
		return controller.TaskInvoke(stub, args[0], args[1:])
	}
	return shim.Error("arguments err, undefined module: " + fn)
}
