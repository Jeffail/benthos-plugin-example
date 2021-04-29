package processor

import (
	"context"

	"github.com/Jeffail/benthos/v3/public/x/service"
)

func init() {
	constructor := func(conf interface{}, mgr *service.Resources) (service.Processor, error) {
		return &reverseProcessor{}, nil
	}

	err := service.RegisterProcessor("reverse", service.NewConfigSpec(), constructor)
	if err != nil {
		panic(err)
	}
}

//------------------------------------------------------------------------------

type reverseProcessor struct{}

func (r *reverseProcessor) Process(ctx context.Context, m *service.Message) ([]*service.Message, error) {
	bytesContent, err := m.AsBytes()
	if err != nil {
		return nil, err
	}

	newBytes := make([]byte, len(bytesContent))
	for i, b := range bytesContent {
		newBytes[len(newBytes)-i-1] = b
	}

	m.SetBytes(newBytes)
	return []*service.Message{m}, nil
}

func (r *reverseProcessor) Close(ctx context.Context) error {
	return nil
}
