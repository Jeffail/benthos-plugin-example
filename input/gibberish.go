package input

import (
	"context"
	"math/rand"

	"github.com/Jeffail/benthos/v3/public/x/service"
)

func init() {
	type gibberishConfig struct {
		Length int `yaml:"length"`
	}

	configSpec := service.NewConfigSpec().WithConstructor(func() interface{} {
		return &gibberishConfig{
			Length: 100,
		}
	})

	constructor := func(conf interface{}, mgr *service.Resources) (service.Input, error) {
		gconf := conf.(*gibberishConfig)
		return &gibberishInput{gconf.Length}, nil
	}

	err := service.RegisterInput("gibberish", configSpec, constructor)
	if err != nil {
		panic(err)
	}
}

//------------------------------------------------------------------------------

type gibberishInput struct {
	length int
}

func (g *gibberishInput) Connect(ctx context.Context) error {
	return nil
}

func (g *gibberishInput) Read(ctx context.Context) (*service.Message, service.AckFunc, error) {
	b := make([]byte, g.length)
	for k := range b {
		b[k] = byte((rand.Int() % 94) + 32)
	}
	return service.NewMessage(b), service.NoopAckFunc, nil
}

func (g *gibberishInput) Close(ctx context.Context) error {
	return nil
}
