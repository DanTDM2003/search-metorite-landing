package database

import (
	"context"
	"errors"

	"github.com/DanTDM2003/search-api-docker-redis/internal/models"
	"github.com/DanTDM2003/search-api-docker-redis/internal/users/repository"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/paginator"
	"github.com/DanTDM2003/search-api-docker-redis/pkg/postgres"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

const (
	usersTable = "users"
)

func (repo impleRepository) getTable() *gorm.DB {
	return repo.db.Table(usersTable)
}

func (repo impleRepository) GetUsers(ctx context.Context, opt repository.GetUsersOptions) ([]models.User, paginator.Paginator, error) {
	table := repo.getTable()

	var total int64
	if err := table.Count(&total).Error; err != nil {
		repo.l.Errorf(ctx, "users.repository.database.GetUsers.db.Count: %v", err)
		return nil, paginator.Paginator{}, err
	}

	cursor := table.
		Offset(int(opt.PaginatorQuery.Offset())).
		Limit(int(opt.Limit))

	var users []models.User
	if err := cursor.Find(&users).Error; err != nil {
		repo.l.Errorf(ctx, "users.repository.database.GetUsers.db.Find: %v", err)
		return nil, paginator.Paginator{}, err
	}

	return users, paginator.Paginator{
		Total:       total,
		Count:       int64(len(users)),
		PerPage:     opt.Limit,
		CurrentPage: opt.Page,
	}, nil
}

func (repo impleRepository) GetOneUser(ctx context.Context, opt repository.GetOneUserOptions) (models.User, error) {
	table := repo.getTable()

	cond, params := repo.buildGetOneUserCondition(opt)

	var user models.User
	if err := table.Where(cond, params).First(&user).Error; err != nil {
		repo.l.Errorf(ctx, "users.repository.database.GetOneUser.db.First: %v", err)
		return models.User{}, err
	}

	return user, nil
}

func (repo impleRepository) CreateUser(ctx context.Context, opt repository.CreateUserOptions) (models.User, error) {
	table := repo.getTable()

	user := repo.buildCreateUserModel(opt)

	if err := table.Create(&user).Error; err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == postgres.ErrDuplicatedKeyCode {
			repo.l.Warnf(ctx, "users.repository.database.CreateUser.db.Create: %v", err)
			return models.User{}, gorm.ErrCheckConstraintViolated
		}
		repo.l.Errorf(ctx, "users.repository.database.CreateUser.db.Create: %v", err)
		return models.User{}, err
	}

	return user, nil
}

func (repo impleRepository) UpdateUser(ctx context.Context, opt repository.UpdateUserOptions, user models.User) (models.User, error) {
	table := repo.getTable()

	update := repo.buildUpdateUserModel(opt, user)

	if err := table.Where("id = ?", user.ID).Updates(&update).Error; err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == postgres.ErrDuplicatedKeyCode {
			repo.l.Warnf(ctx, "users.repository.database.UpdateUser.db.Updates: %v", err)
			return models.User{}, gorm.ErrCheckConstraintViolated
		}
		repo.l.Errorf(ctx, "users.repository.database.UpdateUser.db.Updates: %v", err)
		return models.User{}, nil
	}

	return update, nil
}

func (repo impleRepository) DeleteUser(ctx context.Context, id uint) error {
	table := repo.getTable()

	if err := table.Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		repo.l.Errorf(ctx, "users.repository.database.DeleteUser.db.Delete: %v", err)
		return err
	}

	return nil
}
