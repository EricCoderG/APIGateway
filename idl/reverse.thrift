namespace go reverse

struct ReverseRequest {
    1: required string inputString;
}

struct ReverseResponse {
    1: required string reversedString;
}

service ReverseService {
    ReverseResponse reverseString(1: ReverseRequest req) (api.post="/reverse");
}
