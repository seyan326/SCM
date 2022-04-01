/*
Copyright IBM Corp. 2016 All Rights Reserved.

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
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// impix_scm Chaincode implementation
type impix_scm struct {
	contractapi.Contract
}

type OrderBook struct {
	OrderNumber         string `json:"orderNumber"`
	Compuny             string `json:"compuny"`
	ProductNumber       string `json:"productNumber"`
	ProductName         string `json:"productName"`
	Thickness           string `json:"thickness"`
	Width               string `json:"width"`
	Division            string `json:"division"`
	OrderQuantity       string `json:"orderQuantity"`
	Unit                string `json:"unit"`
	DeliveryRequestDate string `json:"deliveryRequestDate"`
	Destination         string `json:"destination"`
}

func (s *impix_scm) CreateOrderBook(ctx contractapi.TransactionContextInterface,
	orderNumber string, compuny string, productNumber string, productName string,
	thickness string, width string, division string, orderQuantity string, unit string,
	deliveryRequestDate string, destination string) error {

	orderBook := OrderBook{
		OrderNumber:         orderNumber,
		Compuny:             compuny,
		ProductNumber:       productNumber,
		ProductName:         productName,
		Thickness:           thickness,
		Width:               width,
		Division:            division,
		OrderQuantity:       orderQuantity,
		Unit:                unit,
		DeliveryRequestDate: deliveryRequestDate,
		Destination:         destination,
	}
	memberAsBytes, _ := json.Marshal(orderBook)

	return ctx.GetStub().PutState(orderNumber, memberAsBytes)
}

// QueryOrder
func (s *impix_scm) QueryOrder(ctx contractapi.TransactionContextInterface, orderNumber string) (*OrderBook, error) {
	orderAsBytes, err := ctx.GetStub().GetState(orderNumber)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if orderAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", orderNumber)
	}

	orderBook := new(OrderBook)
	_ = json.Unmarshal(orderAsBytes, orderBook)

	return orderBook, nil
}

func main() {
	cc, err := contractapi.NewChaincode(new(impix_scm))
	if err != nil {
		panic(err.Error())
	}
	if err := cc.Start(); err != nil {
		fmt.Printf("Error starting impix_scm chaincode: %s", err)
	}
}
