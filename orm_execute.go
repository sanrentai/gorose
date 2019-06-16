package gorose

import (
	"errors"
	"github.com/gohouse/t"
)

// Insert : insert data and get affected rows
func (dba *Orm) Insert(data ...interface{}) (int64, error) {
	if dba.GetData() == nil && len(data) > 0 {
		dba.data = data[0]
	}
	// 构建sql
	sqlStr, args, err := dba.BuildSql("insert")
	if err != nil {
		return 0, err
	}

	return dba.ISession.Execute(sqlStr, args...)
}

// insertGetId : insert data and get id
func (dba *Orm) InsertGetId() (int64, error) {
	_, err := dba.Insert()
	if err != nil {
		return 0, err
	}
	return dba.ISession.LastInsertId(), nil
}

// Update : update data
func (dba *Orm) Update() (int64, error) {
	// 构建sql
	sqlStr, args, err := dba.BuildSql("update")
	if err != nil {
		return 0, err
	}

	return dba.ISession.Execute(sqlStr, args...)
}

// Force 强制执行没有where的删除和修改
func (dba *Orm) Force() IOrm {
	dba.force = true
	return dba
}

// Delete : delete data
func (dba *Orm) Delete() (int64, error) {
	// 构建sql
	sqlStr, args, err := dba.BuildSql("delete")
	if err != nil {
		return 0, err
	}

	return dba.ISession.Execute(sqlStr, args...)
}

// Increment : auto Increment +1 default
// we can define step (such as 2, 3, 6 ...) if give the second params
// we can use this method as decrement with the third param as "-"
// orm.Increment("top") , orm.Increment("top", 2, "-")=orm.Decrement("top",2)
func (dba *Orm) Increment(args ...interface{}) (int64, error) {
	argLen := len(args)
	var field string
	var mode string = "+"
	var value string = "1"
	switch argLen {
	case 1:
		field = t.New(args[0]).String()
	case 2:
		field = t.New(args[0]).String()
		value = t.New(args[1]).String()
	case 3:
		field = t.New(args[0]).String()
		value = t.New(args[1]).String()
		mode = t.New(args[2]).String()
	default:
		return 0, errors.New("参数数量只允许1个,2个或3个")
	}
	dba.Data(field + "=" + field + mode + value)
	return dba.Update()
}

// Decrement : auto Decrement -1 default
// we can define step (such as 2, 3, 6 ...) if give the second params
func (dba *Orm) Decrement(args ...interface{}) (int64, error) {
	arglen := len(args)
	switch arglen {
	case 1:
		args = append(args, 1)
		args = append(args, "-")
	case 2:
		args = append(args, "-")
	default:
		return 0, errors.New("Decrement参数个数有误")
	}
	return dba.Increment(args...)
}
