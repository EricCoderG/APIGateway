package test

import (
	substring "api-gateway/substring-service/kitex_gen/substring"
	handler "api-gateway/unit-test/utils"
	"context"
	"fmt"
	"testing"
)

func Test_FindSubstring_func_normal(t *testing.T) {
	mainStr := "abcd"
	subStr := "bc"
	expectedResp := "[1]"
	ssi := new(handler.SubstringServiceImpl)
	req := &substring.SubstringRequest{
		MainString: mainStr,
		SubString:  subStr,
	}
	resp, err := ssi.FindSubstring(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	respstr := fmt.Sprintf("%v", resp.Positions)
	if respstr != expectedResp {
		t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, respstr)
	}
}

func Test_FindSubstring_func_zero(t *testing.T) {
	mainStr := "abcd"
	subStr := "ef"
	expectedResp := "[]"
	ssi := new(handler.SubstringServiceImpl)
	req := &substring.SubstringRequest{
		MainString: mainStr,
		SubString:  subStr,
	}
	resp, err := ssi.FindSubstring(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	respstr := fmt.Sprintf("%v", resp.Positions)
	if respstr != expectedResp {
		t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, respstr)
	}
}

func Test_FindSubstring_func_MultiAnswer(t *testing.T) {
	mainStr := "123456123"
	subStr := "123"
	expectedResp := "[0 6]"
	ssi := new(handler.SubstringServiceImpl)
	req := &substring.SubstringRequest{
		MainString: mainStr,
		SubString:  subStr,
	}
	resp, err := ssi.FindSubstring(context.Background(), req)
	if err != nil {
		t.Error(err)
	}
	respstr := fmt.Sprintf("%v", resp.Positions)
	if respstr != expectedResp {
		t.Errorf("expected resp is %s but resp in fact is %s", expectedResp, respstr)
	}
}
