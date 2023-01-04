package controllers

import (
	"fmt"
	"net/http"
)

// pagescontroller 处理静态页面
type PagesController struct {
}

// home 首页
func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>hello, 欢迎来到 goblog !</h1>")
}

// home 首页
func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

// home 首页
func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>请求页面未找到 :(</h1><p>如有疑惑，请联系我们。</h1>")
}
