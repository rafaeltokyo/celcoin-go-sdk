package celcoin

import (
	"encoding/json"
	"fmt"
)

//BankslipRequest - Modelo para criação de um boleto
type BankslipRequest struct {
	Payer        BankslipPayer               `json:"payer"`
	Instructions BankslipInstructionsRequest `json:"instructions"`
	DueDate      string                      `json:"dueDate"`
	Values       BankslipValues              `json:"values"`
	DaysToExpire int32                       `json:"daysToExpire"`
}

//BankslipReceipt - Dados do recebedor
type BankslipReceipt struct {
	ReceiptData      string `json:"receiptData"`
	Receiptformatted string `json:"receiptformatted"`
}

//BankslipResponse - Resposta na geração do boleto
type BankslipResponse struct {
	CodeBarras     string          `json:"codeBarras"`
	Digitable      string          `json:"digitable"`
	BinaryPDF      string          `json:"binaryPDF"`
	Authentication int32           `json:"authentication"`
	Receipt        BankslipReceipt `json:"receipt"`
	SettleDate     string          `json:"settleDate"`
	CreateDate     string          `json:"createDate"`
	TransactionID  int32           `json:"transactionId"`
	Urlreceipt     string          `json:"Urlreceipt"`
	ErrorCode      string          `json:"errorCode"`
	Message        string          `json:"message"`
	Status         int32           `json:"status"`
}

//BankslipPayer - Modelo do cliente
type BankslipPayer struct {
	Name         string `json:"name"`
	DocumentType string `json:"document_type"`
	Document     string `json:"document"`
	Address      string `json:"address"`
	District     string `json:"district"`
	City         string `json:"city"`
	UF           string `json:"uf"`
	ZipCode      string `json:"zipCode"`
	Email        string `json:"email"`
	Ddd          string `json:"ddd"`
	PhoneNumber  string `json:"phoneNumber"`
}

//BankslipValues - valor original para criar boleto
type BankslipValues struct {
	OriginalValue float32 `json:"originalValue"`
}

//BankslipInstructionsRequest - Modelo para instruições do boleto
type BankslipInstructionsRequest struct {
	Instruction1 string `json:"instruction1"`
	Instruction2 string `json:"instruction2"`
	Instruction3 string `json:"instruction3"`
	Instruction4 string `json:"instruction4"`
	Instruction5 string `json:"instruction5"`
}

//TransactionResponse - corpo da resposta da transação
type TransactionResponse struct {
	Transaction *Transaction `json:"transaction"`
	ErrorCode   string       `json:"errorCode"`
	Message     string       `json:"message"`
	Status      int          `json:"status"`
}

//Transaction - strutura da transação
type Transaction struct {
	Authentication int32  `json:"authentication"`
	ErrorCode      string `json:"errorCode"`
	CreateDate     string `json:"createDate"`
	Message        string `json:"message"`
	ExternalNSU    string `json:"externalNSU"`
	TransactionID  int32  `json:"transactionId"`
	Status         int    `json:"status"`
}

//Bankslip - Criar boleto
func (celcoin *CelcoinClient) Bankslip(req BankslipRequest) (*BankslipResponse, *Error, error) {
	data, _ := json.Marshal(req)
	var response *BankslipResponse
	err, errAPI := celcoin.Request("POST", "transactions/bankslip", data, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}

//GetTransactionStatus - status da transação
func (celcoin *CelcoinClient) GetTransactionStatus(transactionID string) (*TransactionResponse, *Error, error) {
	var response *TransactionResponse
	err, errAPI := celcoin.Request("GET", fmt.Sprintf("transactions/status-consult?transactionId=%s", transactionID), nil, &response)
	if err != nil {
		return nil, nil, err
	}
	if errAPI != nil {
		return nil, errAPI, nil
	}
	return response, nil, nil
}
