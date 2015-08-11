package models

import (
	// "fmt"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	_DB_NAME        = "data/Go.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"auto_now_add;type(datetime);index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"null;type(datetime);index"`
	TopicCount      int64
	TopicLastUserId int64
}
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"null;auto_now_add;type(datetime);index"`
	Updated         time.Time `orm:"auto_now;type(datetime);index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"null;type(datetime);index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}
func AddCategory(name string) error {
	o := orm.NewOrm()
	cate := &Category{Title: name}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}
func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}
func AddTopic(title, content string) error {
	// fmt.Sprint(title, content)
	o := orm.NewOrm()
	topic := &Topic{
		Title:   title,
		Content: content,
	}
	_, err := o.Insert(topic)
	return err
}

func ModifyTopic(tid, title, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Content = content
		topic.Updated = time.Now()
		o.Update(topic)
	}
	return nil
	// tid, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	return err
	// }
	// o := orm.NewOrm()
	// topic := &Topic{
	// 	Id:      tid,
	// 	Title:   title,
	// 	Content: content,
	// }
	// _, err = o.Update(topic)
	// return err
}
func GetTopic(id string) (*Topic, error) {

	tid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	t := &Topic{Id: tid}
	qs := o.QueryTable("Topic")
	err = qs.Filter("Id", tid).One(t)
	return t, err
}
func GetAllTopics() ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("Topic")
	_, err := qs.All(&topics)
	return topics, err
}
