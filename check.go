package main

import (
	"database/sql"
	"github.com/astaxie/beego/toolbox"
)

// 健康检查文件，记得在conf/app.conf中开启健康检查

type DataBaseCheck struct {
}

// 链接过程中有错误的发生，就返回错误
func (*DataBaseCheck) Check() error {
	// Check方法是toolbox中自带的
	_, err := sql.Open("mysql", "rich:www.123.com@tcp/blog?charset=utf8mb4")
	if err != nil {
		return err
	}

	return nil
}

func init() {
	// 把DataBaseCheck检查项添加到toolbox的健康检查中
	toolbox.AddHealthCheck("database", &DataBaseCheck{})
}
