package models



type Project struct {
	name   string
	criteria	Criteria
}

//docAddress表示项目文件地址
//proAddress表示项目地址
type CreateProjectResult struct {
	DocAddress interface{}
	ProAddress interface{}
}