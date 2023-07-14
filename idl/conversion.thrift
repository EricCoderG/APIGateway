namespace go conversion

struct ConversionRequest {
    1: required string inputString;
}

struct ConversionResponse {
    1: required string convertedString;
}

service ConversionService {
    ConversionResponse convertCase(1: ConversionRequest req) (api.post="/convert");
}