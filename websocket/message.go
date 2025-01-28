package websocket

import (
	"salimon/nexus/types"
	"time"
)

func handleMessage(body string, ctx *types.WsContext) {
	if body == "" {
		sendInvalidPayload(ctx.Conn)
		return
	}
	if ctx.Entity == "" {
		sendErrorPayload("no entity connected", ctx.Conn)
		return
	}

	tokens := []string{"this", "is", "a", "new", "message"}

	for i := 0; i < len(tokens); i++ {
		payload := types.WsResponse{
			Action: "TOKEN",
			Token:  tokens[i],
		}
		sendPayload(payload, ctx.Conn)
		time.Sleep(2 * time.Second)
	}
}
