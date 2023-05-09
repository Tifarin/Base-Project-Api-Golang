package model

import "time"

type User struct {
    ID           int64      `db:"id"`
    Username     string     `db:"username"`
    Names        string     `db:"names"`
    Email        string     `db:"email"`
    PasswordHash string     `db:"password_hash"`
    RoleID       int64      `db:"id_roles"`
    CreatedAt    time.Time  `db:"created_at"`
    UpdatedAt    *time.Time `db:"updated_at"`
}
