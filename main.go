package main

import (
	"context"
	"log"

	sdk "github.com/apito-io/go-apito-plugin-sdk"
)

func main() {
	log.Printf("🎯 [hc-ai-plugin] Starting AI plugin...")
	plugin := sdk.Init("hc-ai-plugin", "1.0.0", "apito-plugin-key")
	statusType := sdk.NewObjectType("AIStatus", "AI plugin status").
		AddStringField("status", "Plugin status", false).
		AddStringField("version", "Plugin version", false).
		Build()
	plugin.RegisterQuery("getAIStatus",
		sdk.ComplexObjectField("Get AI plugin status", statusType),
		func(ctx context.Context, rawArgs map[string]interface{}) (interface{}, error) {
			return map[string]interface{}{"status": "ready", "version": "1.0.0"}, nil
		})
	log.Printf("🚀 [hc-ai-plugin] Plugin ready")
	plugin.Serve()
}
