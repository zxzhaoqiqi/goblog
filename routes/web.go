package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zxzhaoqiqi/goblog/app/http/controllers"
	"github.com/zxzhaoqiqi/goblog/app/http/middlewares"
)

func RegisterWebRoutes(r *mux.Router) {

	pc := new(controllers.PagesController)

	// 设置 404 页面
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	// 静态页面
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")

	// 文章相关页面
	ac := new(controllers.ArticlesController)

	r.HandleFunc("/articles", ac.Index).Methods("GET").Name("articles.index")
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")

	r.HandleFunc("/articles/create", ac.Create).Methods("GET").Name("articles.create")
	r.HandleFunc("/articles", ac.Store).Methods("POST").Name("articles.store")

	r.HandleFunc("/articles/{id:[0-9]+}/edit", ac.Edit).Methods("GET").Name("articles.edit")
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Update).Methods("POST").Name("articles.update")

	r.HandleFunc("/articles/{id:[0-9]+}/delete", ac.Delete).Methods("POST").Name("articles.delete")

	// 中间件: 强制内容类型为 HTML
	r.Use(middlewares.ForceHTMLMiddle)

}
