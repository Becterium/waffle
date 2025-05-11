package server

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"waffle/app/media/service/internal/biz"
)

const (
	EveryMidNight = "0 0 0 * * *"
)

func NewCronServer(useCase biz.ImageRepo) *CronWorker {
	c := cron.New(cron.WithSeconds())
	registerCornFunction(c, useCase)
	return &CronWorker{
		sche: c,
	}
}

type CronWorker struct {
	sche *cron.Cron
}

func (c *CronWorker) Start(ctx context.Context) error {
	log.Println(fmt.Sprintf("=========schedule start==========="))
	c.sche.Start()
	return nil
}

func (c *CronWorker) Stop(ctx context.Context) error {
	c.sche.Stop()
	return nil
}

func registerCornFunction(c *cron.Cron, u biz.ImageRepo) {
	c.AddFunc(EveryMidNight, u.CronSynchronizeImageViewFromRedis())
}
