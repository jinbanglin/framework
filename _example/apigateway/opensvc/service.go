package opensvc

import (
	"github.com/jinbanglin/moss/_example"
	"github.com/jinbanglin/moss/kernel"
	"github.com/jinbanglin/moss/kernel/payload"

	"context"
	"fmt"
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
	fmt.Println("-----2----", res, err)
	return res, err
}
