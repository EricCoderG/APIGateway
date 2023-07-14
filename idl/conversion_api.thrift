namespace go conversion.api

struct CaseConversionRequest {
    1: required string inputString;
}

struct CaseConversionResponse {
    1: required string convertedString;
    2: required string msg;
}

service CaseConversionService {
    CaseConversionResponse convertCase(1: CaseConversionRequest req) (api.post="/convert");
}