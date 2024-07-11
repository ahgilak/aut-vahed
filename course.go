package main

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"
)

func (ctx *Context) AddCourse(course string) error {
	ctx.Get("https://portal.aut.ac.ir/aportal/regadm/student.portal/student.portal.jsp?action=edit&st_info=register&st_sub_info=0")

	passline := ctx.Solve(ctx.GetCaptcha)
  fmt.Println(course)
	data, err := ctx.PostForm(
		"https://portal.aut.ac.ir/aportal/regadm/student.portal/student.portal.jsp?action=apply_reg&st_info=add",
		url.Values{
			"st_reg_course": {course},
			"addpassline":     {passline},
			"st_course_add":   {"درس را اضافه کن"},
		},
	)

	defer data.Body.Close()

	if data.StatusCode != 200 {
		fmt.Println(data.StatusCode)
		return errors.New(data.Status)
	}

	if err != nil {
		panic(err)
	}

	body, _ := io.ReadAll(data.Body)

	if strings.Contains(string(body), "فيلد حروف تصوير معتبر نميباشد يا به آن دسترسي نداريد!") {
		return errors.New("Wrong Captcha")
	}

  fmt.Println(string(body))

	return nil

}
