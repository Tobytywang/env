package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type SizeOfPicture struct {
  Height int 
  Width  int
}

// 定义栏目（数据库表）
type Picture struct {
	Name    string `form:"name"`
	Path    string
	Link    string
	SizeOfPicture
}