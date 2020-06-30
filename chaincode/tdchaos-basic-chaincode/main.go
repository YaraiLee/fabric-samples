/**
 * @Author: liyalei
 * @Description:
 * @Version:
 * @Date: 2020/6/22 5:46 下午
 */
package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"tdchaos-basic-chaincode/router"
)

func main() {
	err := shim.Start(new(router.BasicChainCode))
	if nil != err {
		fmt.Printf("Error create BasicChainCode: %s", err)
	}
}
