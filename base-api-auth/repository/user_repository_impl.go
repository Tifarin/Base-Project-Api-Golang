package repository

import (
	"base-api-auth/helper"
	"base-api-auth/data/model"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct {

}
func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user model.User) model.User {
    SQL := "INSERT INTO user (username, names, email, password_hash, id_roles, created_at, updated_at) " +
        "VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"
    var id int64
    err := tx.QueryRowContext(ctx, SQL, user.Username, user.Names, user.Email, user.PasswordHash, user.RoleID, user.CreatedAt, user.UpdatedAt).Scan(&id)
    helper.PanicIfError(err)

    user.ID = id
    return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user model.User) model.User {
    SQL := "UPDATE user SET username=$1, names=$2, email=$3, password_hash=$4, id_roles=$5, created_at=$6, updated_at=$7 WHERE id=$8"
    _, err := tx.ExecContext(ctx, SQL, user.Username, user.Names, user.Email, user.PasswordHash, user.RoleID, user.CreatedAt, user.UpdatedAt, user.ID)
	helper.PanicIfError(err)
	return user
}
func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user model.User){
	SQL := "DELETE FROM user WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, user.ID)
	helper.PanicIfError(err)
}
func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (model.User, error) {
	SQL := "SELECT id, username, names, email, password_hash, id_roles, created_at, updated_at FROM user WHERE id = $1"
    rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := model.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Names, &user.Email, &user.PasswordHash, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
        helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}
func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []model.User {
	SQL := "SELECT id, username, names, email, password_hash, id_roles, created_at, updated_at FROM user"
    rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	var users []model.User
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Names, &user.Email, &user.PasswordHash, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)
        helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}