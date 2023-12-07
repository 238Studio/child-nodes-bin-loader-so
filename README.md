# child-nodes-bin-loader-so

## 说明
本包为so加载器，用于在linux环境上加载so文件。

## 接口说明

### func InitSoLoader() *SoLoader
初始化so加载器，返回加载器实例。

---
### func (so *SoPackage) GetName() string
获取so包名。

---
### func (so *SoPackage) GetID() int
获取so包ID。

---
### func (so *SoPackage) GetFunctions() []string
获取so包中的函数列表。

---
### func (so *SoPackage) GetFunctionsArgsTypes(methodName string) ([]string, error)
获取so包中指定函数的参数类型列表。传入函数名，返回参数类型列表和错误信息。

---
### func (so *SoPackage) GetFunctionReturnTypes(methodName string) ([]string, error)
获取so包中指定函数的返回值类型列表。传入函数名，返回返回值类型列表和错误信息。

---
### func (so *SoPackage) GetInfo(key string) (string, error)
获取so包中其他信息。传入key，返回value和错误信息。

---
### func (so *SoPackage) Execute(method string, args []uintptr, re uintptr) (err error)
执行so包中的函数。传入函数名、参数、返回值，返回错误信息。
其中，参数和返回值的类型为uintptr，需要在调用时进行类型转换。返回时通过re参数返回。

---
### func (soLoader *SoLoader) LoadBinPackage(path string) (name string, id int, err error)
加载so包并注册。传入so包路径，返回so包名、so包ID和错误信息。

---
### func (soLoader *SoLoader) ReleasePackage(name string, id int) (err error)
释放so包。传入so包名和so包ID，返回错误信息。

---
### func (soLoader *SoLoader) GetBinPackage(name string, id int) (soPackage *SoPackage, err error)
获取so包。传入so包名和so包ID，返回so包实例和错误信息。