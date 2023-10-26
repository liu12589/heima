package models

type LoginMessage struct {
	User     string
	Password string
}

// Login 客户信息
type Login struct {
	Id       int64  `gorm:"column:id;not null;type:int(11) primary key auto_increment"`
	User     string `gorm:"column:user"`
	Password string `gorm:"column:password"`
}

type Project struct {
	Id          int64  `gorm:"column:id;not null;type:int(11) primary key auto_increment"`
	ProjectId   string `gorm:"project_id"`
	Name        string `gorm:"column:user"`
	ProjectName string `gorm:"project_name"`
	Model       string `gorm:"column:model"`
	Url         string `gorm:"column:url"`
	PromptId    int    `gorm:"column:prompt_id"`
}

type ProjectReq struct {
	Id  string
	Url string
}

type ProjectRes struct {
	User         string `json:"user"`
	ProjectName  string `json:"project_name"`
	ProjectModel string `json:"project_model"`
}

type AllAccount struct {
	User string `json:"user"`
}

type DbInformation struct {
	User        string `gorm:"column:id"`
	MILVUS_HOSt string `gorm:"milvus_host"`
	MILVUS_PORT string `gorm:"milvus_port"`
}

type UserProjectRelation struct {
	Id        string `json:"id"`
	UserName  string `json:"user_name"`
	ProjectId string `json:"project_id"`
}
