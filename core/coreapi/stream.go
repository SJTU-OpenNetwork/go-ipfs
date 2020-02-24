package coreapi

import (
    "context"
    peer "github.com/libp2p/go-libp2p-core/peer"
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

func (api *StreamAPI)StartWorker(ctx context.Context, conf *stream.StreamConfig, pid peer.ID) error {
	return api.stream.StartWorker(ctx, conf, pid)
}
