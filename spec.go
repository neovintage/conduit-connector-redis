package redis

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func Specification() sdk.Specification {
	return sdk.Specification{
		Name:    "redis",
		Summary: "A Redis PubSub source connector for Conduit, written in Go.",
		Version: "v0.1.0",
		Author:  "Rimas Silkaitis",
		SourceParams: map[string]sdk.Parameter{
			"uri": {
				Default:     "",
				Description: "URI for the Redis instance",
				Required:    true,
			},
            "channel": {
                Default:     "",
                Description: "The name of the channel for the subscription",
                Required:    true,
            },
		},
	}
}
