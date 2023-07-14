package main

import (
	api "api-gateway/substring-service/kitex_gen/substring/api"
	"context"
)

// SubstringServiceImpl implements the last service interface defined in the IDL.
type SubstringServiceImpl struct{}

// FindSubstring implements the SubstringServiceImpl interface.
func (s *SubstringServiceImpl) FindSubstring(ctx context.Context, req *api.SubstringRequest) (resp *api.SubstringResponse, err error) {
	// TODO: Your code here...
	return
}
