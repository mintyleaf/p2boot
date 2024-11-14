package main

import (
	"context"

	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
)

func NewKDHT(ctx context.Context, host host.Host) (*dht.IpfsDHT, error) {
	var options []dht.Option

	options = append(options, dht.Mode(dht.ModeAutoServer))

	kdht, err := dht.New(ctx, host, options...)
	if err != nil {
		return nil, err
	}

	if err = kdht.Bootstrap(ctx); err != nil {
		return nil, err
	}

	return kdht, nil
}
