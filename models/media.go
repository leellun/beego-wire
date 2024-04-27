package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Media struct {
	Id             int       `orm:"column(id);pk"`
	Description    string    `orm:"column(description)"`
	Dislikes       int       `orm:"column(dislikes)"`
	Duration       int       `orm:"column(duration)"`
	HlsFile        string    `orm:"column(hls_file)"`
	Likes          int       `orm:"column(likes)"`
	Md5sum         string    `orm:"column(md5sum);null"`
	MediaInfo      string    `orm:"column(media_info)"`
	MediaType      string    `orm:"column(media_type)"`
	Poster         string    `orm:"column(poster)"`
	Size           string    `orm:"column(size);null"`
	Thumbnail      string    `orm:"column(thumbnail)"`
	VideoHeight    int       `orm:"column(video_height)"`
	Views          int       `orm:"column(views)"`
	ChannelId      int       `orm:"column(channel_id);null"`
	UserId         int       `orm:"column(user_id)"`
	CreateTime     time.Time `orm:"column(create_time);type(timestamp with time zone);null"`
	UpdateTime     time.Time `orm:"column(update_time);type(timestamp with time zone)"`
	Status         int16     `orm:"column(status)"`
	Name           string    `orm:"column(name)"`
	Identification string    `orm:"column(identification);type(uuid)"`
}

func (t *Media) TableName() string {
	return "media"
}

func init() {
	orm.RegisterModel(new(Media))
}

// AddMedia insert a new Media into database and returns
// last inserted Id on success.
func AddMedia(m *Media) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMediaById retrieves Media by Id. Returns error if
// Id doesn't exist
func GetMediaById(id int) (v *Media, err error) {
	o := orm.NewOrm()
	v = &Media{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMedia retrieves all Media matches certain condition. Returns empty list if
// no records exist
func GetAllMedia(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Media))
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

	var l []Media
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

// UpdateMedia updates Media by Id and returns error if
// the record to be updated doesn't exist
func UpdateMediaById(m *Media) (err error) {
	o := orm.NewOrm()
	v := Media{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMedia deletes Media by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMedia(id int) (err error) {
	o := orm.NewOrm()
	v := Media{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Media{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
