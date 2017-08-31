package models

import (
	"os"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	qrcode "github.com/skip2/go-qrcode"
)

type QRCode struct {
	Id       int    `orm:"pk;auto"form:"id"`
	Name     string `orm:"size(100)"form:"name"`
	Link     string // 页面的路径
	Pic      string // 图片存储的路径
	Desc     string `orm:"type(text)"form:"desc-html"`
	Markdown string `orm:"type(text)"form:"desc-markdown"`
	Read     uint
}

type QRCodeExt struct {
	QRCode
	PicExist bool
}

var (
	QR_PATH string = beego.AppConfig.String("QR_PATH")
	WEB_URL string = beego.AppConfig.String("WEB_URL")
)

func QRAddOne(code *QRCode) error {
	o := orm.NewOrm()

	_, err := o.Insert(code)
	if err != nil {
		return err
	}

	// update the Link
	err = o.Read(code)
	if err != nil {
		return err
	}
	code.Link = "/plant?id=" + strconv.Itoa(code.Id)
	_, err = o.Update(code)
	if err != nil {
		return err
	}
	return nil
}

func QRUpdate(code *QRCode) error {
	o := orm.NewOrm()
	temp := QRCode{Id: code.Id}
	err := o.Read(&temp)
	if err != nil {
		return err
	}
	temp.Name = code.Name
	temp.Pic = code.Pic
	temp.Desc = code.Desc
	temp.Markdown = code.Markdown
	_, err = o.Update(&temp, "Desc", "Markdown", "Name", "Pic")
	if err != nil {
		return err
	}
	return nil
}

func QRRead(id int) error {
	o := orm.NewOrm()
	temp := QRCode{Id: id}
	err := o.Read(&temp)
	if err != nil {
		return err
	}
	temp.Read = temp.Read + 1
	_, err = o.Update(&temp, "Read")
	if err != nil {
		return err
	}
	return nil
}

func QRDel(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&QRCode{Id: id})
	if err != nil {
		return err
	}
	return nil
}

func QRReadAll(qrlist *[]*QRCode) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("qrcode").OrderBy("-id").All(qrlist)
	if err != nil {
		return err
	}
	return nil
}

func QRReadById(id int) (*QRCode, error) {
	o := orm.NewOrm()
	a := new(QRCode)
	err := o.QueryTable("qrcode").Filter("id", id).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func CountCodes() int64 {
	o := orm.NewOrm()
	cnt, _ := o.QueryTable("qrcode").Count()
	return cnt
}

func ListCodesByOffsetAndLimit(sort string, offset int, codeperpage int) (destlist []QRCodeExt) {
	o := orm.NewOrm()
	var qrlist []QRCode
	var templist []QRCode
	var top int

	if sort == "readup" {
		o.QueryTable("qrcode").OrderBy("read").All(&templist)
	} else if sort == "readdown" {
		o.QueryTable("qrcode").OrderBy("-read").All(&templist)
	} else if sort == "iddown" {
		o.QueryTable("qrcode").OrderBy("-id").All(&templist)
	} else {
		o.QueryTable("qrcode").OrderBy(sort).All(&templist)
	}

	if (offset + codeperpage) > len(templist) {
		top = len(templist)
	} else {
		top = (offset + codeperpage)
	}
	qrlist = templist[offset:top]

	var qrcodeext QRCodeExt
	for _, qr := range qrlist {
		qrcodeext.QRCode = qr
		// if os.IsExist(qr.Pic) {
		if _, err := os.Stat(qr.Pic); os.IsNotExist(err) {
			qrcodeext.PicExist = false
		} else {
			qrcodeext.PicExist = true
		}
		destlist = append(destlist, qrcodeext)
	}
	return destlist
}

// 重点改造
func QRSearch(content string) (destlist []QRCodeExt) {
	o := orm.NewOrm()
	cond := orm.NewCondition()
	cond1 := cond.Or("name__contains", content).Or("markdown__contains", content)
	var qrlist []QRCode
	o.QueryTable("qrcode").SetCond(cond1).All(&qrlist)

	var qrcodeext QRCodeExt
	for _, qr := range qrlist {
		qrcodeext.QRCode = qr
		if _, err := os.Stat(qr.Pic); os.IsNotExist(err) {
			qrcodeext.PicExist = false
		} else {
			qrcodeext.PicExist = true
		}
		destlist = append(destlist, qrcodeext)
	}
	return destlist
}

func Create_qrcode(code *QRCode, name string) (path string, err error) {
	os.Mkdir(QR_PATH, 0777)
	link_string := "http://" + WEB_URL + code.Link
	code_string := "static/" + QR_PATH + "/" + strconv.Itoa(code.Id) + "-" + name + ".png"
	err = qrcode.WriteFile(link_string, qrcode.Medium, 256, code_string)
	if err != nil {
		return "", err
	}
	return code_string, nil
}

// 自定义表名
func (u *QRCode) TableName() string {
	return "qrcode"
}

// 注册表
func init() {
	orm.RegisterModel(new(QRCode))
}
