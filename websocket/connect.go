package websocket

import (
	"salimon/nexus/types"
)

func handleConnect(entity string, ctx *types.WsContext) {
	if entity == "" {
		sendInvalidPayload(ctx.Conn)
		return
	}
	ctx.Entity = entity

	result := true
	payload := types.WsResponse{
		Action: "CONNECT",
		Result: &result,
	}
	sendPayload(payload, ctx.Conn)
}
