package coreapi

import (
	"github.com/SJTU-OpenNetwork/go-stream/stream"
)

type StreamAPI CoreAPI

func (api *SwarmAPI)StartStream(s stream.Stream) error {
	return api.stream.StartStream(s)
}

func (api *SwarmAPI)AddStreamBlock(b stream.StreamBlock) error {
	return api.stream.AddStreamBlock(b)
}

func (api *SwarmAPI)SubscribeStream(conf stream.StreamConfig) error {
	return api.stream.SubscribeStream(conf)
}