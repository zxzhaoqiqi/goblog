package article

import (
	"github.com/zxzhaoqiqi/goblog/pkg/route"
	"github.com/zxzhaoqiqi/goblog/pkg/types"
)

type Article struct {
	ID    uint64
	Title string
	Body  string
}

func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", types.Uint64ToString(article.ID))
}
