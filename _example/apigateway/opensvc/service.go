package opensvc

import (
	"github.com/jinbanglin/moss/_example"
	"github.com/jinbanglin/moss/kernel"
	"github.com/jinbanglin/moss/kernel/payload"

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
	res, err := kernel.WatcherInvoking(_example.Addgrpcsvc, ctx, request)
	return res, err
}
