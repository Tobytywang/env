// qrcode提供二维码的增删改查功能
package models

import (
	"errors"
	"strconv"
	"os"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	qrcode "github.com/skip2/go-qrcode"
)

// 定义二维码结构体（数据库表）
type QRCode struct {
	Id        int    `orm:"pk;auto"form:"id"`
	Name      string `orm:"size(100)"form:"name"`
	Link		  string // 页面的路径
	Pic       string // 图片存储的路径
	Code		  string // 二维码存储的路径
	Desc      string `orm:"type(text)"form:"desc-html"`
	Markdown  string `orm:"type(text)"form:"desc-markdown"`
	Read      uint
}

// 定义上传二维码的储存位置
var (
	QRPATH string = beego.AppConfig.String("QRPATH")
	WEB_URL string = beego.AppConfig.String("WEB_URL")
)

// 新增一个二维码
func QRAddOne(code *QRCode) error{
  o := orm.NewOrm()

  _, err := o.Insert(code)
  if err != nil {
		beego.Debug(err)
    return errors.New("添加植物失败(1)")
  }

  o.Read(code)
  intid := (int)(code.Id)
  code.Link = "http://" + WEB_URL + "/plant?id=" + strconv.Itoa(intid)
  _, err = o.Update(code)
  if err != nil {
		beego.Debug(err)
    return errors.New("添加植物失败(2)")
  }

	name := strings.Split(code.Name, ".")
	os.Mkdir(QRPATH, 0777)
	err = qrcode.WriteFile(code.Link, qrcode.Medium, 256, "static/" + QRPATH + "/" + strconv.Itoa(code.Id) + "-" + name[0] + ".png")
	if err != nil {
		beego.Debug(err)
		return errors.New("生成二维码失败！")
	}
	o.Read(code)

	beego.Debug(name)
	code.Code = "static/" + QRPATH + "/" + strconv.Itoa(code.Id) + "-" + name[0] + ".png"
  _, err = o.Update(code)
	beego.Debug(code)
  if err != nil {
		beego.Debug(err)
    return errors.New("生成二维码路径失败！")
  }
  return nil
}

// 更新一个二维码
func QRUpdate(code *QRCode) error{
	o := orm.NewOrm()
	beego.Debug(code)
	temp := QRCode{Id: code.Id}
	if o.Read(&temp) == nil {
		beego.Debug(temp)
		temp.Desc = code.Desc
		temp.Markdown = code.Markdown
		o.Update(&temp, "Desc", "Markdown")
	}
	return nil
}

// 增加二维码的阅读次数
func QRRead(id int) error{
	beego.Debug("批准")
	beego.Debug(id)
	o := orm.NewOrm()
	temp := QRCode{Id: id}
	var err error
	if err := o.Read(&temp); err == nil {
		temp.Read = temp.Read + 1
		o.Update(&temp, "Read")
	}
	beego.Debug(err)
	return nil
}
// 删除一个二维码
func QRDel(id int) error{
  o := orm.NewOrm()
	if _, err := o.Delete(&QRCode{Id: id}); err != nil {
		return errors.New("删除植物失败")
	}
	return nil
}

// 查找所有二维码
func QRReadAll(qrlist *[]*QRCode) {
	o := orm.NewOrm()
	o.QueryTable("qrcode").OrderBy("-id").All(qrlist)
}

// 返回数据库中二维码的数目（分页功能）
func CountCodes() int64{
	o := orm.NewOrm()
	cnt, _ := o.QueryTable("qrcode").Count()
  return cnt
}

// 根据偏移和数量获取二维码（分页功能）
func ListCodesByOffsetAndLimit(offset int, codeperpage int) (qrlist []QRCode){
	o := orm.NewOrm()
	// templist := make([]QRCode, 0)
	var templist []QRCode
	o.QueryTable("qrcode").OrderBy("id").All(&templist)
	var top int
	if ((offset+codeperpage)>len(templist)){
		top = len(templist)
	} else {
		top = (offset+codeperpage)
	}
	qrlist = templist[offset:top]
	return qrlist
}

// 根据ID查找二维码
func QRReadById(id int) (*QRCode, error){
  o := orm.NewOrm()
  a := new(QRCode)
  o.QueryTable("qrcode").Filter("id", id).One(a)
  if a.Id == 0 {
    return a, errors.New("没有该数据")
  }
  return a, nil
}

// 根据Name或者Desc的内容查找匹配的二维码（查找功能）
func QRSearch(content string) (qrlist []QRCode){
	o := orm.NewOrm()
	o.QueryTable("qrcode").OrderBy("id").Filter("name__contains", content).Filter("markdown__contains", content).All(&qrlist)
	beego.Debug(qrlist)
	return
}

// 自定义表名
func (u *QRCode) TableName() string {
    return "qrcode"
}

// 注册表
func init() {
	orm.RegisterModel(new(QRCode))
}
