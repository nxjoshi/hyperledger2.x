/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

func main() {
	fhirdltContract := new(FhirdltContract)
	fhirdltContract.Info.Version = "0.0.1"
	fhirdltContract.Info.Description = "My Smart Contract"
	fhirdltContract.Info.License = new(metadata.LicenseMetadata)
	fhirdltContract.Info.License.Name = "Apache-2.0"
	fhirdltContract.Info.Contact = new(metadata.ContactMetadata)
	fhirdltContract.Info.Contact.Name = "John Doe"

	chaincode, err := contractapi.NewChaincode(fhirdltContract)
	chaincode.Info.Title = "V2 chaincode"
	chaincode.Info.Version = "0.0.1"

	if err != nil {
		panic("Could not create chaincode from FhirdltContract." + err.Error())
	}

	err = chaincode.Start()

	if err != nil {
		panic("Failed to start chaincode. " + err.Error())
	}
}
