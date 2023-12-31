namespace go substring

struct SubstringRequest {
    1: required string mainString;
    2: required string subString;
}

struct SubstringResponse {
    1: required list<i32> positions;
}

service SubstringService {
    SubstringResponse findSubstring(1: SubstringRequest req) (api.post="/substring");
}