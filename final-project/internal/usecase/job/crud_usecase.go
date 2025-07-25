package todo_list_usecase

import (
	"context"
	"fmt"
	"time"

	generalEntity "github.com/kriti-mauludin/final-project/entity"
	"github.com/kriti-mauludin/final-project/internal/helper"
	"github.com/kriti-mauludin/final-project/internal/repository/mysql"
	mentity "github.com/kriti-mauludin/final-project/internal/repository/mysql/entity"
	"github.com/kriti-mauludin/final-project/internal/usecase"
	"github.com/kriti-mauludin/final-project/internal/usecase/job/entity"
	errwrap "github.com/pkg/errors"
)

type CrudJobUsecase struct {
	jobRepo mysql.IJobRepository
}

func NewCrudJobUsecase(
	jobRepo mysql.IJobRepository,
) *CrudJobUsecase {
	return &CrudJobUsecase{jobRepo}
}

type ICrudJobUsecase interface {
	GetByUserID(ctx context.Context, userID int64) (res []*entity.JobResponse, err error)
	GetByID(ctx context.Context, jobID int64) (*entity.JobResponse, error)
	Create(ctx context.Context, jobReq entity.JobReq) (*entity.JobResponse, error)
	UpdateByID(ctx context.Context, jobReq entity.JobReq) error
	DeleteByID(ctx context.Context, jobID int64) error
}

func (t *CrudJobUsecase) GetByUserID(ctx context.Context, userID int64) (res []*entity.JobResponse, err error) {
	funcName := "CrudJobUsecase.GetByUserID"
	captureFieldError := generalEntity.CaptureFields{
		"user_id": helper.ToString(userID),
	}

	result, err := t.jobRepo.GetByUserID(ctx, userID)
	if err != nil {
		helper.LogError("jobRepo.GetByUserID", funcName, err, captureFieldError, "")

		return nil, err
	}

	for _, v := range result {
		res = append(res, &entity.JobResponse{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			CreatedAt:   helper.ConvertToJakartaTime(v.CreatedAt),
			UpdatedAt:   helper.ConvertToJakartaTime(v.UpdatedAt),
		})
	}

	return res, nil
}

func (t *CrudJobUsecase) GetByID(ctx context.Context, jobID int64) (*entity.JobResponse, error) {
	funcName := "CrudJobUsecase.GetByID"
	captureFieldError := generalEntity.CaptureFields{
		"user_id": helper.ToString(jobID),
	}

	data, err := t.jobRepo.GetByID(ctx, jobID)
	if err != nil {
		helper.LogError("jobRepo.GetByID", funcName, err, captureFieldError, "")

		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	return &entity.JobResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   helper.ConvertToJakartaTime(data.CreatedAt),
		UpdatedAt:   helper.ConvertToJakartaTime(data.UpdatedAt),
	}, nil
}

func (t *CrudJobUsecase) Create(ctx context.Context, jobReq entity.JobReq) (*entity.JobResponse, error) {
	funcName := "CrudJobUsecase.Create"
	captureFieldError := generalEntity.CaptureFields{
		"user_id": helper.ToString(jobReq.UserID),
		"payload": helper.ToString(jobReq),
	}

	if errMsg := usecase.ValidateStruct(jobReq); errMsg != "" {
		return nil, errwrap.Wrap(fmt.Errorf(generalEntity.INVALID_PAYLOAD_CODE), errMsg)
	}

	jobPayload := &mentity.Job{
		UserID:      jobReq.UserID,
		Name:        jobReq.Name,
		Description: jobReq.Description,
		CreatedAt:   time.Now(),
	}

	err := t.jobRepo.Create(ctx, nil, jobPayload, false)
	if err != nil {
		helper.LogError("jobRepo.Create", funcName, err, captureFieldError, "")

		return nil, err
	}

	return &entity.JobResponse{
		ID:          jobPayload.ID,
		Name:        jobPayload.Name,
		Description: jobPayload.Description,
		CreatedAt:   helper.ConvertToJakartaTime(jobPayload.CreatedAt),
	}, nil
}

func (t *CrudJobUsecase) UpdateByID(ctx context.Context, jobReq entity.JobReq) error {
	funcName := "CrudJobUsecase.UpdateByID"
	jobID := jobReq.ID

	captureFieldError := generalEntity.CaptureFields{
		"user_id": helper.ToString(jobReq.UserID),
		"payload": helper.ToString(jobReq),
	}

	// Start DB Transaction
	if err := mysql.DBTransaction(t.jobRepo, func(trx mysql.TrxObj) error {
		// Locking Data
		lockedData, err := t.jobRepo.LockByID(ctx, trx, jobID)
		if err != nil {
			helper.LogError("jobRepo.LockByID", funcName, err, captureFieldError, "")

			return err
		}
		if lockedData == nil {
			return fmt.Errorf("DATA IS NOT EXIST")
		}

		// Process Update
		if err := t.jobRepo.Update(ctx, trx, lockedData, &mentity.Job{
			Name:        jobReq.Name,
			Description: jobReq.Description,
			UpdatedAt:   time.Now(),
		}); err != nil {
			helper.LogError("jobRepo.Update", funcName, err, captureFieldError, "")

			return err
		}

		return nil
	}); err != nil {
		helper.LogError("roleRepo.DBTransaction", funcName, err, captureFieldError, "")

		return err
	}

	return nil
}

func (t *CrudJobUsecase) DeleteByID(ctx context.Context, jobID int64) error {
	funcName := "CrudJobUsecase.DeleteByID"
	captureFieldError := generalEntity.CaptureFields{
		"job_id": helper.ToString(jobID),
	}

	err := t.jobRepo.DeleteByID(ctx, nil, jobID)
	if err != nil {
		helper.LogError("jobRepo.DeleteByID", funcName, err, captureFieldError, "")

		return err
	}

	return nil
}
