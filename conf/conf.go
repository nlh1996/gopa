package conf

const (
	// MgoURL 数据库连接地址
	MgoURL = "mongodb://localhost:27017"
	// CheckURL 爬取目标网站的地址用于代理ip测速，增加爬取成功率
	CheckURL = "http://httpbin.org/get"
)