package adt

import (
	"context"

	cid "github.com/ipfs/go-cid"
)

// Store defines an interface required to back the ADTs in this package.
type Store interface {
	Context() context.Context
	IpldStore
}

type IpldStore interface {
	Get(ctx context.Context, c cid.Cid, out interface{}) error
	Put(ctx context.Context, v interface{}) (cid.Cid, error)
}
