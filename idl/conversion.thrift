namespace go conversion

struct CaseConversionRequest {
    1: required string inputString;
}

struct CaseConversionResponse {
    1: required string convertedString;
}

service CaseConversionService {
    CaseConversionResponse convertCase(1: CaseConversionRequest req) (api.post="/convert");
}