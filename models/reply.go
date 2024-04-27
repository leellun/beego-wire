package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Reply struct {
	Id           int       `orm:"column(id);pk"`
	CommentId    int64     `orm:"column(comment_id)"`
	Pid          int64     `orm:"column(pid)"`
	Content      string    `orm:"column(content)"`
	FromUserId   int64     `orm:"column(from_user_id)"`
	FromNickname string    `orm:"column(from_nickname)"`
	FromAvatar   string    `orm:"column(from_avatar)"`
	ToUserId     int64     `orm:"column(to_user_id)"`
	ToNickname   string    `orm:"column(to_nickname)"`
	ToAvatar     string    `orm:"column(to_avatar)"`
	Status       int16     `orm:"column(status)"`
	CreateTime   time.Time `orm:"column(create_time);type(timestamp without time zone)"`
	UpdateTime   time.Time `orm:"column(update_time);type(timestamp without time zone);null"`
}

func (t *Reply) TableName() string {
	return "reply"
}

func init() {
	orm.RegisterModel(new(Reply))
}

// AddReply insert a new Reply into database and returns
// last inserted Id on success.
func AddReply(m *Reply) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetReplyById retrieves Reply by Id. Returns error if
// Id doesn't exist
func GetReplyById(id int) (v *Reply, err error) {
	o := orm.NewOrm()
	v = &Reply{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllReply retrieves all Reply matches certain condition. Returns empty list if
// no records exist
func GetAllReply(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Reply))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Reply
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateReply updates Reply by Id and returns error if
// the record to be updated doesn't exist
func UpdateReplyById(m *Reply) (err error) {
	o := orm.NewOrm()
	v := Reply{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteReply deletes Reply by Id and returns error if
// the record to be deleted doesn't exist
func DeleteReply(id int) (err error) {
	o := orm.NewOrm()
	v := Reply{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Reply{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
