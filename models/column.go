package models

import (
	"errors"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 定义栏目
type Column struct {
	Id      int    `orm:"pk;auto"`
	Name    string `form:"name"`
	Link    string `form:"link"`
	Pri     int    `form:"pri"`
	Deep    int
	Father  int    `form:"father"`
	Type    string `form:"type"`
	Content string `orm:"type(text)"form:"content"`
}

type ColumnExt struct {
	Column
	HasSon bool
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

	// update the deep
	columnlist := make([]*Column, 0)
	CReadAll(&columnlist)
	c.Deep = deepProcess(c, columnlist, 1)

	// 插入数据库
	_, err := o.Insert(c)
	if err != nil {
		return err
	}

	err = o.Read(c)
	if err != nil {
		return err
	}

	id := strconv.Itoa(c.Id)
	c.Link = "/column?Id=" + id
	_, err = o.Update(c)
	if err != nil {
		return err
	}
	return nil
}

// 计算新加入项目的深度
func deepProcess(c *Column, clist []*Column, deep int) int {
	if c.Father != 0 {
		for i := 0; i < len(clist); i++ {
			if c.Father == clist[i].Id {
				deep = deep + 1
				ctemp := clist[i]
				clist[i].Id = 0
				deep = deepProcess(ctemp, clist, deep)
			}
		}
	}
	return deep
}

// 如何能做到在排序的时候对优先级大的项目优先排序
func SortColumn(id int, slist []*Column, dlist []*Column) ([]*Column, []*Column) {

	var templist []*Column
	for i := 0; i < len(slist); i++ {
		if slist[i].Father == id {
			templist = append(templist, slist[i])
		}
	}

	for i := 0; i < (len(templist) - 1); i++ {
		for j := 0; j < (len(templist) - 1); j++ {
			if templist[j].Pri < templist[j+1].Pri {
				temp := templist[j]
				templist[j] = templist[j+1]
				templist[j+1] = temp
			}
		}
	}

	for i := 0; i < len(templist); i++ {
		dlist = append(dlist, templist[i])
		dlist, slist = SortColumn(templist[i].Id, slist, dlist)
	}

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

func CReadAllColumn(clist *[]*Column) {
	o := orm.NewOrm()
	o.QueryTable("tax").Filter("type", "column").Filter("deep__lte", 2).All(clist)
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
	c, _ := CReadById(id)

	return c.Content
}

// 根据Id获得一个Tax
func CReadById(id int) (*Column, error) {
	o := orm.NewOrm()
	// c := new(Column)
	var c Column

	o.QueryTable("tax").Filter("id", id).One(&c)
	if c.Id == 0 {
		return &c, errors.New("没有该数据")
	}

	// 可能会出现id不在数据库中的错误
	return &c, nil
}

// 作为一个过滤函数，数据增强结构体
// 帮助实现Column里的功能
func CHasSon(columnlist []*Column) (destlist []ColumnExt) {
	var columnext ColumnExt
	for _, column := range columnlist {
		columnext.Column = *column

		columnlist := make([]*Column, 0)
		CReadAll(&columnlist)
		dlist := make([]*Column, 0)
		dlist, columnlist = CFindSon(column.Id, columnlist, dlist)

		if len(dlist) > 0 {
			columnext.HasSon = true
		} else {
			columnext.HasSon = false
		}
		destlist = append(destlist, columnext)
	}

	return destlist
}

// 根据结构体修改
func CModify(c *Column) (err error) {
	o := orm.NewOrm()
	beego.Debug(c)
	beego.Debug(c.Id)
	ctemp := Column{Id: c.Id}
	if err = o.Read(&ctemp); err == nil {
		ctemp = *c
		// 修改优先级
		columnlist := make([]*Column, 0)
		CReadAll(&columnlist)
		ctemp.Deep = deepProcess(c, columnlist, 1)
		//beego.Debug(ttemp)
		_, err := o.Update(&ctemp)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return err
}

// 删除一个目录
func CDel(id int) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Column{Id: id})
	if err != nil {
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
