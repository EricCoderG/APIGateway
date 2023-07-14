// Code generated by hertz generator.

package gateway

import (
	"api-gateway/hertz-http-server/biz/model/conversion"
	"api-gateway/hertz-http-server/biz/model/length"
	"api-gateway/hertz-http-server/biz/model/reverse"
	"api-gateway/hertz-http-server/biz/model/substring/api"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CalculateLength .
// @router /length [POST]
func CalculateLength(ctx context.Context, c *app.RequestContext) {
	var err error
	var req length.LengthRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(length.LengthResponse)

	c.JSON(consts.StatusOK, resp)
}

// ReverseString .
// @router /reverse [POST]
func ReverseString(ctx context.Context, c *app.RequestContext) {
	var err error
	var req reverse.ReverseRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(reverse.ReverseResponse)

	c.JSON(consts.StatusOK, resp)
}

// ConvertCase .
// @router /convert [POST]
func ConvertCase(ctx context.Context, c *app.RequestContext) {
	var err error
	var req conversion.ConversionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(conversion.ConversionResponse)

	c.JSON(consts.StatusOK, resp)
}

// FindSubstring .
// @router /substring [POST]
func FindSubstring(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.SubstringRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(api.SubstringResponse)

	c.JSON(consts.StatusOK, resp)
}
