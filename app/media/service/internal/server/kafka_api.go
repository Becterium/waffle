package server

import (
	"context"
	"fmt"
	"waffle/model/mq_kafka"

	"github.com/tx7do/kratos-transport/broker"
)

func imageCreator() broker.Any  { return &mq_kafka.Image{} }
func avatarCreator() broker.Any { return &mq_kafka.Avatar{} }

type imageHandler func(_ context.Context, topic string, headers broker.Headers, msg *mq_kafka.Image) error
type avatarHandler func(_ context.Context, topic string, headers broker.Headers, msg *mq_kafka.Avatar) error

func registerImageHandler(fnc imageHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		switch t := event.Message().Body.(type) {
		case *mq_kafka.Image:
			if err := fnc(ctx, event.Topic(), event.Message().Headers, t); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}
		return nil
	}
}

func registerAvatarHandler(fnc avatarHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		switch t := event.Message().Body.(type) {
		case *mq_kafka.Avatar:
			if err := fnc(ctx, event.Topic(), event.Message().Headers, t); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}
		return nil
	}
}
