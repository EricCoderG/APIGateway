namespace go gateway

include "length.thrift"
include "reverse.thrift"
include "conversion.thrift"
include "substring.thrift"

service GatewayService {
    length.LengthResponse calculateLength(1: length.LengthRequest req) (api.post="/length");
    reverse.ReverseResponse reverseString(1: reverse.ReverseRequest req) (api.post="/reverse");
    conversion.ConversionResponse convertCase(1: conversion.ConversionRequest req) (api.post="/convert");
    substring.SubstringResponse findSubstring(1: substring.SubstringRequest req) (api.post="/substring");
}
