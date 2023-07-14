namespace go reverse.api

struct ReverseRequest {
    1: required string inputString;
}

struct ReverseResponse {
    1: required string reversedString;
    2: required string msg;
}

service ReverseService {
    ReverseResponse reverseString(1: ReverseRequest req) (api.post="/reverse");
}
