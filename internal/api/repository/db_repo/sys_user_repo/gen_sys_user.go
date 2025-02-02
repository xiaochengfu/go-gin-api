///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package sys_user_repo

import (
	"fmt"

	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *SysUser {
	return new(SysUser)
}

func NewQueryBuilder() *sysUserRepoQueryBuilder {
	return new(sysUserRepoQueryBuilder)
}

func (t *SysUser) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type sysUserRepoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *sysUserRepoQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *sysUserRepoQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&SysUser{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *sysUserRepoQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&SysUser{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *sysUserRepoQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&SysUser{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *sysUserRepoQueryBuilder) First(db *gorm.DB) (*SysUser, error) {
	ret := &SysUser{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *sysUserRepoQueryBuilder) QueryOne(db *gorm.DB) (*SysUser, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *sysUserRepoQueryBuilder) QueryAll(db *gorm.DB) ([]*SysUser, error) {
	var ret []*SysUser
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *sysUserRepoQueryBuilder) Limit(limit int) *sysUserRepoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *sysUserRepoQueryBuilder) Offset(offset int) *sysUserRepoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereId(p db_repo.Predicate, value int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereIdIn(value []int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereIdNotIn(value []int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) OrderById(asc bool) *sysUserRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereUsername(p db_repo.Predicate, value string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", p),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereUsernameIn(value []string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", "IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereUsernameNotIn(value []string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "username", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) OrderByUsername(asc bool) *sysUserRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "username "+order)
	return qb
}

func (qb *sysUserRepoQueryBuilder) WherePassword(p db_repo.Predicate, value string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "password", p),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WherePasswordIn(value []string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "password", "IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WherePasswordNotIn(value []string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "password", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) OrderByPassword(asc bool) *sysUserRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "password "+order)
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereNickname(p db_repo.Predicate, value string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", p),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereNicknameIn(value []string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", "IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereNicknameNotIn(value []string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "nickname", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) OrderByNickname(asc bool) *sysUserRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "nickname "+order)
	return qb
}

func (qb *sysUserRepoQueryBuilder) WherePhone(p db_repo.Predicate, value string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phone", p),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WherePhoneIn(value []string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phone", "IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WherePhoneNotIn(value []string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "phone", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) OrderByPhone(asc bool) *sysUserRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "phone "+order)
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereEmail(p db_repo.Predicate, value string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "email", p),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereEmailIn(value []string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "email", "IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereEmailNotIn(value []string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "email", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) OrderByEmail(asc bool) *sysUserRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "email "+order)
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereAvatar(p db_repo.Predicate, value string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "avatar", p),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereAvatarIn(value []string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "avatar", "IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereAvatarNotIn(value []string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "avatar", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) OrderByAvatar(asc bool) *sysUserRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "avatar "+order)
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereSex(p db_repo.Predicate, value int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sex", p),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereSexIn(value []int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sex", "IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereSexNotIn(value []int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "sex", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) OrderBySex(asc bool) *sysUserRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "sex "+order)
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereCreateBy(p db_repo.Predicate, value string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_by", p),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereCreateByIn(value []string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_by", "IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereCreateByNotIn(value []string) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_by", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) OrderByCreateBy(asc bool) *sysUserRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_by "+order)
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereStatus(p db_repo.Predicate, value int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", p),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereStatusIn(value []int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereStatusNotIn(value []int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "status", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) OrderByStatus(asc bool) *sysUserRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "status "+order)
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereLastLoginTime(p db_repo.Predicate, value int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "last_login_time", p),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereLastLoginTimeIn(value []int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "last_login_time", "IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereLastLoginTimeNotIn(value []int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "last_login_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) OrderByLastLoginTime(asc bool) *sysUserRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "last_login_time "+order)
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereUpdateTime(p db_repo.Predicate, value int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_time", p),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereUpdateTimeIn(value []int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_time", "IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereUpdateTimeNotIn(value []int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "update_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) OrderByUpdateTime(asc bool) *sysUserRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "update_time "+order)
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereCreateTime(p db_repo.Predicate, value int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", p),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereCreateTimeIn(value []int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereCreateTimeNotIn(value []int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "create_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) OrderByCreateTime(asc bool) *sysUserRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "create_time "+order)
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereDeleteTime(p db_repo.Predicate, value int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "delete_time", p),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereDeleteTimeIn(value []int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "delete_time", "IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) WhereDeleteTimeNotIn(value []int32) *sysUserRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "delete_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *sysUserRepoQueryBuilder) OrderByDeleteTime(asc bool) *sysUserRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "delete_time "+order)
	return qb
}
