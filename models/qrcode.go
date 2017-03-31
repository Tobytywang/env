package models

import (
	"errors"
	"strconv"
	"os"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// qrcode
	qrcode "github.com/skip2/go-qrcode"
)

// 单页面（富文本处理）
// ID
// 学名
// 链接网站链接
// 图片链接
// 植物描述
// 阅读热度
type QRCode struct {
	Id      int   `orm:"pk;auto"form:"id"`
	Name    string `orm:"size(100)"form:"name"`
	Link		string // 页面的路径
	Pic     string // 图片存储的路径
	Code		string // 二维码存储的路径
	Desc    string `orm:"type(text)"form:"desc-html"`
	Markdown    string `orm:"type(text)"form:"desc-markdown"`
	Read    uint
}

var (
	QRPATH string = "static/public/"
)

// 增加一个二维码
func QRAddOne(code *QRCode) error{
  o := orm.NewOrm()

  // 第一次插入
  _, err := o.Insert(code)
  if err != nil {
		beego.Debug(err)
    return errors.New("添加植物失败(1)")
  }

  // 第二次插入，根据其Id修改它的Link属性
	// Link属性用来在扫码时通过id定位渲染对象
  o.Read(code)
  intid := (int)(code.Id)
  code.Link = "http://xxx.org/plant?Id=" + strconv.Itoa(intid)
  _, err = o.Update(code)
  if err != nil {
		beego.Debug(err)
    return errors.New("添加植物失败(2)")
  }

	name := strings.Split(code.Name, ".")
	// 更新二维码信息
	// 根据链接生成二维码并存储在public目录下
	os.Mkdir("static/public/", 0777)
	err = qrcode.WriteFile(code.Link, qrcode.Medium, 256, QRPATH + strconv.Itoa(code.Id) + "-" + name[0] + ".png")
	if err != nil {
		beego.Debug(err)
		return errors.New("生成二维码失败！")
	}
	// 将二维码的存储路径填充在Code字段
	o.Read(code)
  // intid := (int)(code.Id)
  // code.Link = "http://xxx.org/plant?Id=" + strconv.Itoa(intid)

	beego.Debug(name)
	code.Code = "static/public/" + strconv.Itoa(code.Id) + "-" + name[0] + ".png"
  _, err = o.Update(code)
	beego.Debug(code)
  if err != nil {
		beego.Debug(err)
    return errors.New("生成二维码路径失败！")
  }
  return nil
}

// 修改一个二维码
func QRUpdate(code *QRCode) error{
	// 将原先的id删掉
	// 重新add并调用原先的id
	o := orm.NewOrm()
	beego.Debug(code)
	temp := QRCode{Id: code.Id}
	if o.Read(&temp) == nil {
		temp = *code
		//beego.Debug(ttemp)
		// 将原先的删掉
		QRDel(code.Id)
		// 用原先的id新建一个
		if err := QRAddOne(&temp); err != nil {
			beego.Debug(err)
		} else {
			return nil
		}
		// beego.Debug(ttemp)
	}
	return errors.New("修改目录失败")
}

// 删除一个二维码
func QRDel(id int) error{
  o := orm.NewOrm()
	if _, err := o.Delete(&QRCode{Id: id}); err != nil {
		return errors.New("删除植物失败")
	}
	return nil
}

// 查找二维码
// 查看所有文章
func QRReadAll(qrlist *[]*QRCode) {
	o := orm.NewOrm()
	o.QueryTable("qrcode").OrderBy("-id").All(qrlist)
}

func CountCodes() int64{
	o := orm.NewOrm()
	cnt, _ := o.QueryTable("qrcode").Count()
  return cnt
}

// ListPostsByOffsetAndLimit
func ListCodesByOffsetAndLimit(offset int, codeperpage int) (qrlist []QRCode){
	o := orm.NewOrm()
	// templist := make([]QRCode, 0)
	var templist []QRCode
	o.QueryTable("qrcode").OrderBy("-id").All(&templist)
	var top int
	if ((offset+codeperpage)>len(templist)){
		top = len(templist)
	} else {
		top = (offset+codeperpage)
	}
	qrlist = templist[offset:top]
	return qrlist
}

// 根据ID查找
func QRReadById(id int) (*QRCode, error){
  o := orm.NewOrm()
  a := new(QRCode)
  o.QueryTable("qrcode").Filter("id", id).One(a)
  if a.Id == 0 {
    return a, errors.New("没有该数据")
  }
  return a, nil
}

func (u *QRCode) TableName() string {
    return "qrcode"
}

func init() {
	orm.RegisterModel(new(QRCode))
}
