package mysql

import (
	"context"

	"github.com/kriti-mauludin/final-project/config"
	"github.com/kriti-mauludin/final-project/internal/helper"
	"github.com/kriti-mauludin/final-project/internal/repository/mysql/entity"

	apperr "github.com/kriti-mauludin/final-project/error"

	errwrap "github.com/pkg/errors"
	"gorm.io/gorm"
)

type IJobRepository interface {
	TrxSupportRepo
	GetByUserID(ctx context.Context, ID int64) (result []*entity.Job, err error)
	GetByID(ctx context.Context, ID int64) (result *entity.Job, err error)
	Create(ctx context.Context, dbTrx TrxObj, params *entity.Job, nonZeroVal bool) error
	LockByID(ctx context.Context, dbTrx TrxObj, ID int64) (result *entity.Job, err error)
	Update(ctx context.Context, dbTrx TrxObj, params *entity.Job, changes *entity.Job) (err error)
	DeleteByID(ctx context.Context, dbTrx TrxObj, id int64) error
}

type JobRepository struct {
	GormTrxSupport
}

func NewJobRepository(mysql *config.Mysql) *JobRepository {
	return &JobRepository{GormTrxSupport{db: mysql.DB}}
}

func (r *JobRepository) GetByUserID(ctx context.Context, userID int64) (result []*entity.Job, err error) {
	funcName := "JobRepository.GetByUserID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	err = r.db.Raw("SELECT * FROM job WHERE user_id = ?", userID).Scan(&result).Error
	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperr.ErrRecordNotFound()
	}

	return result, err
}

func (r *JobRepository) GetByID(ctx context.Context, ID int64) (result *entity.Job, err error) {
	funcName := "JobRepository.GetByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	err = r.db.Raw("SELECT * FROM job WHERE id = ? LIMIT 1", ID).Scan(&result).Error
	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperr.ErrRecordNotFound()
	}

	return result, err
}

func (r *JobRepository) Create(ctx context.Context, dbTrx TrxObj, params *entity.Job, nonZeroVal bool) error {
	funcName := "JobRepository.Create"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	cols := helper.NonZeroCols(params, nonZeroVal)
	return r.Trx(dbTrx).Select(cols).Create(&params).Error
}

func (r *JobRepository) LockByID(ctx context.Context, dbTrx TrxObj, ID int64) (result *entity.Job, err error) {
	funcName := "JobRepository.LockByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	err = r.Trx(dbTrx).
		Raw("SELECT * FROM job WHERE id = ? FOR UPDATE", ID).
		Scan(&result).Error

	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperr.ErrRecordNotFound()
	}

	return result, err
}

func (r *JobRepository) Update(ctx context.Context, dbTrx TrxObj, params *entity.Job, changes *entity.Job) (err error) {
	funcName := "JobRepository.Update"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	db := r.Trx(dbTrx).Model(params)
	if changes != nil {
		err = db.Updates(*changes).Error
	} else {
		err = db.Updates(helper.StructToMap(params, false)).Error
	}

	if err != nil {
		return errwrap.Wrap(err, funcName)
	}

	return nil
}

func (r *JobRepository) DeleteByID(ctx context.Context, dbTrx TrxObj, id int64) error {
	funcName := "JobRepository.DeleteByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	err := r.Trx(dbTrx).Where("id = ?", id).Delete(&entity.Job{}).Error
	if err != nil {
		return err
	}

	return nil
}
