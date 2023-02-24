package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error couldn't read body: %v \n", err)
		return
	}
	e := json.Unmarshal(body, x)
	if e != nil {
		fmt.Printf("Error couldn't convert body: %v \n", e)
		return
	}
	return
}
