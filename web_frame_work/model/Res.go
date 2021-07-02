package model

// 默认返回体
type DefaultRes struct {
	Errinfo *Errinfo
}

// 返回错误题
type Errinfo struct {
	Code int
	Msg  string
}
