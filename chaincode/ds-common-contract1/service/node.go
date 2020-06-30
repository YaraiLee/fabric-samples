package service

import (
	"ds-common-contract/common"
	"ds-common-contract/types"
	"github.com/hyperledger/fabric-chaincode-go/shim"

	"encoding/json"
)

func SaveNode(stub shim.ChaincodeStubInterface, key string, data []byte) error {
	err := save(stub, common.NodePrefix, key, data)
	if err != nil {
		return err
	}
	return nil
}

func GetNode(stub shim.ChaincodeStubInterface, key string) ([]byte, error) {
	return query(stub, common.NodePrefix, key)
}

func DelNode(stub shim.ChaincodeStubInterface, key string) error {
	err := del(stub, common.NodePrefix, key)
	if err != nil {
		return err
	}

	return nil
}

func GetNodes(stub shim.ChaincodeStubInterface, key string) ([]byte, error) {
	var result types.QueryOrgResult
	resultsIterator, err := stub.GetStateByPartialCompositeKey(common.NodePrefix, []string{key})
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	for resultsIterator.HasNext() {
		res, err := resultsIterator.Next()
		if nil != err {
			return nil, err
		}
		var node types.NodeItem
		err = json.Unmarshal(res.Value, &node)
		if nil != err {
			continue
		}
		result.Records = append(result.Records, node)
	}
	return json.Marshal(result)
}
