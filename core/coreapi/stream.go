package coreapi

import (
    "context"

	"github.com/SJTU-OpenNetwork/go-stream"
)

type StreamAPI CoreAPI

func (api *StreamAPI)StartStream(ctx context.Context, s *stream.Stream) error {
	return api.stream.StartStream(ctx, s)
}

func (api *StreamAPI)AddStreamBlock(ctx context.Context, b *stream.StreamBlock) error {
	return api.stream.AddStreamBlock(ctx, b)
}

func (api *StreamAPI)SubscribeStream(ctx context.Context, conf *stream.StreamConfig) error {
	return api.stream.SubscribeStream(ctx, conf)
}
