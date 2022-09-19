package pkg

import (
	"bytes"
	"io"
	"net/http"

	"gopkg.in/ini.v1"
)

// Thanks https://github.com/optix2000/totsugeki

var SERVER_URL string = "https://ggst-game.guiltygear.com"

var UID string = ""

func init() {
	cfg, err := ini.Load("./configs/config.ini")
	if err != nil {
		panic(err)
	}
	UID = cfg.Section("api").Key("UID").String()
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
