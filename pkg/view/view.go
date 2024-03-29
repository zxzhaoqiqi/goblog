package view

import (
	"html/template"
	"io"
	"path/filepath"
	"strings"

	"github.com/zxzhaoqiqi/goblog/pkg/logger"
	"github.com/zxzhaoqiqi/goblog/pkg/route"
)

// render 渲染视图
func Render(w io.Writer, data interface{}, tplFiles ...string) {
	// 1 设置模板相对路径
	viewDir := "resources/views/"

	// 2. 遍历传参文件列表 Slice, 设置正确的路径，支持 dir.filename 语法糖
	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	// 3. 所有布局模板文件 Slice
	files, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.Error(err)

	// 4. 在 Slice 中新增目标文件
	allFiles := append(files, tplFiles...)

	// 5. 解析所有模板文件
	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"RouteName2Url": route.Name2URL,
		}).ParseFiles(allFiles...)
	logger.Error(err)

	// 6. 渲染模板
	err = tmpl.ExecuteTemplate(w, "app", data)
	logger.Error(err)
}
