package main

import (
	"encoding/json"
	"fmt"
)

type model struct {
	Code    int
	Descrip string
}

func main() {
	var data []model
	rcvd := `[{"Code":200,"Descrip":"StatusOk"},
	{"Code":301,"Descrip":"StatusMovedPermanently"},
	{"Code":302,"Descrip":"StatusFound"},
	{"Code":303,"Descrip":"StatusSeeOther"},
	{"Code":307,"Descrip":"StatusTemporaryRedirect"},
	{"Code":400,"Descrip":"StatusBadRequest"},
	{"Code":401,"Descrip":"StatusUnauthorized"},
	{"Code":402,"Descrip":"StatusPaymentRequired"},
	{"Code":403,"Descrip":"StatusForbidden"},
	{"Code":404,"Descrip":"StatusNotFound"},
	{"Code":405,"Descrip":"StatusMethodNotAllowed"},
	{"Code":418,"Descrip":"StatusTeapot"},
	{"Code":500,"Descrip":"StatusInternalServerError"}]`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		fmt.Println(err)
	}
	for _, r := range data {
		fmt.Println(r.Code, "--", r.Descrip)
	}

}
