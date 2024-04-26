package models

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"reflect"
	"strings"
	"time"
)

type FilesMedia struct {
	Id                int           `orm:"column(id);pk"`
	AddDate           time.Time     `orm:"column(add_date);type(timestamp with time zone);null"`
	AllowDownload     bool          `orm:"column(allow_download)"`
	Description       string        `orm:"column(description)"`
	Dislikes          int           `orm:"column(dislikes)"`
	Duration          int           `orm:"column(duration)"`
	EditDate          time.Time     `orm:"column(edit_date);type(timestamp with time zone)"`
	EnableComments    bool          `orm:"column(enable_comments)"`
	EncodingStatus    string        `orm:"column(encoding_status)"`
	Featured          bool          `orm:"column(featured)"`
	FriendlyToken     string        `orm:"column(friendly_token)"`
	HlsFile           string        `orm:"column(hls_file)"`
	IsReviewed        bool          `orm:"column(is_reviewed)"`
	Likes             int           `orm:"column(likes)"`
	Listable          bool          `orm:"column(listable)"`
	Md5sum            string        `orm:"column(md5sum);null"`
	MediaFile         string        `orm:"column(media_file)"`
	MediaInfo         string        `orm:"column(media_info)"`
	MediaType         string        `orm:"column(media_type)"`
	Password          string        `orm:"column(password)"`
	PreviewFilePath   string        `orm:"column(preview_file_path)"`
	Poster            string        `orm:"column(poster)"`
	ReportedTimes     int           `orm:"column(reported_times)"`
	Search            string        `orm:"column(search);null"`
	Size              string        `orm:"column(size);null"`
	Sprites           string        `orm:"column(sprites)"`
	State             string        `orm:"column(state)"`
	Title             string        `orm:"column(title)"`
	Thumbnail         string        `orm:"column(thumbnail)"`
	ThumbnailTime     float64       `orm:"column(thumbnail_time);null"`
	Uid               string        `orm:"column(uid);type(uuid)"`
	UploadedThumbnail string        `orm:"column(uploaded_thumbnail)"`
	UploadedPoster    string        `orm:"column(uploaded_poster)"`
	UserFeatured      bool          `orm:"column(user_featured)"`
	VideoHeight       int           `orm:"column(video_height)"`
	Views             int           `orm:"column(views)"`
	ChannelId         *UsersChannel `orm:"column(channel_id);rel(fk)"`
	LicenseId         *FilesLicense `orm:"column(license_id);rel(fk)"`
	UserId            *UsersUser    `orm:"column(user_id);rel(fk)"`
}

func (t *FilesMedia) TableName() string {
	return "files_media"
}

func init() {
	orm.RegisterModel(new(FilesMedia))
}

// AddFilesMedia insert a new FilesMedia into database and returns
// last inserted Id on success.
func AddFilesMedia(m *FilesMedia) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFilesMediaById retrieves FilesMedia by Id. Returns error if
// Id doesn't exist
func GetFilesMediaById(id int) (v *FilesMedia, err error) {
	o := orm.NewOrm()
	v = &FilesMedia{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFilesMedia retrieves all FilesMedia matches certain condition. Returns empty list if
// no records exist
func GetAllFilesMedia(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FilesMedia))
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

	var l []FilesMedia
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

// UpdateFilesMedia updates FilesMedia by Id and returns error if
// the record to be updated doesn't exist
func UpdateFilesMediaById(m *FilesMedia) (err error) {
	o := orm.NewOrm()
	v := FilesMedia{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFilesMedia deletes FilesMedia by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFilesMedia(id int) (err error) {
	o := orm.NewOrm()
	v := FilesMedia{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FilesMedia{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
