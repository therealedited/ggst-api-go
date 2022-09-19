package pkg

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/therealedited/ggst-api-go/internal"
	"github.com/vmihailenco/msgpack"
	"gopkg.in/ini.v1"
)

// Thanks https://github.com/optix2000/totsugeki

var SERVER_URL string = "https://ggst-game.guiltygear.com"

var UserSteamID string = ""

func init() {
	cfg, err := ini.Load("./configs/api.ini")
	if err != nil {
		panic(err)
	}
	UserSteamID = cfg.Section("api").Key("steamID").String()
}

func Login() {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf).StructAsArray(true).UseCompactEncoding(true)
	UserSteamID_int64, err := strconv.ParseInt(UserSteamID, 10, 64)
	if err != nil {
		panic(err)
	}
	UserSteamID_hex := strconv.FormatInt(UserSteamID_int64, 16)
	payload := &internal.POSTStriveLoginPacket{
		Header: internal.POSTRequestHeader{
			PlayerID: "211027113123008384", //https://github.com/xynxynxyn/ggst-api-rs/blob/master/src/requests.rs#L700
			Token:    "61a5ed4f461c2",      //https://github.com/xynxynxyn/ggst-api-rs/blob/master/src/requests.rs#L701
			Unknown3: 6,
			Version:  "0.1.5",
			Unknown4: 3,
		},
		Payload: internal.POSTLoginPayload{
			Unknown1:    1,
			SteamID:     UserSteamID,
			SteamID_hex: string(UserSteamID_hex),
			Unknown2:    256,
			Unknown3:    "",
		},
	}
	err = enc.Encode(payload)
	if err != nil {
		panic(err)
	}
	//fmt.Print(hex.EncodeToString(buf.Bytes()))
	packet := append([]byte("data="), buf.Bytes()...)
	resp := Post("/api/user/login", packet)
	fmt.Printf("%s", resp)
}

func Post(endpoint string, message_pack []byte) []byte {
	resp, err := http.Post(SERVER_URL+endpoint, "application/x-www-form-urlencoded", bytes.NewBuffer(message_pack))
	if err != nil {
		panic(err)
	}
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return buf
}
