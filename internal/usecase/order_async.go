package usecase

import (
	"context"
	"log"

	"github.com/Helltale/tz-telecom/internal/domain"
)

type OrderJob struct {
	UserID int64
	Items  []domain.OrderItem
}

type OrderWorker struct {
	queue chan OrderJob
	uc    *OrderUseCase
}

func NewOrderWorker(uc *OrderUseCase, size int) *OrderWorker {
	w := &OrderWorker{
		queue: make(chan OrderJob, size),
		uc:    uc,
	}
	go w.run()
	return w
}

func (w *OrderWorker) run() {
	for job := range w.queue {
		if err := w.uc.CreateOrder(context.Background(), job.UserID, job.Items); err != nil {
			log.Printf("async order failed: %v", err)
		}
	}
}

func (w *OrderWorker) Enqueue(job OrderJob) {
	w.queue <- job
}
