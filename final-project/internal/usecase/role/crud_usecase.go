package todo_list_usecase

import (
	"context"
	"fmt"
	"time"

	generalEntity "github.com/kriti-mauludin/final-project/entity"
	"github.com/kriti-mauludin/final-project/internal/helper"
	"github.com/kriti-mauludin/final-project/internal/queue"
	"github.com/kriti-mauludin/final-project/internal/repository/mysql"
	mentity "github.com/kriti-mauludin/final-project/internal/repository/mysql/entity"
	"github.com/kriti-mauludin/final-project/internal/usecase"
	"github.com/kriti-mauludin/final-project/internal/usecase/role/entity"
	errwrap "github.com/pkg/errors"
)

type CrudRoleUsecase struct {
	roleRepo mysql.IRoleRepository
	queue    queue.Queue
}

func NewCrudRoleUsecase(
	roleRepo mysql.IRoleRepository,
	queue queue.Queue,
) *CrudRoleUsecase {
	return &CrudRoleUsecase{roleRepo, queue}
}

type ICrudRoleUsecase interface {
	GetByUserID(ctx context.Context, userID int64) (res []*entity.RoleResponse, err error)
	GetByID(ctx context.Context, roleID int64) (*entity.RoleResponse, error)
	Create(ctx context.Context, roleReq entity.RoleReq) (*entity.RoleResponse, error)
	UpdateByID(ctx context.Context, roleReq entity.RoleReq) error
	DeleteByID(ctx context.Context, roleID int64) error
}

func (t *CrudRoleUsecase) GetByUserID(ctx context.Context, userID int64) (res []*entity.RoleResponse, err error) {
	funcName := "CrudRoleUsecase.GetByUserID"
	captureFieldError := generalEntity.CaptureFields{
		"user_id": helper.ToString(userID),
	}

	result, err := t.roleRepo.GetByUserID(ctx, userID)
	if err != nil {
		helper.LogError("roleRepo.GetByUserID", funcName, err, captureFieldError, "")

		return nil, err
	}

	for _, v := range result {
		res = append(res, &entity.RoleResponse{
			ID:          v.ID,
			Role:        v.Role,
			Description: v.Description,
			DoingAt:     helper.ConvertToJakartaDate(v.DoingAt),
			CreatedAt:   helper.ConvertToJakartaTime(v.CreatedAt),
			UpdatedAt:   helper.ConvertToJakartaTime(v.UpdatedAt),
		})
	}

	return res, nil
}

func (t *CrudRoleUsecase) GetByID(ctx context.Context, roleID int64) (*entity.RoleResponse, error) {
	funcName := "CrudRoleUsecase.GetByID"
	captureFieldError := generalEntity.CaptureFields{
		"user_id": helper.ToString(roleID),
	}

	data, err := t.roleRepo.GetByID(ctx, roleID)
	if err != nil {
		helper.LogError("roleRepo.GetByID", funcName, err, captureFieldError, "")

		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	return &entity.RoleResponse{
		ID:          data.ID,
		Role:        data.Role,
		Description: data.Description,
		DoingAt:     helper.ConvertToJakartaDate(data.DoingAt),
		CreatedAt:   helper.ConvertToJakartaTime(data.CreatedAt),
		UpdatedAt:   helper.ConvertToJakartaTime(data.UpdatedAt),
	}, nil
}

func (t *CrudRoleUsecase) Create(ctx context.Context, roleReq entity.RoleReq) (*entity.RoleResponse, error) {
	funcName := "CrudRoleUsecase.Create"
	captureFieldError := generalEntity.CaptureFields{
		"user_id": helper.ToString(roleReq.UserID),
		"payload": helper.ToString(roleReq),
	}

	if errMsg := usecase.ValidateStruct(roleReq); errMsg != "" {
		return nil, errwrap.Wrap(fmt.Errorf(generalEntity.INVALID_PAYLOAD_CODE), errMsg)
	}

	doingAt, _ := helper.ParseDate(roleReq.DoingAt)

	rolePayload := &mentity.Role{
		UserID:      roleReq.UserID,
		Role:        roleReq.Role,
		Description: roleReq.Description,
		DoingAt:     doingAt,
		CreatedAt:   time.Now(),
	}

	err := t.roleRepo.Create(ctx, nil, rolePayload, false)
	if err != nil {
		helper.LogError("roleRepo.Create", funcName, err, captureFieldError, "")

		return nil, err
	}

	sendEmailReq := entity.SendEmailReq{
		UserID:    roleReq.UserID,
		CreatedAt: rolePayload.CreatedAt,
	}

	sendEmailReqJson, _ := helper.Serialize(sendEmailReq)
	err = t.queue.Publish(queue.ProcessSendEmail, sendEmailReqJson, 1)
	if err != nil {
		helper.LogError("queue.PublishMessage", funcName, err, generalEntity.CaptureFields{
			"topic":   queue.ProcessSendEmail,
			"payload": helper.ToString(sendEmailReq),
		}, "")
		return nil, err
	}

	return &entity.RoleResponse{
		ID:          rolePayload.ID,
		Role:        rolePayload.Role,
		Description: rolePayload.Description,
		DoingAt:     helper.ConvertToJakartaDate(rolePayload.DoingAt),
		CreatedAt:   helper.ConvertToJakartaTime(rolePayload.CreatedAt),
	}, nil
}

func (t *CrudRoleUsecase) UpdateByID(ctx context.Context, roleReq entity.RoleReq) error {
	funcName := "CrudRoleUsecase.UpdateByID"
	roleID := roleReq.ID

	captureFieldError := generalEntity.CaptureFields{
		"user_id": helper.ToString(roleReq.UserID),
		"payload": helper.ToString(roleReq),
	}

	// Start DB Transaction
	if err := mysql.DBTransaction(t.roleRepo, func(trx mysql.TrxObj) error {
		// Locking Data
		lockedData, err := t.roleRepo.LockByID(ctx, trx, roleID)
		if err != nil {
			helper.LogError("roleRepo.LockByID", funcName, err, captureFieldError, "")

			return err
		}
		if lockedData == nil {
			return fmt.Errorf("DATA IS NOT EXIST")
		}

		// Process Update
		doingAt, _ := helper.ParseDate(roleReq.DoingAt)
		if err := t.roleRepo.Update(ctx, trx, lockedData, &mentity.Role{
			Role:        roleReq.Role,
			Description: roleReq.Description,
			DoingAt:     doingAt,
			UpdatedAt:   time.Now(),
		}); err != nil {
			helper.LogError("roleRepo.Update", funcName, err, captureFieldError, "")

			return err
		}

		return nil
	}); err != nil {
		helper.LogError("roleRepo.DBTransaction", funcName, err, captureFieldError, "")

		return err
	}

	return nil
}

func (t *CrudRoleUsecase) DeleteByID(ctx context.Context, roleID int64) error {
	funcName := "CrudRoleUsecase.DeleteByID"
	captureFieldError := generalEntity.CaptureFields{
		"role_id": helper.ToString(roleID),
	}

	err := t.roleRepo.DeleteByID(ctx, nil, roleID)
	if err != nil {
		helper.LogError("roleRepo.DeleteByID", funcName, err, captureFieldError, "")

		return err
	}

	return nil
}
