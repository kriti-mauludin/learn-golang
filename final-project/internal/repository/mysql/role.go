package mysql

import (
	"context"

	"github.com/kriti-mauludin/try-consumer-rabbitmq/config"
	"github.com/kriti-mauludin/try-consumer-rabbitmq/internal/helper"
	"github.com/kriti-mauludin/try-consumer-rabbitmq/internal/repository/mysql/entity"

	apperr "github.com/kriti-mauludin/try-consumer-rabbitmq/error"

	errwrap "github.com/pkg/errors"
	"gorm.io/gorm"
)

type IRoleRepository interface {
	TrxSupportRepo
	GetByUserID(ctx context.Context, ID int64) (result []*entity.Role, err error)
	GetByID(ctx context.Context, ID int64) (result *entity.Role, err error)
	Create(ctx context.Context, dbTrx TrxObj, params *entity.Role, nonZeroVal bool) error
	LockByID(ctx context.Context, dbTrx TrxObj, ID int64) (result *entity.Role, err error)
	Update(ctx context.Context, dbTrx TrxObj, params *entity.Role, changes *entity.Role) (err error)
	DeleteByID(ctx context.Context, dbTrx TrxObj, id int64) error
}

type RoleRepository struct {
	GormTrxSupport
}

func NewRoleRepository(mysql *config.Mysql) *RoleRepository {
	return &RoleRepository{GormTrxSupport{db: mysql.DB}}
}

func (r *RoleRepository) GetByUserID(ctx context.Context, userID int64) (result []*entity.Role, err error) {
	funcName := "RoleRepository.GetByUserID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	err = r.db.Raw("SELECT * FROM role WHERE user_id = ?", userID).Scan(&result).Error
	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperr.ErrRecordNotFound()
	}

	return result, err
}

func (r *RoleRepository) GetByID(ctx context.Context, ID int64) (result *entity.Role, err error) {
	funcName := "RoleRepository.GetByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	err = r.db.Raw("SELECT * FROM role WHERE id = ? LIMIT 1", ID).Scan(&result).Error
	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperr.ErrRecordNotFound()
	}

	return result, err
}

func (r *RoleRepository) Create(ctx context.Context, dbTrx TrxObj, params *entity.Role, nonZeroVal bool) error {
	funcName := "RoleRepository.Create"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	cols := helper.NonZeroCols(params, nonZeroVal)
	return r.Trx(dbTrx).Select(cols).Create(&params).Error
}

func (r *RoleRepository) LockByID(ctx context.Context, dbTrx TrxObj, ID int64) (result *entity.Role, err error) {
	funcName := "RoleRepository.LockByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return nil, errwrap.Wrap(err, funcName)
	}

	err = r.Trx(dbTrx).
		Raw("SELECT * FROM role WHERE id = ? FOR UPDATE", ID).
		Scan(&result).Error

	if errwrap.Is(err, gorm.ErrRecordNotFound) {
		return nil, apperr.ErrRecordNotFound()
	}

	return result, err
}

func (r *RoleRepository) Update(ctx context.Context, dbTrx TrxObj, params *entity.Role, changes *entity.Role) (err error) {
	funcName := "RoleRepository.Update"

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

func (r *RoleRepository) DeleteByID(ctx context.Context, dbTrx TrxObj, id int64) error {
	funcName := "RoleRepository.DeleteByID"

	if err := helper.CheckDeadline(ctx); err != nil {
		return errwrap.Wrap(err, funcName)
	}

	err := r.Trx(dbTrx).Where("id = ?", id).Delete(&entity.Role{}).Error
	if err != nil {
		return err
	}

	return nil
}
