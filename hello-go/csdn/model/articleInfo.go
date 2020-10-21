package model

type CArticleInfo struct {
	url      string
	nickName string
	// 是否原创
	original bool
	// 分类专栏
	category string
	// 文章标签
	tagList string
	//阅读量
	readCount int

	// 点赞数
	like int
	// 首次发布
	createTime string
	// 最后发布
	updateTime string
}
