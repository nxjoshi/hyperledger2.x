/*
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Fhirdlt struct {
	Value string `json:"value"`
}

type FHIR_Compliant_Patient struct {
	Resourcetype string    `json:"Resourcetype"`
	Name         []Name    `json:"NameArray"`
	Gender       string    `json:"Gender"`
	BirthDate    string    `json:"BirthDate"`
	Telecom      []Telecom `json:"TelecomArray"`
	Address      []Address `json:"AddressArray"`
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

// FhirdltContract contract for managing CRUD for Fhirdlt
type FhirdltContract struct {
	contractapi.Contract
}

// FhirdltExists returns true when asset with given ID exists in world state
func (c *FhirdltContract) FhirdltExists(ctx contractapi.TransactionContextInterface, fhirdltID string) (bool, error) {
	data, err := ctx.GetStub().GetState(fhirdltID)

	if err != nil {
		return false, err
	}

	return data != nil, nil
}

// CreateFhirdlt creates a new instance of Fhirdlt

/*	Resourcetype string,Name []Name,Gender string,BirthDate string,Telecom []Telecom,Address []Address,Id string,Active bool */
func (c *FhirdltContract) CreatePatientFhirdlt(ctx contractapi.TransactionContextInterface, resourcetype string, name string, gender string, birthDate string, telecom string, address string, period string, id string, active bool) error {
	exists, err := c.FhirdltExists(ctx, id)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if exists {
		return fmt.Errorf("The asset %s already exists", id)
	}

	str1 := "`"
	str2 := "`"
	valuestrname := str1 + name + str2
	telecomArr := str1 + telecom + str2
	addressArr := str1 + address + str2

	FHIR_Compliant_Patient := FHIR_Compliant_Patient{}

	namearray := []Name{}
	var val []byte = []byte(valuestrname)
	name_array, _ := strconv.Unquote(string(val))
	err1 := json.Unmarshal([]byte(name_array), &namearray)
	if err1 != nil {

	}

	fmt.Println(namearray)
	fmt.Println(name_array)

	telecomarray := []Telecom{}
	var val_telecom []byte = []byte(telecomArr)
	telecom_array, _ := strconv.Unquote(string(val_telecom))
	err2 := json.Unmarshal([]byte(telecom_array), &telecomarray)
	if err2 != nil {

	}

	addressarray := []Address{}
	var val_address []byte = []byte(addressArr)
	address_array, _ := strconv.Unquote(string(val_address))
	err3 := json.Unmarshal([]byte(address_array), &addressarray)

	if err3 != nil {

	}

	// period_value := Period{}
	// var val_period []byte = []byte(period)
	// period_val, _ := strconv.Unquote(string(val_period))
	// err4 := json.Unmarshal([]byte(period_val), &period_value)

	// if err4 != nil {

	// }

	FHIR_Compliant_Patient.Resourcetype = resourcetype
	FHIR_Compliant_Patient.Name = namearray
	FHIR_Compliant_Patient.Gender = gender
	FHIR_Compliant_Patient.BirthDate = birthDate
	FHIR_Compliant_Patient.Telecom = telecomarray
	FHIR_Compliant_Patient.Address = addressarray

	bytes, _ := json.Marshal(FHIR_Compliant_Patient)

	return ctx.GetStub().PutState(id, bytes)
}

// ReadFhirdlt retrieves an instance of Fhirdlt from the world state
func (c *FhirdltContract) ReadFhirdlt(ctx contractapi.TransactionContextInterface, fhirdltID string) (*FHIR_Compliant_Patient, error) {
	exists, err := c.FhirdltExists(ctx, fhirdltID)
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

// UpdateFhirdlt retrieves an instance of Fhirdlt from the world state and updates its value
func (c *FhirdltContract) UpdateFhirdlt(ctx contractapi.TransactionContextInterface, fhirdltID string, newValue string) error {
	exists, err := c.FhirdltExists(ctx, fhirdltID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fhirdltID)
	}

	fhirdlt := new(Fhirdlt)
	fhirdlt.Value = newValue

	bytes, _ := json.Marshal(fhirdlt)

	return ctx.GetStub().PutState(fhirdltID, bytes)
}

// DeleteFhirdlt deletes an instance of Fhirdlt from the world state
func (c *FhirdltContract) DeleteFhirdlt(ctx contractapi.TransactionContextInterface, fhirdltID string) error {
	exists, err := c.FhirdltExists(ctx, fhirdltID)
	if err != nil {
		return fmt.Errorf("Could not read from world state. %s", err)
	} else if !exists {
		return fmt.Errorf("The asset %s does not exist", fhirdltID)
	}

	return ctx.GetStub().DelState(fhirdltID)
}
