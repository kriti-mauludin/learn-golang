package todo_list_category_usecase

import (
	"context"
	"fmt"
	"time"

	generalEntity "github.com/kriti-mauludin/learn-golang/entity"

	"github.com/kriti-mauludin/learn-golang/internal/helper"
	"github.com/kriti-mauludin/learn-golang/internal/repository/mysql"
	myentity "github.com/kriti-mauludin/learn-golang/internal/repository/mysql/entity"
	"github.com/kriti-mauludin/learn-golang/internal/usecase/todo_list_category/entity"
)

//todo
// - Create
// - Get By ID
// - Update
// - Delete
// - Get All Data

type ICrudTodoListCategoryUsecase interface {
	Create(ctx context.Context, req entity.TodoListCategoryReq) error
	GetByID(ctx context.Context, ID int64) (*entity.TodoListCategoryResponse, error)
	UpdateByID(ctx context.Context, req entity.TodoListCategoryReq) error
	DeleteByID(ctx context.Context, ID int64) error
}

type CrudTodoListCategoryUsecase struct {
	TodoListCategoryRepo mysql.ITodoListCategoryRepository
}

func NewCrudTodoListCategoryUsecase(TodoListCategoryRepo mysql.ITodoListCategoryRepository) *CrudTodoListCategoryUsecase {
	return &CrudTodoListCategoryUsecase{
		TodoListCategoryRepo: TodoListCategoryRepo,
	}
}

func (u *CrudTodoListCategoryUsecase) Create(ctx context.Context, req entity.TodoListCategoryReq) error {
	funcName := "CrudTodoListCategoryUsecase.Create"

	logFields := generalEntity.CaptureFields{
		"name":        req.Name,
		"description": req.Description,
	}

	data := &myentity.TodoListCategory{
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now(),
	}

	err := u.TodoListCategoryRepo.Create(ctx, nil, data, false)
	if err != nil {
		helper.LogError("TodoListCategoryRepo.Create", funcName, err, logFields, "")
		return err
	}
	return nil
}

func (t *CrudTodoListCategoryUsecase) GetByID(ctx context.Context, todoListID int64) (*entity.TodoListCategoryResponse, error) {
	funcName := "CrudTodoListCategoryUsecase.GetByID"
	captureFieldError := generalEntity.CaptureFields{
		"user_id": helper.ToString(todoListID),
	}

	data, err := t.TodoListCategoryRepo.GetByID(ctx, todoListID)
	if err != nil {
		helper.LogError("todoListCategoryRepo.GetByID", funcName, err, captureFieldError, "")

		return nil, err
	}
	if data == nil {
		return nil, nil
	}

	return &entity.TodoListCategoryResponse{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		CreatedAt:   helper.ConvertToJakartaTime(data.CreatedAt),
	}, nil
}

func (t *CrudTodoListCategoryUsecase) UpdateByID(ctx context.Context, todoListReq entity.TodoListCategoryReq) error {
	funcName := "CrudTodoListCategoryUsecase.UpdateByID"
	todoListID := todoListReq.ID

	captureFieldError := generalEntity.CaptureFields{
		"payload": helper.ToString(todoListReq),
	}

	if err := mysql.DBTransaction(t.TodoListCategoryRepo, func(trx mysql.TrxObj) error {
		lockedData, err := t.TodoListCategoryRepo.LockByID(ctx, trx, todoListID)
		if err != nil {
			helper.LogError("todoListCategoryRepo.LockByID", funcName, err, captureFieldError, "")

			return err
		}
		if lockedData == nil {
			return fmt.Errorf("DATA IS NOT EXIST")
		}

		if err := t.TodoListCategoryRepo.Update(ctx, trx, lockedData, &myentity.TodoListCategory{
			Name:        todoListReq.Name,
			Description: todoListReq.Description,
		}); err != nil {
			helper.LogError("todoListCategoryRepo.Update", funcName, err, captureFieldError, "")

			return err
		}

		return nil
	}); err != nil {
		helper.LogError("todoListCategoryRepo.DBTransaction", funcName, err, captureFieldError, "")

		return err
	}

	return nil
}

func (t *CrudTodoListCategoryUsecase) DeleteByID(ctx context.Context, todoListID int64) error {
	funcName := "CrudTodoListCategoryUsecase.DeleteByID"
	captureFieldError := generalEntity.CaptureFields{
		"todo_list_category_id": helper.ToString(todoListID),
	}

	err := t.TodoListCategoryRepo.DeleteByID(ctx, nil, todoListID)
	if err != nil {
		helper.LogError("todoListCategoryRepo.DeleteByID", funcName, err, captureFieldError, "")

		return err
	}

	return nil
}
