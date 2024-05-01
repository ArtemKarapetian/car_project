package model

import "time"

type Car struct {
	RegNum    string     `db:"reg_num"`
	Mark      string     `db:"mark"`
	Model     string     `db:"model"`
	Year      int32      `db:"year,omitempty"`
	Owner     *People    `db:"owner"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
