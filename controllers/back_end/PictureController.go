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
        file.Link = "http://" + beego.AppConfig.String("WEB_URL") + "/static/picture" + PthSep + fi.Name()
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
  beego.Debug(files)
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
  // 增加上传图片功能
  filepath := "static/picture"
  _, _, err := c.GetFile("pic")
  name := c.GetString("name")
  filetype := c.GetString("filetype")
  beego.Debug(err)
  if err == nil {
    os.MkdirAll(filepath, 0777)
    if err:=c.SaveToFile("pic", filepath+"/"+name + filetype); err!=nil{
      beego.Debug(err)
    }
  }
  c.TplName = "back_end/public.html"
  c.Data["Tpl"] = "picture"
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