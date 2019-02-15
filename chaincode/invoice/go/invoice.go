package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time" //timestamp

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

// Parameter Use
type Invoice struct {
	InvoiceNumber   string  `json:"invoicenum"`
	BilledTo        string  `json:"billedto"`
	InvoiceDate     string  `json:"invoicedate"`
	InvoiceAmount   float64 `json:"invoiceamount"`
	ItemDescription string  `json:"itemdescription"`
	GR              bool    `json:"gr"`
	IsPaid          bool    `json:"ispaid"`
	PaidAmount      float64 `json:"paidamount"`
	Repaid          bool    `json:"repaid"`
	RepaymentAmount float64 `json:"repaymentamount"`
}

/*
 * The Init method is called when the Smart Contract "invoice" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "raiseInvoice" {
		return s.raiseInvoice(APIstub, args)
	} else if function == "displayAllInvoices" {
		return s.displayAllInvoices(APIstub) //display all invoices without any arguments
	} else if function == "isGoodsReceived" {
		return s.isGoodsReceived(APIstub, args)
	} else if function == "isPaidToSupplier" {
		return s.isPaidToSupplier(APIstub, args)
	} else if function == "isPaidToBank" {
		return s.isPaidToBank(APIstub, args)
	} else if function == "getAuditHistoryForInvoice" {
		return s.getAuditHistoryForInvoice(APIstub, args)
	} else if function == "getUser" {
		return s.getUser(APIstub, args)
	} else if function == "raiseInvoiceWithJsonInput" {
		return s.raiseInvoiceWithJsonInput(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {

	//the first data that is Initialize or register to the network
	invoice := []Invoice{
		Invoice{
			InvoiceNumber:   "1001",
			BilledTo:        "ASUS",
			InvoiceDate:     "07FEB2019",
			InvoiceAmount:   10000.00,
			ItemDescription: "LAPTOP",
			GR:              false,
			IsPaid:          false,
			PaidAmount:      0.00,
			Repaid:          false,
			RepaymentAmount: 0.00},
	}

	i := 0
	for i < len(invoice) {
		fmt.Println("i is ", i)
		invoiceAsBytes, _ := json.Marshal(invoice[i])
		APIstub.PutState("INVOICE"+strconv.Itoa(i), invoiceAsBytes)
		fmt.Println("Added", invoice[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) raiseInvoice(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 { // change size Note: We have only 6 args since other parameter has a default value to initialize and it will automatically generated, those parameters need to be update.
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	invAmount, _ := strconv.ParseFloat(args[4], 64)
	//parameters that will execute data on the insomia api
	var invoice = Invoice{InvoiceNumber: args[1], BilledTo: args[2], InvoiceDate: args[3], InvoiceAmount: invAmount, ItemDescription: args[5], GR: false, IsPaid: false, PaidAmount: 0.00, Repaid: false, RepaymentAmount: 0.00}
	//change parameters

	invoiceAsBytes, _ := json.Marshal(invoice)
	APIstub.PutState(args[0], invoiceAsBytes)

	return shim.Success(nil)
}

// display the all invoices through json format
func (s *SmartContract) displayAllInvoices(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "INVOICE0"
	endKey := "INVOICE999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"INVOICE\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"RECORD\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- displayAllInvoices:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) isGoodsReceived(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	invoiceAsBytes, _ := APIstub.GetState(args[0])
	invoice := Invoice{}

	json.Unmarshal(invoiceAsBytes, &invoice)
	invoice.GR = true

	invoiceAsBytes, _ = json.Marshal(invoice)
	APIstub.PutState(args[0], invoiceAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) isPaidToSupplier(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	invoiceAsBytes, _ := APIstub.GetState(args[0])
	invoice := Invoice{}

	pAmount, _ := strconv.ParseFloat(args[1], 64)

	json.Unmarshal(invoiceAsBytes, &invoice)
	// if the paid amount is less than the invoice amount
	if pAmount < invoice.InvoiceAmount {
		invoice.PaidAmount = pAmount
		invoice.IsPaid = true
	} else {
		return shim.Error("Paid Amount must be always less than invoice amount")
	}
	invoiceAsBytes, _ = json.Marshal(invoice)
	APIstub.PutState(args[0], invoiceAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) isPaidToBank(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	invoiceAsBytes, _ := APIstub.GetState(args[0])
	invoice := Invoice{}

	rAmount, _ := strconv.ParseFloat(args[1], 64)

	json.Unmarshal(invoiceAsBytes, &invoice)
	// if the invoice amount is less than the repayment amount
	if invoice.InvoiceAmount < rAmount {
		invoice.RepaymentAmount = rAmount
		invoice.Repaid = true
	} else {
		return shim.Error("Repayment Amount must be always greater than invoice amount")
	}

	invoiceAsBytes, _ = json.Marshal(invoice)
	APIstub.PutState(args[0], invoiceAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) getAuditHistoryForInvoice(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	invoiceKey := args[0]

	resultsIterator, err := APIstub.GetHistoryForKey(invoiceKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	//Json arrays for Invoice
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":") // The hash of the transaction ID
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		buffer.WriteString(string(response.Value))

		buffer.WriteString(", \"Timestamp\":") // Timestamp that transacts
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) getUser(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	return shim.Success(nil)
}

func main() {

	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
