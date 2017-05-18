package models

import (
	"errors"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Column struct {
	Id      int   `orm:"pk;auto"`
	Name    string `form:"name"`
	Link    string `form:"link"`
	Pri     int   `form:"pri"`
	Deep    int
	Father  int   `form:"father"`
	Type    string `form:"type"`
	Content string `orm:"type(text)"form:"content"`
}

type MainMenu struct {
	Id       int
	Child    int
	Menu     *Column
	Menulist []*Column
}

// 增加一个目录
func CAdd(c *Column) error {
	o := orm.NewOrm()

	columnlist := make([]*Column, 0)
	CReadAll(&columnlist)
	c.Deep = DeepProcess(c, columnlist, 1)
	_, err := o.Insert(c)
	if err != nil {
		return errors.New("插入目录失败")
	}

	// 第二次插入，根据其Id修改它的Link属性
	o.Read(c)
	id :=strconv.Itoa(c.Id) 
	c.Link = "/column?Id=" + id
	_, err = o.Update(c)
	if err != nil {
		return errors.New("修改link失败")
	}
	return nil
}

// 计算新加入项目的深度
func DeepProcess(c *Column, clist []*Column, deep int) int {
	if c.Father != 0 {
		for i := 0; i < len(clist); i++ {
			if c.Father == clist[i].Id {
				deep = deep + 1
				ctemp := clist[i]
				clist[i].Id = 0
				for j := 0; j < len(clist); j++ {
					//beego.Debug("   ", tlist[j])
				}
				deep = DeepProcess(ctemp, clist, deep)
			}
		}
	}
	return deep
}

// 如何能做到在排序的时候对优先级大的项目优先排序
func SortColumn(id int, slist []*Column, dlist []*Column) ([]*Column, []*Column) {

	//	beego.Debug("=========开始==========")
	//	beego.Debug("这次寻找FatherId=", id, "的项目")
	var templist []*Column
	for i := 0; i < len(slist); i++ {
		if slist[i].Father == id {
			templist = append(templist, slist[i])
		}
	}
	// 第二步
	// 沉底法
	// 	beego.Debug("沉底排序开始")
	// 沉底排序有问题？
	for i := 0; i < (len(templist) - 1); i++ {
		for j := 0; j < (len(templist) - 1); j++ {
			if templist[j].Pri < templist[j+1].Pri {
				temp := templist[j]
				templist[j] = templist[j+1]
				templist[j+1] = temp
			}
		}
	}
	//	beego.Debug("沉底排序结束")

	for i := 0; i < len(templist); i++ {
		//		beego.Debug(templist[i])
	}
	for i := 0; i < len(templist); i++ {
		dlist = append(dlist, templist[i])
		dlist, slist = SortColumn(templist[i].Id, slist, dlist)
	}
	//	beego.Debug("=========结束==========")
	return dlist, slist
}

// 输入一个节点，返回它的所有子节点
func CFindSon(id int, slist []*Column, dlist []*Column) ([]*Column, []*Column) {
	// 找出所有子节点
	var templist []*Column
	for i := 0; i < len(slist); i++ {
		if slist[i].Father == id {
			templist = append(templist, slist[i])
		}
	}
	// 对子节点进行排序
	for i := 0; i < (len(templist) - 1); i++ {
		for j := 0; j < (len(templist) - 1); j++ {
			if templist[j].Pri < templist[j+1].Pri {
				temp := templist[j]
				templist[j] = templist[j+1]
				templist[j+1] = temp
			}
		}
	}
	// 返回这些子节点
	return templist, slist
}

// 输入一个节点，依次返回他的父系节点
// 输入的应该是一个Father
func CFindFather(id int, slist []*Column, dlist []*Column) []*Column {
	o := orm.NewOrm()
	u := new(Column)
	o.QueryTable("tax").Filter("id", id).One(u)
	if u.Id != 0 {
		for i := 0; i < len(slist); i++ {
			if slist[i].Id == id {
				dlist = append(dlist, slist[i])
				dlist = CFindFather(slist[i].Father, slist, dlist)
			}
		}
	}
	for i := 0; i < len(dlist); i++ {
		//		beego.Debug(dlist[i])
	}
	return dlist
}

// 输入一个父节点，依次返回它的所有一级子节点
func CFindBrother(id int, slist []*Column) []*Column {
	var templist []*Column
	for i := 0; i < len(slist); i++ {
		if slist[i].Father == id {
			templist = append(templist, slist[i])
		}
	}
	return templist
}

// 查看所有目录
func CReadAll(clist *[]*Column) {
	o := orm.NewOrm()
	o.QueryTable("tax").All(clist)
}

// 根据父类Id获得所有子类目录
func CReadByFather(father_id int) []*Column {
	o := orm.NewOrm()
	var clist []*Column
	o.QueryTable("tax").Filter("father", father_id).All(&clist)
	for i := 0; i < len(clist); i++ {
		//beego.Debug(tlist[i])
	}
	return clist
}

// 根据Id获得Tax的Content
func CReadContentById(id int) string {
	c := CReadById(id)
	return c.Content
}

// 根据Id获得一个Tax
func CReadById(id int) *Column {
	o := orm.NewOrm()

	var c Column
	o.QueryTable("tax").Filter("id", id).One(&c)

	// 可能会出现id不在数据库中的错误
	return &c
}

// 根据结构体修改
func CModify(c *Column) error {
	o := orm.NewOrm()

	ctemp := Column{Id: c.Id}
	if o.Read(&ctemp) == nil {
		ctemp = *c
		//beego.Debug(ttemp)
		if _, err := o.Update(&ctemp); err != nil {
			beego.Debug(err)
			errors.New("修改目录失败")
		} else {
			return nil
		}
		beego.Debug(ctemp)
	}

	return errors.New("修改目录失败")
}

// 删除一个目录
func CDel(id int) error {
	o := orm.NewOrm()

	if _, err := o.Delete(&Column{Id: id}); err != nil {
		return errors.New("删除目录失败")
	}

	return nil
}

// 自定义表名
func (u *Column) TableName() string {
    return "tax"
}

func init() {
	orm.RegisterModel(new(Column))
}