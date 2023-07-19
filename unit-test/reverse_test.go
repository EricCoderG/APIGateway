package test

import (
	reverse "api-gateway/reverse-service/kitex_gen/reverse"
	handler "api-gateway/unit-test/utils"
	"context"
	"testing"
)

func Test_reverseString_func_normal(t *testing.T) {
	testString := "abcd"
	expectedResp := "dcba"
	csi := new(handler.ReverseServiceImpl)
	req := &reverse.ReverseRequest{
		InputString: testString,
	}
	resp, err := csi.ReverseString(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if resp.ReversedString != expectedResp {
		t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, resp.ReversedString)
	}
}

func Test_reverseString_func_zero(t *testing.T) {
	testString := ""
	expectedResp := ""
	csi := new(handler.ReverseServiceImpl)
	req := &reverse.ReverseRequest{
		InputString: testString,
	}
	resp, err := csi.ReverseString(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if resp.ReversedString != expectedResp {
		t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, resp.ReversedString)
	}
}

func Test_reverseString_func_Mixed(t *testing.T) {
	testString := "abcd123ABC"
	expectedResp := "CBA321dcba"
	csi := new(handler.ReverseServiceImpl)
	req := &reverse.ReverseRequest{
		InputString: testString,
	}
	resp, err := csi.ReverseString(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	if resp.ReversedString != expectedResp {
		t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, resp.ReversedString)
	}
}
