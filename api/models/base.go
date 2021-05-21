package models

import (
	"database/sql"
	"time"
)

type Base struct {
	CreatedBy string
	CreatedAt sql.NullTime
	UpdatedBy string
	UpdatedAt sql.NullTime
	IsDeleted bool
	DeletedAt *sql.NullTime
}

func (base *Base) SetCreatedBy(user string) {
	base.CreatedBy = user
}

func (base *Base) SetCreatedAt(date time.Time) {
	nullTime := sql.NullTime{Time: date, Valid: true}
	base.CreatedAt = nullTime
}

func (base *Base) SetUpdatedBy(user string) {
	base.UpdatedBy = user
}

func (base *Base) SetUpdatedAt(date time.Time) {
	nullTime := sql.NullTime{Time: date, Valid: true}
	base.UpdatedAt = nullTime
}

func (base *Base) SetIsDeleted(deleted bool) {
	base.IsDeleted = deleted
}

func (base *Base) SetDeletedAt(date time.Time) {
	nullTime := sql.NullTime{Time: date, Valid: true}
	base.DeletedAt = &nullTime
}
