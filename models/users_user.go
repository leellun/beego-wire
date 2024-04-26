package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type UsersUser struct {
	Id                     int       `orm:"column(id);pk"`
	Password               string    `orm:"column(password)"`
	LastLogin              time.Time `orm:"column(last_login);type(timestamp with time zone);null"`
	IsSuperUser            bool      `orm:"column(is_super_user)"`
	Username               string    `orm:"column(username)"`
	FirstName              string    `orm:"column(first_name)"`
	LastName               string    `orm:"column(last_name)"`
	Email                  string    `orm:"column(email)"`
	IsStaff                bool      `orm:"column(is_staff)"`
	IsActive               bool      `orm:"column(is_active)"`
	DateJoined             time.Time `orm:"column(date_joined);type(timestamp with time zone)"`
	Logo                   string    `orm:"column(logo)"`
	Description            string    `orm:"column(description)"`
	Name                   string    `orm:"column(name)"`
	DateAdded              time.Time `orm:"column(date_added);type(timestamp with time zone)"`
	IsFeatured             bool      `orm:"column(is_featured)"`
	Title                  string    `orm:"column(title)"`
	AdvancedUser           bool      `orm:"column(advanced_user)"`
	MediaCount             int       `orm:"column(media_count)"`
	NotificationOnComments bool      `orm:"column(notification_on_comments)"`
	AllowContact           bool      `orm:"column(allow_contact)"`
	Location               string    `orm:"column(location)"`
	IsEditor               bool      `orm:"column(is_editor)"`
	IsManager              bool      `orm:"column(is_manager)"`
}

func (t *UsersUser) TableName() string {
	return "users_user"
}

func init() {
	orm.RegisterModel(new(UsersUser))
}
func (t *UsersUser) IsMediacmsEditor() bool {
	//Whether user is MediaCMS editor
	editor := false
	if t.IsSuperUser || t.IsManager || t.IsEditor {
		editor = true
	}
	return editor
}
func (t *UsersUser) IsMediacmsManager() bool {
	//Whether user is MediaCMS manager
	manager := false
	if t.IsSuperUser || t.IsManager {
		manager = false
	}
	return manager
}

// AddUsersUser insert a new UsersUser into database and returns
// last inserted Id on success.
func AddUsersUser(m *UsersUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUsersUserById retrieves UsersUser by Id. Returns error if
// Id doesn't exist
func GetUsersUserById(id int) (v *UsersUser, err error) {
	o := orm.NewOrm()
	v = &UsersUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUsersUser retrieves all UsersUser matches certain condition. Returns empty list if
// no records exist
func GetAllUsersUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UsersUser))
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

	var l []UsersUser
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

// UpdateUsersUser updates UsersUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateUsersUserById(m *UsersUser) (err error) {
	o := orm.NewOrm()
	v := UsersUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUsersUser deletes UsersUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUsersUser(id int) (err error) {
	o := orm.NewOrm()
	v := UsersUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UsersUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
