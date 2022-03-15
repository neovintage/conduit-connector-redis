package redis

import (
	"context"
    "fmt"
    "time"

	sdk "github.com/conduitio/conduit-connector-sdk"
    "github.com/go-redis/redis/v8"
)

// Source connector
type Source struct {
	sdk.UnimplementedSource

	config map[string]string
    rdbOptions *redis.Options
    rdb *redis.Client
    pubsub *redis.PubSub
}

func NewSource() sdk.Source {
	return &Source{}
}

func (s *Source) Configure(ctx context.Context, m map[string]string) error {
	uri, ok := m[ConfigURI]
	if !ok {
		return requiredConfigErr(ConfigURI)
	}
	s.config = m

    options, uriErr := redis.ParseURL(uri)
    if uriErr != nil {
        return uriErr
    }
    s.rdbOptions = options
	return nil
}

func (s *Source) Open(ctx context.Context, position sdk.Position) error {
    s.rdb = redis.NewClient(s.rdbOptions)
    s.pubsub = s.rdb.Subscribe(ctx, s.config[ConfigChannel])

    // Wait for confirmation that subscription is created before recieving
    // any messages
    _, err := s.pubsub.Receive(ctx)
    if err != nil {
        return err
    }

	return nil
}

func (s *Source) Read(ctx context.Context) (sdk.Record, error) {
    select {
        case msg := <-s.pubsub.Channel():
            fmt.Printf("Record: %s", msg.Payload)
            return sdk.Record{
                Position:   nil,
                CreatedAt:  time.Now(),
                Payload:    sdk.RawData(msg.Payload),
            }, nil
        case <-ctx.Done():
            return sdk.Record{}, ctx.Err()
    }
}

func (s *Source) Ack(ctx context.Context, position sdk.Position) error {
	return nil // no ack needed
}

func (s *Source) Teardown(ctx context.Context) error {
	if s.rdb != nil {
		return s.rdb.Close()
	}
	return nil
}

