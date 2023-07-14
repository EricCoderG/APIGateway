namespace go substring.api

struct SubstringRequest {
    1: required string mainString;
    2: required string subString;
}

struct SubstringResponse {
    1: required list<i32> positions;
    2: required string msg;
}

service SubstringService {
    SubstringResponse findSubstring(1: SubstringRequest req) (api.post="/substring");
}