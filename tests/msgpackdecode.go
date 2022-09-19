package tests

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/therealedited/ggst-api-go/internal"
	"github.com/therealedited/ggst-api-go/pkg"
	"github.com/vmihailenco/msgpack/v5"
)

func Test1() {
	api_sys_getenv := "9295a0a002a5302e312e350391cd0100"
	data, err := hex.DecodeString(strings.TrimRight(api_sys_getenv, "\x00"))
	if err != nil {
		panic(err)
	}
	packet := &internal.StriveInitPacket{}
	err2 := msgpack.Unmarshal(data, packet)
	if err2 != nil {
		panic(err2)
	}

	fmt.Printf("%v", packet)
}

func Test2() {
	resp := pkg.Post("/api/sys/get_env", []byte("data=9295a0a002a5302e312e350391cd0100"))
	parsedResponse := &internal.EnvAnswer{
		Hash: string(resp[3:16]),
		Date: string(resp[18:37]),
		Ver1: string(resp[38:43]),
		Ver2: string(resp[44:49]),
		Ver3: string(resp[50:55]),
		Link: string(resp[61:]),
	}
	fmt.Printf("%v", parsedResponse)
}
