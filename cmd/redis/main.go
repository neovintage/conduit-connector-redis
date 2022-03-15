package main

import (
	redis "github.com/conduitio/conduit-connector-redis"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func main() {
	sdk.Serve(redis.Specification, redis.NewSource, nil)
}
