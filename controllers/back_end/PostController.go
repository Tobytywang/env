package back_end

import (
  "env/models"
	"env/controllers"
	"github.com/astaxie/beego"
  "github.com/astaxie/beego/utils/pagination"
)

type PostController struct {
  controllers.BaseController
}

// Get方法查看所有的文章
func (c *PostController) Get() {
  postlist := make([]*models.Post, 0)
  models.PReadAll(&postlist)

  postsPerPage := 15
  paginator := pagination.SetPaginator(c.Ctx, postsPerPage, models.PCountPosts())

  c.Data["URL"] = beego.AppConfig.String("WEB_URL")
  c.Data["QRList"] = models.ListPostsByOffsetAndLimit(paginator.Offset(), postsPerPage)
  beego.Debug(models.ListPostsByOffsetAndLimit(paginator.Offset(), postsPerPage))
  // c.TplName = "back_end/qrcode.html"
  c.TplName = "back_end/public.html"
  c.Data["Tpl"] = "post"
}

func (c *PostController) Post() {
  // var content = c.GetString("content")
  var post models.Post
  if err := c.ParseForm(&post); err !=nil {
    beego.Debug(err)
    // 说明有错误，跳转到查看所有项目界面？
    // 使用flash提示
    c.Redirect("/post", 302)
  }
  if post.Id != 0 {
    // 有Id，表示修改
    if _, err := models.PReadById(post.Id); err == nil {
      if err := models.PModify(&post); err != nil {
        beego.Debug(err)
      }
  }
  } else {
    // 没有Id，表示新增
    if err := models.PAdd(&post); err != nil {
      beego.Debug(err)
    } 
  }
  c.Redirect("/post", 302)
}

// 增加一篇文章
func (c *PostController) Add() {
  // if id, err := c.GetInt("id"); err == nil{
  //   if code, err := models.QRReadById(id); err == nil{
  //     beego.Debug(code)
  //     c.Data["Modify"] = true
  //     c.Data["Code"] = code
  //   }
  // }
  // // c.TplName = "back_end/qrcode_add.html"
  ////////////////////////////////////////////////////
  // if(c.Ctx.Input.Is("POST") == true) {

    // c.TplName = "index.tpl"
  if id, err := c.GetInt("id"); err == nil{
    if post, err := models.PReadById(id); err == nil{
      beego.Debug(post)
      c.Data["Modify"] = true
      c.Data["Post"] = post
    }
  }
    c.TplName = "back_end/public.html"
    c.Data["Tpl"] = "post_add"
}

// 删除一篇文章
func (c *PostController) Del() {
  // id, err := c.GetInt("id");
  // if err != nil {
  //   beego.Debug(err)
  // }
  // err = models.QRDel(id);
  // if err != nil {
  //   beego.Debug(err)
  // }
  id, err := c.GetInt("id");
  if err != nil {
    beego.Debug(err)
  }
  err = models.PDel(id);
  if err != nil {
    beego.Debug(err)
  }
  c.Redirect("/post", 302)
}

// 查找一篇文章
func (c *PostController) Search() {
  // content := c.GetString("content")
  // beego.Debug(content)
  // qrlist := models.QRSearch(content)
  // beego.Debug(qrlist)
  // c.Data["QRList"] = qrlist
  // // c.TplName = "back_end/qrcode.html"
  c.TplName = "back_end/public.html"
  c.Data["Tpl"] = "post_add"
}