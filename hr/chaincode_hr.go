/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}
// details map[string]interface{}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, id []string) ([]byte, error) {
	if len(id) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	// mapString := make(map[string]string)

	// for key, value := range details {
 //        strKey := fmt.Sprintf("%v", key)
 //        strValue := fmt.Sprintf("%v", value)

 //        mapString[strKey] = strValue
 //    }

	err := stub.PutState(id[0], []byte(id[1]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (t *SimpleChaincode) Update(stub shim.ChaincodeStubInterface, function string, id []string) ([]byte, error) {
	// var err error
	var key string

	if len(id) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	key = id[0]

	valAsbytes, _ := stub.GetState(key)

	return valAsbytes, nil
}


// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, id []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "Init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "Init", id)
	}
	if function == "Update" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "Update", id)
	}
	fmt.Println("invoke did not find func: " + function)					//error

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, id []string) ([]byte, error) {
	var jsonResp string
	var key string
	var err error

	if len(id) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	key = id[0]

	valAsbytes, err := stub.GetState(key)

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for \"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil					//error

	// return nil, errors.New("Received unknown function query: " + function)
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
