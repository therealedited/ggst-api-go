package tests

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/therealedited/ggst-api-go/internal"
	"github.com/therealedited/ggst-api-go/pkg"
	"github.com/vmihailenco/msgpack"
)

func Test1() {
	api_sys_getenv := "9295a0a002a5302e312e350391cd0100"
	data, err := hex.DecodeString(strings.TrimRight(api_sys_getenv, "\x00"))
	if err != nil {
		panic(err)
	}
	packet := &internal.POSTStriveInitPacket{}
	err2 := msgpack.Unmarshal(data, packet)
	if err2 != nil {
		panic(err2)
	}

	fmt.Printf("%v", packet)
}

func Test3() {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf).StructAsArray(true).UseCompactEncoding(true)
	payload := &internal.POSTStriveReplayPacket{
		Header: internal.POSTRequestHeader{
			PlayerID: "211027113123008384", //https://github.com/xynxynxyn/ggst-api-rs/blob/master/src/requests.rs#L700
			Token:    "61a5ed4f461c2",      //https://github.com/xynxynxyn/ggst-api-rs/blob/master/src/requests.rs#L701
			Unknown3: 2,
			Version:  "0.1.5",
			Unknown4: 3,
		},
		Payload: internal.POSTReplayPayload{
			Unknown1:      1,
			Index:         0,
			ReplayPerPage: 5,
			Query: internal.ReplayQuery{
				Unknown1:           -1,
				PlayerSearch:       0, //All
				MinFloor:           1,
				MaxFloor:           99,
				Seq:                []int{},
				Char1:              -1,
				Char2:              -1,
				Winner:             0,
				PrioritizeBestBout: 0,
				Unknown2:           1,
			},
		},
	}
	err := enc.Encode(payload)
	if err != nil {
		panic(err)
	}
	fmt.Print(hex.EncodeToString(buf.Bytes()))
	packet := append([]byte("data="), buf.Bytes()...)
	resp := pkg.Post("/api/user/login", packet)
	fmt.Printf("%s", resp)
}
