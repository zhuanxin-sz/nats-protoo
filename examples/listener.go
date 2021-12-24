package main

import (
	"encoding/json"

	nprotoo "github.com/zhuanxin-sz/nats-protoo"
	"github.com/zhuanxin-sz/nats-protoo/logger"
)

func JsonEncode(str string) map[string]interface{} {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		panic(err)
	}
	return data
}

func main() {
	logger.Init("debug")
	npc := nprotoo.NewNatsProtoo(nprotoo.DefaultNatsURL)
	npc.OnRequest("channel1", func(request map[string]interface{}, accept nprotoo.AcceptFunc, reject nprotoo.RejectFunc) {
		method := request["method"].(string)
		data := request["data"].(map[string]interface{})
		logger.Infof("method => %s, data => %v", method, data)

		//accept(JsonEncode(`{}`))
		reject(404, "Not found")
	})

	npc.OnBroadcast("even1", func(data map[string]interface{}, subj string) {
		logger.Infof("Got Broadcast1 subj => %s, data => %v", subj, data)
	})

	npc.OnBroadcast("even1", func(data map[string]interface{}, subj string) {
		logger.Infof("Got Broadcast2 subj => %s, data => %v", subj, data)
	})

	npc.OnBroadcast("even1", func(data map[string]interface{}, subj string) {
		logger.Infof("Got Broadcast3 subj => %s, data => %v", subj, data)
	})

	npc.OnBroadcast("even1", func(data map[string]interface{}, subj string) {
		logger.Infof("Got Broadcast4 subj => %s, data => %v", subj, data)
	})
	select {}
}
