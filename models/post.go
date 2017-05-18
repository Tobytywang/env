package models

import (
	"errors"
	"strconv"
	_ "os"
	_ "strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 定义文章结构体（数据库表）
type Post struct {
	Id          int    `orm:"pk;auto"form:"id"`
	Name        string `orm:"size(100)"form:"name"`
	Link        string // 页面的路径
	Content     string `orm:"type(text)"form:"content"`
	Read        uint
}

// 新增一个文章
func PAddOne(post *Post) error{
  o := orm.NewOrm()

  // 
  _, err := o.Insert(post)
  if err != nil {
    // beego.Debug(err)
    return errors.New("添加文章失败(1)")
  }

  // 第二次：读取id构造链接
  o.Read(post)
  intid := (int)(post.Id)
  post.Link = "http://" + WEB_URL + "/post?id=" + strconv.Itoa(intid)
  _, err = o.Update(post)
  if err != nil {
    // beego.Debug(err)
    return errors.New("添加文章失败(2)")
  }

  // 以下是二维码专属代码，不再需要
  //  name := strings.Split(code.Name, ".")
  //  os.Mkdir(QRPATH, 0777)
  //  err = qrcode.WriteFile(code.Link, qrcode.Medium, 256, "static/" + QRPATH + "/" + strconv.Itoa(code.Id) + "-" + name[0] + ".png")
  //  if err != nil {
	// beego.Debug(err)
	// return errors.New("生成二维码失败！")
  //  }
  //  o.Read(code)

  //  beego.Debug(name)
  //  code.Code = "static/" + QRPATH + "/" + strconv.Itoa(code.Id) + "-" + name[0] + ".png"
  //  _, err = o.Update(code)
	// beego.Debug(code)
  //  if err != nil {
	// 	beego.Debug(err)
  //    return errors.New("生成二维码路径失败！")
  //  }
  return nil
}

// 更新一个文章
func PUpdate(post *Post) error{
  o := orm.NewOrm()
  beego.Debug(post)
  temp := Post{Id: post.Id}
  if o.Read(&temp) == nil {
    beego.Debug(temp)
    temp.Content = post.Content
    temp.Name = post.Name
    o.Update(&temp, "Name", "Content")
  }
  return nil
}

// 删除一个文章
func PDelOne(id int) error{
  o := orm.NewOrm()
  if _, err := o.Delete(&Post{Id: id}); err != nil {
    return errors.New("删除文章失败")
  }
  return nil
}

// 增加阅读数
func PRead(id int) error{
  beego.Debug(id)
  o := orm.NewOrm()
  temp := Post{Id: id}
  var err error
  if err := o.Read(&temp); err == nil {
    temp.Read = temp.Read + 1
    o.Update(&temp, "Read")
  }
  beego.Debug(err)
  return nil
}

// 计数（分页功能）
func PCountPosts() int64{
  o := orm.NewOrm()
  cnt, _ := o.QueryTable("post").Count()
  return cnt

}

// 根据偏移和数量获取二维码（分页功能）
func ListPostsByOffsetAndLimit(offset int, postperpage int) (postlist []Post){
  o := orm.NewOrm()
  // templist := make([]QRCode, 0)
  var templist []Post
  o.QueryTable("post").OrderBy("id").All(&templist)
  var top int
  if ((offset+postperpage)>len(templist)){
    top = len(templist)
  } else {
    top = (offset+postperpage)
  }
  postlist = templist[offset:top]
  return postlist
}

// 根据ID查找二维码
func PReadById(id int) (*Post, error){
  o := orm.NewOrm()
  a := new(Post)
  o.QueryTable("post").Filter("id", id).One(a)
  if a.Id == 0 {
    return a, errors.New("没有该数据")
  }
  return a, nil
}

// 查找所有文章
func PReadAll(posts *[]*Post) {
	o := orm.NewOrm()
	o.QueryTable("post").OrderBy("-id").All(posts)
}

// 根据Name或者Desc的内容查找匹配的二维码（查找功能）
func PSearch(content string) (posts []Post){
	o := orm.NewOrm()
	o.QueryTable("post").OrderBy("id").Filter("name__contains", content).Filter("content__contains", content).All(&posts)
	beego.Debug(posts)
	return
}



// 自定义表名
func (u *Post) TableName() string {
    return "post"
}

// 注册表
func init() {
	orm.RegisterModel(new(Post))
}