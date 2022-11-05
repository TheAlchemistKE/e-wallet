package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func TopUpAccount(phoneNumber int, amount float64) []byte {
	url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	method := "POST"
	requestBody, _ := json.Marshal(map[string]any{
		"BusinessShortCode": 174379,
		"Password":          "MTc0Mzc5YmZiMjc5ZjlhYTliZGJjZjE1OGU5N2RkNzFhNDY3Y2QyZTBjODkzMDU5YjEwZjc4ZTZiNzJhZGExZWQyYzkxOTIwMjIxMTAyMjEwNzU1",
		"Timestamp":         time.Now().Unix(),
		"TransactionType":   "CustomerPayBillOnline",
		"Amount":            amount,
		"PartyA":            phoneNumber,
		"PartyB":            174379,
		"PhoneNumber":       phoneNumber,
		"CallBackURL":       "https://mydomain.com/path",
		"AccountReference":  "LetaWallet",
		"TransactionDesc":   "Payment of X",
	})

	payload := bytes.NewBuffer(requestBody)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer lAwV6ODiTOLjjHYcIke9p23GTZq3")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	return body
}
