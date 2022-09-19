package internal

//Thanks https://github.com/optix2000/totsugeki

type RequestHeader struct {
	_msgpack struct{} `msgpack:",as_array"`
	PlayerID string
	Token    string
	Unknown3 int
	Version  string
	Unknown4 int
}

type GameData struct {
	_msgpack struct{} `msgpack:",as_array"`
	Unknown1 int
	Unknown2 int
	Unknown3 int
	JsonData map[string]interface{}
}

type EnvPayload struct {
	_msgpack struct{} `msgpack:",as_array"`
	Unknown4 int      // 256
}

type EnvAnswer struct {
	Hash string
	Date string
	Ver1 string
	Ver2 string
	Ver3 string
	Link string
}

type LoginPayload struct {
}

type StriveInitPacket struct {
	_msgpack struct{} `msgpack:",as_array"`
	Header   RequestHeader
	Payload  EnvPayload
}

type StriveLoginPacket struct {
	_msgpack struct{} `msgpack:",as_array"`
	Header   RequestHeader
	Payload  LoginPayload
}
