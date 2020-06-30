/**
 * @Author: liyalei
 * @Description:
 * @Version:
 * @Date: 2020/6/22 3:29 下午
 */
package service

import (
	"crypto/x509"
	"ds-common-contract/common"
	"encoding/pem"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/msp"
)

func GetCertX509(stub shim.ChaincodeStubInterface) (*x509.Certificate, error) {
	creatorCert, err := stub.GetCreator()
	if err != nil {
		return nil, err
	}

	identity := msp.SerializedIdentity{}
	err = proto.Unmarshal(creatorCert, &identity)
	if err != nil {
		return nil, err
	}

	cert, _ := pem.Decode(identity.IdBytes)
	cert509, err := x509.ParseCertificate(cert.Bytes)
	if err != nil {
		return nil, err
	}

	return cert509, nil
}

func save(stub shim.ChaincodeStubInterface, prefix string, key string, data []byte) error {
	//var bs []byte
	uniKey, err := stub.CreateCompositeKey(prefix, []string{key})
	if err != nil {
		return err
	}
	////结构体转json字符串
	//bs, err = json.Marshal(data)
	//if err != nil {
	//	return nil, err
	//}
	//保存
	err = stub.PutState(uniKey, data)
	if err != nil {
		return err
	}
	return nil
}

func query(stub shim.ChaincodeStubInterface, prefix string, key string) ([]byte, error) {
	var result []byte
	uniKey, err := stub.CreateCompositeKey(prefix, []string{key})
	if err != nil {
		return result, err
	}
	data, err := stub.GetState(uniKey)
	if err != nil {
		return result, err
	} else if data == nil || len(data) == 0 {
		return result, common.ErrNotFound
	}

	return data, nil
}

func del(stub shim.ChaincodeStubInterface, prefix string, key string) error {
	var err error
	uniKey, err := stub.CreateCompositeKey(prefix, []string{key})
	if err != nil {
		return err
	}

	//删除
	err = stub.DelState(uniKey)
	if err != nil {
		return err
	}
	return nil
}

func store(stub shim.ChaincodeStubInterface, prefix string, key string, data []byte) ([]byte, error) {
	uniKey, err := stub.CreateCompositeKey(prefix, []string{key})
	if err != nil {
		return nil, err
	}
	//保存
	err = stub.PutState(uniKey, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func queryAll(stub shim.ChaincodeStubInterface, prefix string, key string) ([][]byte, error) {
	result := [][]byte{}
	resultsIterator, err := stub.GetStateByPartialCompositeKey(prefix, []string{key})
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	for resultsIterator.HasNext() {
		res, err := resultsIterator.Next()
		if nil != err {
			return nil, err
		}
		result = append(result, res.Value)
	}
	return result, nil
}
