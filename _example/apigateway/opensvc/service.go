package opensvc

import (
	"moss/_example"
	"moss/kernel"
	"moss/kernel/payload"

	"context"
)

type Service interface {
	Register(ctx context.Context, request *payload.MossPacket) (*payload.MossPacket, error)
}

type EntryService struct{}

func NewEntryService() Service {
	return &EntryService{}
}

func (*EntryService) Register(ctx context.Context, request *payload.MossPacket) (*payload.MossPacket, error) {
	return kernel.WatcherInvoking(_example.Addgrpcsvc, ctx, request)
}
