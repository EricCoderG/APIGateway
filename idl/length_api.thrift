namespace go length.api

struct LengthRequest {
    1: required string inputString;
}

struct LengthResponse {
    1: required i32 length;
    2: required string msg;
}

service LengthService {
    LengthResponse calculateLength(1: LengthRequest req) (api.post="/length");
}
