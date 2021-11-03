///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package remind_library_repo

import (
	"fmt"

	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *RemindLibrary {
	return new(RemindLibrary)
}

func NewQueryBuilder() *remindLibraryRepoQueryBuilder {
	return new(remindLibraryRepoQueryBuilder)
}

func (t *RemindLibrary) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type remindLibraryRepoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *remindLibraryRepoQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *remindLibraryRepoQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&RemindLibrary{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *remindLibraryRepoQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&RemindLibrary{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *remindLibraryRepoQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&RemindLibrary{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *remindLibraryRepoQueryBuilder) First(db *gorm.DB) (*RemindLibrary, error) {
	ret := &RemindLibrary{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *remindLibraryRepoQueryBuilder) QueryOne(db *gorm.DB) (*RemindLibrary, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *remindLibraryRepoQueryBuilder) QueryAll(db *gorm.DB) ([]*RemindLibrary, error) {
	var ret []*RemindLibrary
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *remindLibraryRepoQueryBuilder) Limit(limit int) *remindLibraryRepoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) Offset(offset int) *remindLibraryRepoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereId(p db_repo.Predicate, value int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereIdIn(value []int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereIdNotIn(value []int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) OrderById(asc bool) *remindLibraryRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereType(p db_repo.Predicate, value int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", p),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereTypeIn(value []int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "IN"),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereTypeNotIn(value []int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) OrderByType(asc bool) *remindLibraryRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "type "+order)
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereUserId(p db_repo.Predicate, value int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", p),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereUserIdIn(value []int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", "IN"),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereUserIdNotIn(value []int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "user_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) OrderByUserId(asc bool) *remindLibraryRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "user_id "+order)
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereBody(p db_repo.Predicate, value string) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "body", p),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereBodyIn(value []string) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "body", "IN"),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereBodyNotIn(value []string) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "body", "NOT IN"),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) OrderByBody(asc bool) *remindLibraryRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "body "+order)
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereCategoryId(p db_repo.Predicate, value int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", p),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereCategoryIdIn(value []int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", "IN"),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereCategoryIdNotIn(value []int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "category_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) OrderByCategoryId(asc bool) *remindLibraryRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "category_id "+order)
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereCreateTime(p db_repo.Predicate, value int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereCreateTimeIn(value []int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereCreateTimeNotIn(value []int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) OrderByCreateTime(asc bool) *remindLibraryRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereUpdateTime(p db_repo.Predicate, value int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_time", p),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereUpdateTimeIn(value []int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_time", "IN"),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) WhereUpdateTimeNotIn(value []int32) *remindLibraryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *remindLibraryRepoQueryBuilder) OrderByUpdateTime(asc bool) *remindLibraryRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "update_time "+order)
	return qb
}
