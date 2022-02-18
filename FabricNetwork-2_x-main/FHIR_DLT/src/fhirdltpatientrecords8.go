/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-contract-api-go/metadata"
)

type Fhirdlt struct {
	Value string `json:"value"`
}

type FHIR_Compliant_Patient struct {
	Resourcetype string    `json:"Resourcetype"`
	Name         []Name    `json:"Name"`
	Gender       string    `json:"Gender"`
	BirthDate    string    `json:"BirthDate"`
	Telecom      []Telecom `json:"Telecom"`
	Address      []Address `json:"Address"`
	Id           string    `json:"Id"`
	Active       bool      `json:"active"`
}

type Name struct {
	Use    string `json:"Use"`
	Family string `json:"Family"`
	Given  string `json:"Given"`
}

type Telecom struct {
	Value  string `json:"Value"`
	Use    string `json:"Use"`
	System string `json:"System"`
}

type Address struct {
	Use        string `json:"Use"`
	Line       string `json:"Line"`
	City       string `json:"City"`
	State      string `json:"State"`
	Country    string `json:"Country"`
	PostalCode string `json:"PostalCode"`
}

// FhirdltpatientContract contract for managing CRUD for Fhirdltpatient
type FhirdltpatientContract struct {
	contractapi.Contract
}

type CounterNO struct {
	Counter int `json:"counter"`
}

// FhirdltpatientExists returns true when asset with given ID exists in world state
func (c *FhirdltpatientContract) FhirdltpatientExists(ctx contractapi.TransactionContextInterface, fhirdltpatientID string) (bool, error) {
	data, err := ctx.GetStub().GetState(fhirdltpatientID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateFhirdltpatient creates a new instance of Fhirdltpatient
func (c *FhirdltpatientContract) CreatePatientFhirdlt(ctx contractapi.TransactionContextInterface, input string) (string, error) {

	if len(input) == 0 {
		return "", fmt.Errorf("Please pass the correct  data")
	}

	var FHIR_Patient FHIR_Compliant_Patient
	err := json.Unmarshal([]byte(input), &FHIR_Patient)
	if err != nil {
		return "", fmt.Errorf("Failed while unmarshling records. %s", err.Error())
	}

	PatientAsBytes, err := json.Marshal(FHIR_Patient)
	if err != nil {
		return "", fmt.Errorf("Failed while marshling records. %s", err.Error())
	}
	counter := GetCounter(ctx)
	var newid = FHIR_Patient.Resourcetype + "_" + FHIR_Patient.BirthDate + "_" + counter

	ctx.GetStub().SetEvent("CreatePatientRecordAsset", PatientAsBytes)
	return newid, ctx.GetStub().PutState(newid, PatientAsBytes)

	// exists, err := c.FhirdltpatientExists(ctx, id)
	// if err != nil {
	// 	return fmt.Errorf("Could not read from world state. %s", err)
	// } else if exists {
	// 	return fmt.Errorf("The asset %s already exists", id)
	// }
	// counter := GetCounter(ctx)
	// var newid = resourcetype + "_" + birthDate + "_" + counter

	// //int_val++
	// str1 := "`"
	// str2 := "`"
	// valuestrname := str1 + name + str2
	// telecomArr := str1 + telecom + str2
	// addressArr := str1 + address + str2

	// FHIR_Compliant_Patient_OBJ := FHIR_Compliant_Patient{}

	// namearray := []Name{}
	// var val []byte = []byte(valuestrname)
	// name_array, _ := strconv.Unquote(string(val))
	// err6 := json.Unmarshal([]byte(name_array), &namearray)
	// if err6 != nil {
	// 	//return "", fmt.Errorf("Could not read from world state. %s", err6)
	// }

	// fmt.Println(namearray)
	// fmt.Println(name_array)

	// telecomarray := []Telecom{}
	// var val_telecom []byte = []byte(telecomArr)
	// telecom_array, _ := strconv.Unquote(string(val_telecom))
	// err2 := json.Unmarshal([]byte(telecom_array), &telecomarray)
	// if err2 != nil {

	// }

	// addressarray := []Address{}
	// var val_address []byte = []byte(addressArr)
	// address_array, _ := strconv.Unquote(string(val_address))
	// err3 := json.Unmarshal([]byte(address_array), &addressarray)

	// if err3 != nil {

	// }

	// FHIR_Compliant_Patient_OBJ.Resourcetype = resourcetype
	// FHIR_Compliant_Patient_OBJ.Name = namearray
	// FHIR_Compliant_Patient_OBJ.Gender = gender
	// FHIR_Compliant_Patient_OBJ.BirthDate = birthDate
	// FHIR_Compliant_Patient_OBJ.Telecom = telecomarray
	// FHIR_Compliant_Patient_OBJ.Address = addressarray
	// FHIR_Compliant_Patient_OBJ.Id = newid
	// FHIR_Compliant_Patient_OBJ.Active = active

	// bytes_1, _ := json.Marshal(FHIR_Compliant_Patient_OBJ)
	// ctx.GetStub().SetEvent("CreatePatientRecordAsset", bytes_1)

	// return newid, ctx.GetStub().PutState(newid, bytes_1)

}

// ReadFhirdlt retrieves an instance of Fhirdlt from the world state
func (c *FhirdltpatientContract) ReadFhirdlt(ctx contractapi.TransactionContextInterface, fhirdltID string) (*FHIR_Compliant_Patient, error) {
	exists, err := c.FhirdltpatientExists(ctx, fhirdltID)
	if err != nil {
		return nil, fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return nil, fmt.Errorf("The asset %s does not exist", fhirdltID)
	}

	bytes, _ := ctx.GetStub().GetState(fhirdltID)

	fhirdlt_Val := new(FHIR_Compliant_Patient)

	err = json.Unmarshal(bytes, fhirdlt_Val)

	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal world state data to type Fhirdlt")
	}

	return fhirdlt_Val, nil
}

// UpdateFhirdltpatient retrieves an instance of Fhirdltpatient from the world state and updates its value
// func (c *FhirdltpatientContract) UpdateFhirdltpatient(ctx contractapi.TransactionContextInterface, fhirdltpatientID string, newValue string) error {
// 	exists, err := c.FhirdltpatientExists(ctx, fhirdltpatientID)
// 	if err != nil {
// 		return fmt.Errorf("Could not read from world state. %s", err)
// 	} else if !exists {
// 		return fmt.Errorf("The asset %s does not exist", fhirdltpatientID)
// 	}

// 	fhirdltpatient := new(Fhirdltpatient)
// 	fhirdltpatient.Value = newValue

// 	bytes, _ := json.Marshal(fhirdltpatient)

// 	return ctx.GetStub().PutState(fhirdltpatientID, bytes)
// }

// DeleteFhirdltpatient deletes an instance of Fhirdltpatient from the world state
func (c *FhirdltpatientContract) DeleteFhirdltpatient(ctx contractapi.TransactionContextInterface, fhirdltpatientID string) error {
	exists, err := c.FhirdltpatientExists(ctx, fhirdltpatientID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fhirdltpatientID)
	}

	return ctx.GetStub().DelState(fhirdltpatientID)
}

func GetCounter(ctx contractapi.TransactionContextInterface) string {
	counterAsBytes, _ := ctx.GetStub().GetState("CounterNo")
	counterAsset := CounterNO{}

	err := json.Unmarshal(counterAsBytes, &counterAsset)
	counterAsset.Counter++
	counterAsBytes, _ = json.Marshal(counterAsset)

	err = ctx.GetStub().PutState("CounterNo", counterAsBytes)
	if err != nil {

		fmt.Sprintf("Failed to Increment Counter")

	}
	return strconv.Itoa(counterAsset.Counter)
}

func main() {
	fhirdltpatientContract := new(FhirdltpatientContract)
	fhirdltpatientContract.Info.Version = "0.0.1"
	fhirdltpatientContract.Info.Description = "My Smart Contract"
	fhirdltpatientContract.Info.License = new(metadata.LicenseMetadata)
	fhirdltpatientContract.Info.License.Name = "Apache-2.0"
	fhirdltpatientContract.Info.Contact = new(metadata.ContactMetadata)
	fhirdltpatientContract.Info.Contact.Name = "John Doe"

	chaincode, err := contractapi.NewChaincode(fhirdltpatientContract)
	chaincode.Info.Title = "chaincode chaincode"
	chaincode.Info.Version = "0.0.1"

	if err != nil {
		panic("Could not create chaincode from FhirdltpatientContract." + err.Error())
	}

	err = chaincode.Start()

	if err != nil {
		panic("Failed to start chaincode. " + err.Error())
	}
}
