// Code generated by Kitex v0.6.1. DO NOT EDIT.

package substringservice

import (
	substring "api-gateway/substring-service/kitex_gen/substring"
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return substringServiceServiceInfo
}

var substringServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "SubstringService"
	handlerType := (*substring.SubstringService)(nil)
	methods := map[string]kitex.MethodInfo{
		"findSubstring": kitex.NewMethodInfo(findSubstringHandler, newSubstringServiceFindSubstringArgs, newSubstringServiceFindSubstringResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "substring",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func findSubstringHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*substring.SubstringServiceFindSubstringArgs)
	realResult := result.(*substring.SubstringServiceFindSubstringResult)
	success, err := handler.(substring.SubstringService).FindSubstring(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSubstringServiceFindSubstringArgs() interface{} {
	return substring.NewSubstringServiceFindSubstringArgs()
}

func newSubstringServiceFindSubstringResult() interface{} {
	return substring.NewSubstringServiceFindSubstringResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FindSubstring(ctx context.Context, req *substring.SubstringRequest) (r *substring.SubstringResponse, err error) {
	var _args substring.SubstringServiceFindSubstringArgs
	_args.Req = req
	var _result substring.SubstringServiceFindSubstringResult
	if err = p.c.Call(ctx, "findSubstring", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
