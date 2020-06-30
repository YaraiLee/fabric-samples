/**
 * @Author: liyalei
 * @Description:
 * @Version:
 * @Date: 2020/6/23 5:02 下午
 */
package service

import (
	"ds-common-contract/common"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func SaveTask(stub shim.ChaincodeStubInterface, key string, data []byte) error {
	err := save(stub, common.TaskPrefix, key, data)
	if err != nil {
		return err
	}
	return nil
}

func GetTask(stub shim.ChaincodeStubInterface, key string) ([]byte, error) {
	return query(stub, common.TaskPrefix, key)
}

func DelTask(stub shim.ChaincodeStubInterface, key string) error {
	var err error

	err = del(stub, common.TaskPrefix, key)
	if err != nil {
		return err
	}

	return nil
}
