package model

// 默认返回体
type DefaultRes struct {
	Errinfo *Errinfo `protobuf:"bytes,1,opt,name=errinfo,proto3" json:"errinfo"`
}

// 返回错误题
type Errinfo struct {
	Code int    `protobuf:"varint,1,opt,name=Code,proto3" json:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"msg"`
}
