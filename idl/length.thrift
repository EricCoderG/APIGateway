namespace go length

struct LengthRequest {
    1: required string inputString;
}

struct LengthResponse {
    1: required i32 length;
}

service LengthService {
    LengthResponse calculateLength(1: LengthRequest req) (api.post="/length");
}
