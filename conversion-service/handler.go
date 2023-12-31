package main

import (
	conversion "api-gateway/conversion-service/kitex_gen/conversion"
	"context"
	"strings"
)

// ConversionServiceImpl implements the last service interface defined in the IDL.
type ConversionServiceImpl struct{}

// ConvertCase implements the ConversionServiceImpl interface.
func (s *ConversionServiceImpl) ConvertCase(ctx context.Context, req *conversion.ConversionRequest) (resp *conversion.ConversionResponse, err error) {
	resp = &conversion.ConversionResponse{
		ConvertedString: strings.ToUpper(req.InputString),
	}
	return
}
