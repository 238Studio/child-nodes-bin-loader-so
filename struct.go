package so

import "plugin"

// SoPackage so包结构
type SoPackage struct {
	name                 string              //包名。全局唯一。
	id                   int                 //id
	functions            []string            //支持的函数名称
	functionsArgTypes    map[string][]string //函数入参类型 函数名-入参类型表
	functionsReturnTypes map[string][]string //函数返回值类型 函数名-返回值类型表
	info                 map[string]string   //其他信息
	so                   *plugin.Plugin      //so对象
}

// SoLoader so加载器
type SoLoader struct {
	Sos       map[string]map[int]*SoPackage //已加载的so包集合。name->id->so
	soCounter map[string]int                //so计数器。每个name分配独立id。自增，单增。
}
