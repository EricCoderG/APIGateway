package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_substring_normal(t *testing.T) {
	url := "http://127.0.0.1:8088/substring"
	jsonStr := `{"mainString" : "123456","subString" : "123"}`
	expectedResp := `{"positions":[0]}`
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		t.Fatal(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = resp.Body.Close()

	if err != nil {
		return
	}

	respString := string(respBytes)

	if respString != expectedResp {
		t.Error("expected resp is " + expectedResp + " but real resp is " + respString + "!")
	} else {
		t.Log("Normal Test Pass!")
	}
}

func Test_substring_zero(t *testing.T) {
	url := "http://127.0.0.1:8088/substring"
	jsonStr := `{"mainString" : "","subString" : ""}`
	expectedResp := `{"positions":[]}`
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		t.Fatal(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = resp.Body.Close()

	if err != nil {
		return
	}

	respString := string(respBytes)

	if respString != expectedResp {
		t.Error("expected resp is " + expectedResp + " but real resp is " + respString + "!")
	} else {
		t.Log("Zero Test Pass!")
	}
}

func Test_substring_multiAnswer(t *testing.T) {
	url := "http://127.0.0.1:8088/substring"
	jsonStr := `{"mainString" : "123456123","subString" : "123"}`
	expectedResp := `{"positions":[0,6]}`
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		t.Fatal(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = resp.Body.Close()

	if err != nil {
		return
	}

	respString := string(respBytes)

	if respString != expectedResp {
		t.Error("expected resp is " + expectedResp + " but real resp is " + respString + "!")
	} else {
		t.Log("MultiAnswer Test Pass!")
	}
}
