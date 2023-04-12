package main

type SubstituteRewardReq struct {
	Body `json:"BODY"`
}

type Body struct {
	EncryptData string `json:"ENCRYPT_DATA"`
	EncryptKey  string `json:"ENCRYPT_KEY"`
}
