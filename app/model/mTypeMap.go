package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"time"
)

var (
	_                          MTypeMapModel = (*defaultMTypeMapModel)(nil)
	cacheMTypeMapPrimaryPrefix               = "cache:MTypeMap:primary:"
)

type (
	MTypeMap struct {
		Id          int64      `gorm:"id;primary_key"` //
		CreateUser  int64      `gorm:"create_user"`    //id
		CreateTime  *time.Time `gorm:"create_time"`    //
		UpdateTime  *time.Time `gorm:"update_time"`    //
		DeleteTime  *time.Time `gorm:"delete_time"`    //
		DeleteState int64      `gorm:"delete_state"`   //
		Pid         int64      `gorm:"pid"`            //上级分类id
		Name        string     `gorm:"name"`           //类型名称
		Des         string     `gorm:"des"`            //描述
		Status      int64      `gorm:"status"`         //状态(1启用用，2停用)
		Ext1        string     `gorm:"ext1"`           //
		Ext2        string     `gorm:"ext2"`           //
		Ext3        string     `gorm:"ext3"`           //
	}

	MTypeMapModel interface {
		// Trans Transaction
		Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error)
		// Builder Custom assembly conditions
		Builder(ctx context.Context) (b *gorm.DB)
		Insert(ctx context.Context, db *gorm.DB, data *MTypeMap) (err error)
		Update(ctx context.Context, db *gorm.DB, data *MTypeMap) (err error)
		// Delete When a tombstone field is set, it is tombstone. Otherwise, it is physically deleted
		Delete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		// ForceDelete Physical deletion
		ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error)
		Count(ctx context.Context, cond *MTypeMapQuery) (total int64, err error)
		FindOne(ctx context.Context, Id int64) (data *MTypeMap, err error)
		// FindListByPage Normal pagination
		FindListByPage(ctx context.Context, cond *MTypeMapQuery) (list []*MTypeMap, err error)
		// FindListByCursor Cursor is required based on cursor paging
		FindListByCursor(ctx context.Context, cond *MTypeMapQuery) (list []*MTypeMap, err error)
		FindAll(ctx context.Context, cond *MTypeMapQuery) (list []*MTypeMap, err error)
		// ---------------Write your other interfaces below---------------
	}

	defaultMTypeMapModel struct {
		*customConn
		table string
	}
)

func NewMTypeMapMode(db *gorm.DB, c cache.CacheConf) MTypeMapModel {
	return &defaultMTypeMapModel{
		customConn: newCustomConn(db, c),
		table:      "m_type_map",
	}
}

func (m *defaultMTypeMapModel) conn(ctx context.Context, db *gorm.DB) *gorm.DB {
	if db == nil {
		return m.db.Table(m.table).Session(&gorm.Session{Context: ctx})
	}
	return db
}

func (m *defaultMTypeMapModel) Builder(ctx context.Context) (b *gorm.DB) {
	return m.conn(ctx, nil)
}

func (m *defaultMTypeMapModel) Trans(ctx context.Context, fn func(ctx context.Context, db *gorm.DB) (err error)) (err error) {
	return m.conn(ctx, nil).Transaction(func(tx *gorm.DB) error {
		return fn(ctx, tx)
	})
}

func (m *defaultMTypeMapModel) Insert(ctx context.Context, db *gorm.DB, data *MTypeMap) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMTypeMapPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Create(data).Error
	}, cacheKey)
}

func (m *defaultMTypeMapModel) Update(ctx context.Context, db *gorm.DB, data *MTypeMap) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMTypeMapPrimaryPrefix, data.Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Updates(data).Error
	}, cacheKey)
}

func (m *defaultMTypeMapModel) Delete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMTypeMapPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMTypeMapModel) ForceDelete(ctx context.Context, db *gorm.DB, Id int64) (err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMTypeMapPrimaryPrefix, Id)
	return m.Exec(ctx, func() error {
		return m.conn(ctx, db).Unscoped().Delete(Id).Error
	}, cacheKey)
}

func (m *defaultMTypeMapModel) Count(ctx context.Context, cond *MTypeMapQuery) (total int64, err error) {
	err = m.conn(ctx, nil).Where(cond.MTypeMap).Count(&total).Error
	return total, err
}

func (m *defaultMTypeMapModel) FindOne(ctx context.Context, Id int64) (data *MTypeMap, err error) {
	cacheKey := fmt.Sprintf("%s%v", cacheMTypeMapPrimaryPrefix, Id)
	err = m.QueryRow(ctx, &data, func(v interface{}) error {
		tx := m.conn(ctx, nil).Find(v, Id)
		if tx.RowsAffected == 0 {
			return sql.ErrNoRows
		}
		return tx.Error
	}, cacheKey)
	switch err {
	case nil:
		return data, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}
}

func (m *defaultMTypeMapModel) FindListByPage(ctx context.Context, cond *MTypeMapQuery) (list []*MTypeMap, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
		pageScope(cond.PageCurrent, cond.PageSize),
	)
	err = conn.Where(cond.MTypeMap).Find(&list).Error
	return list, err
}

func (m *defaultMTypeMapModel) FindListByCursor(ctx context.Context, cond *MTypeMapQuery) (list []*MTypeMap, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		cursorScope(cond.Cursor, cond.CursorAsc, cond.PageSize, "id"),
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MTypeMap).Find(&list).Error
	return list, err
}

func (m *defaultMTypeMapModel) FindAll(ctx context.Context, cond *MTypeMapQuery) (list []*MTypeMap, err error) {
	conn := m.conn(ctx, nil)
	conn = conn.Scopes(
		orderScope(cond.OrderSort...),
	)
	err = conn.Where(cond.MTypeMap).Find(&list).Error
	return list, err
}
