package main

import (
	"ahgilak/aut-vahed/captcha"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type Context struct {
	http.Client
	captcha.Solver
}

func NewContext() *Context {
	jar, _ := cookiejar.New(nil)
	return &Context{
		Client: http.Client{
			Jar: jar,
		},
		Solver: *captcha.NewSolver(20),
	}
}

func (ctx *Context) InitLogin() {
	login, err := ctx.PostForm(
		"https://portal.aut.ac.ir/aportal/loginRedirect.jsp",
		url.Values{
			"login": {"ورود+از+درگاه+قبلي+پورتال"},
		},
	)

	defer login.Body.Close()

	if err != nil {
		panic(err)
	}
}

func (ctx *Context) GetCaptcha() io.Reader {
	res, _ := ctx.Get("https://portal.aut.ac.ir/aportal/PassImageServlet")

	if res.StatusCode == 200 {
		return res.Body
	} else {
		body, _ := io.ReadAll(res.Body)
		panic(string(body))
	}
}

func (ctx *Context) Login(username, password string) error {
	passline := ctx.Solve(ctx.GetCaptcha)

	data, err := ctx.PostForm(
		"https://portal.aut.ac.ir/aportal/login.jsp",
		url.Values{
			"username": {username},
			"password": {password},
			"passline": {passline},
			"login":    {"ورود+به+پورتال"},
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
	fmt.Println(string(body))
	if strings.Contains(string(body), "حروف تصویر صحیح نمیباشد") {
		return errors.New("Worng Captcha")
	}

	return nil
}
