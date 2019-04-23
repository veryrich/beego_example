package controllers

import (
	"encoding/json"
	"firstbee/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
)

type MainController struct {
	beego.Controller
}

// 获取题目信息
func (c *MainController) Get() {
	var subject models.Subject

	// 对用户传入的参数做校验
	// 将数据校验的部分统一封装到一个匿名函数中 这样的做法可以将校验相关的代码模块化处理，让代码的可读性和可维护性高一些
	err := func() error {

		id, err := c.GetInt("id") // 获取浏览器传递过来的id
		logs.Info(id)             // 日志记录
		if err != nil {
			// 如果id是空的,或者id的值不合法,默认返回给用户id为1的请求
			id = 1
		}
		subject, err = models.GetSubject(id) // 通过id获取题目
		if err != nil {
			// 如果有错误,返回subject不存在
			return errors.New("subject not exist")
		}
		return nil
	}()

	if err != nil {
		// 判断最终整个模板是否有错误的产生
		c.Ctx.WriteString("wrong params") // 如果有错误返回
	}

	// 正常逻辑开始，开始处理查询到的题目信息
	var option map[string]string // models表中option字段使用json格式存储，所以这里要把option做json编码

	//进行json decode，将选项转换成map的格式放到option中
	if err = json.Unmarshal([]byte(subject.Option), &option); err != nil {
		c.Ctx.WriteString("woring params, json decode") // 如果有错误写进上下文中
	}

	// 开始给用户数据
	c.Data["ID"] = subject.Id
	c.Data["Option"] = option
	c.Data["Img"] = "/static" + subject.Img
	c.TplName = "guess.html"

}

func (c *MainController) Post() {
	var subject models.Subject

	err := func() error {

		id, err := c.GetInt("id")
		logs.Info(id)
		if err != nil {

			id = 1
		}
		subject, err = models.GetSubject(id)
		if err != nil {
			return errors.New("subject not exist")
		}
		return nil
	}()

	if err != nil {
		c.Ctx.WriteString("wrong params")
	}

	// 获取用户传入的答案选项
	answer := c.GetString("key")
	right := models.Answer(subject.Id, answer)

	// 将数据写入c.Data
	c.Data["Right"] = right
	c.Data["Next"] = subject.Id + 1
	c.Data["ID"] = subject.Id
	c.TplName = "guess.html"
}
