package back_end

import (
  "os"
  "io/ioutil"
  "image/png"
  // "path/filepath"
  "strings"
  "env/models"
  "env/controllers"
  "github.com/astaxie/beego"
  //"github.com/astaxie/beego/utils/pagination"
)

type PictureController struct {
  controllers.BaseController
}

// 查看（Get）功能
func (c *PictureController) Get() {
  // postlist := make([]*models.Post, 0)
  // models.PReadAll(&postlist)

  // postsPerPage := 15
  // paginator := pagination.SetPaginator(c.Ctx, postsPerPage, models.PCountPosts())

  // c.Data["URL"] = beego.AppConfig.String("WEB_URL")
  // c.Data["QRList"] = models.ListPostsByOffsetAndLimit(paginator.Offset(), postsPerPage)
  // beego.Debug(models.ListPostsByOffsetAndLimit(paginator.Offset(), postsPerPage))
  // c.TplName = "back_end/qrcode.html"

  // 查找指定目录下的所有图片以供显示
  // 显示内容包括名称，引用路径和大小
  // 以及一个预览按钮，一个删除按钮
  // 一个查找框
  ////////////////////////////////////////////////////////
  // var pic_list []String
  // pwd, _ := os.Getwd()
  // beego.Debug(pwd)
  // dir, err := os.OpenFile(pwd, os.O_RDONLY, os.ModeDir)
  // if err != nil {
  //   defer dir.Close()
  //   // 打debug
  //   return
  // }
  // //fileinfo, _ := dir.Stat()
  // names, _ := dir.Readdir(-1)
  // for _, name := range names{
  //   if !name.IsDir() {
  //     beego.Debug(name.Name())
  //   }
  // }
  beego.Debug("图片")
  ///////////////////////////////////////////////////////
  // files := make([]string, 0)
  files := make([]models.Picture, 0)
  dirPath, _ := os.Getwd()
  dirPath = dirPath + "/static/picture"
  // 获得所有文件
  if dir, err := ioutil.ReadDir(dirPath); err == nil{
    PthSep := string(os.PathSeparator)
    suffix := strings.ToUpper(".png")
    for _, fi := range dir {
        if fi.IsDir() {
        continue
      }
      if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
        // files = append(files, dirPath+PthSep+fi.Name())
        beego.Debug(fi)
        var file models.Picture
        file.Name = fi.Name()
        file.Path = "static/picture" + PthSep + fi.Name()
        pngfile, err := os.Open(file.Path)
        if err!=nil {
          beego.Debug(err)
          continue
        }else{
          img, _ := png.DecodeConfig(pngfile)
          file.Height = img.Height
          file.Width = img.Width
        }
        // files = append(files, "static/picture"+PthSep+fi.Name())
        files = append(files, file)
      }
    }
  }
  // 将所有文件都显示出来
  // pngs := make([]models.Picture, 0)
  // for _, file := range files {
  //     png := make(models.Picture)
  //     png.Name = file
  //     png.Path = 
  //     pngs = append(files, )
  // }
  //////////////////////////////////////////////////////
  beego.Debug(files)
  // c.Ctx.Output.Download(files[0])
  // pngfile, err := os.Open(files[0])
  // if err != nil {
  //   beego.Debug(err)
  // }
  // img, err := png.DecodeConfig(pngfile)
  // beego.Debug(img)
  //////////////////////////////////////////////////////
  c.Data["Picture"] = files
  c.TplName = "back_end/public.html"
  c.Data["Tpl"] = "picture"
}

// 修改（Post）功能
func (c *PictureController) Post() {
  // var content = c.GetString("content")
  var post models.Post
  if err := c.ParseForm(&post); err !=nil {
    beego.Debug(err)
    // 说明有错误，跳转到查看所有项目界面？
    // 使用flash提示
    c.Redirect("/picture", 302)
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
  c.Redirect("/picture", 302)
}

// 增加（Add）功能
func (c *PictureController) Add() {
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
    c.Data["Tpl"] = "picture_add"
}

// 删除（Del）功能
func (c *PictureController) Del() {
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
  c.Redirect("/picture", 302)
}

// 查找一篇文章
func (c *PictureController) Search() {
  // content := c.GetString("content")
  // beego.Debug(content)
  // qrlist := models.QRSearch(content)
  // beego.Debug(qrlist)
  // c.Data["QRList"] = qrlist
  // // c.TplName = "back_end/qrcode.html"
  c.TplName = "back_end/public.html"
  c.Data["Tpl"] = "picture_add"
}

func ListAllPicture() {

}