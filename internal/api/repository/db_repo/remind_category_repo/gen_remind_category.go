///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package remind_category_repo

import (
	"fmt"

	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *RemindCategory {
	return new(RemindCategory)
}

func NewQueryBuilder() *remindCategoryRepoQueryBuilder {
	return new(remindCategoryRepoQueryBuilder)
}

func (t *RemindCategory) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type remindCategoryRepoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *remindCategoryRepoQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *remindCategoryRepoQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&RemindCategory{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *remindCategoryRepoQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&RemindCategory{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *remindCategoryRepoQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&RemindCategory{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *remindCategoryRepoQueryBuilder) First(db *gorm.DB) (*RemindCategory, error) {
	ret := &RemindCategory{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *remindCategoryRepoQueryBuilder) QueryOne(db *gorm.DB) (*RemindCategory, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *remindCategoryRepoQueryBuilder) QueryAll(db *gorm.DB) ([]*RemindCategory, error) {
	var ret []*RemindCategory
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *remindCategoryRepoQueryBuilder) Limit(limit int) *remindCategoryRepoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) Offset(offset int) *remindCategoryRepoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) WhereId(p db_repo.Predicate, value int32) *remindCategoryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) WhereIdIn(value []int32) *remindCategoryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) WhereIdNotIn(value []int32) *remindCategoryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) OrderById(asc bool) *remindCategoryRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) WhereName(p db_repo.Predicate, value string) *remindCategoryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) WhereNameIn(value []string) *remindCategoryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) WhereNameNotIn(value []string) *remindCategoryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) OrderByName(asc bool) *remindCategoryRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) WhereCreateTime(p db_repo.Predicate, value int32) *remindCategoryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) WhereCreateTimeIn(value []int32) *remindCategoryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) WhereCreateTimeNotIn(value []int32) *remindCategoryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) OrderByCreateTime(asc bool) *remindCategoryRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) WhereUpdateTime(p db_repo.Predicate, value int32) *remindCategoryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_time", p),
		value,
	})
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) WhereUpdateTimeIn(value []int32) *remindCategoryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_time", "IN"),
		value,
	})
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) WhereUpdateTimeNotIn(value []int32) *remindCategoryRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *remindCategoryRepoQueryBuilder) OrderByUpdateTime(asc bool) *remindCategoryRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "update_time "+order)
	return qb
}
