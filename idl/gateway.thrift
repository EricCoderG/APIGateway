namespace go api

include "length_api.thrift"
include "reverse_api.thrift"
include "conversion_api.thrift"
include "substring_api.thrift"

namespace go gateway.api



service GatewayService {
    length_api.LengthResponse calculateLength(1: length_api.LengthRequest req) (api.post="/length");
    reverse_api.ReverseResponse reverseString(1: reverse_api.ReverseRequest req) (api.post="/reverse");
    conversion_api.CaseConversionResponse convertCase(1: conversion_api.CaseConversionRequest req) (api.post="/convert");
    substring_api.SubstringResponse findSubstring(1: substring_api.SubstringRequest req) (api.post="/substring");
}
