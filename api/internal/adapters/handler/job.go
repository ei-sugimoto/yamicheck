package handler

import (
	"context"

	"connectrpc.com/connect"
	jobv1 "github.com/ei-sugimoto/yamicheck/api/gen/job/v1"
	"github.com/ei-sugimoto/yamicheck/api/internal/domain"
	"github.com/ei-sugimoto/yamicheck/api/internal/ports/inner"
)

type JobHandler struct {
	JobService inner.JobService
}

func NewJobHandler(uc inner.JobService) *JobHandler {
	return &JobHandler{
		JobService: uc,
	}
}

func (jh *JobHandler) Check(ctx context.Context, req *connect.Request[jobv1.CheckRequest]) (*connect.Response[jobv1.CheckResponse], error) {
	newJob := domain.Job{
		CompanyName: req.Msg.CompanyName,
		HourlyWage:  int(req.Msg.HourlyWage),
		Description: req.Msg.Description,
	}
	res, err := jh.JobService.Check(&newJob)
	if err != nil {
		return nil, err
	}
	msg, err := res.String()
	if err != nil {
		return nil, err
	}
	return &connect.Response[jobv1.CheckResponse]{Msg: &jobv1.CheckResponse{Level: int64(res.Integer()), Message: msg}}, nil
}
