package internal

//Thanks https://github.com/optix2000/totsugeki

type POSTRequestHeader struct {
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

type POSTEnvPayload struct {
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

type POSTLoginPayload struct {
	_msgpack    struct{} `msgpack:",as_array"`
	Unknown1    int
	SteamID     string
	SteamID_hex string
	Unknown2    int
	Unknown3    string
}

type ReplayQuery struct {
	Unknown1           int
	PlayerSearch       int
	MinFloor           int
	MaxFloor           int
	Seq                []int
	Char1              int
	Char2              int
	Winner             int
	PrioritizeBestBout int
	Unknown2           int
}

type POSTReplayPayload struct {
	_msgpack      struct{} `msgpack:",as_array"`
	Unknown1      int
	Index         int
	ReplayPerPage int
	Query         ReplayQuery
}

type POSTStriveInitPacket struct {
	_msgpack struct{} `msgpack:",as_array"`
	Header   POSTRequestHeader
	Payload  POSTEnvPayload
}

type POSTStriveLoginPacket struct {
	Header  POSTRequestHeader
	Payload POSTLoginPayload
}

type POSTStriveReplayPacket struct {
	Header  POSTRequestHeader
	Payload POSTReplayPayload
}
