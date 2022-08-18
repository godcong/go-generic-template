package _1_difinition

const (
	String = "string"
	Int    = "int"
	Float  = "float"
	Bool   = "bool"
	Byte   = "byte"
	Bytes  = "[]byte"
)

// GenericInterface 定义泛型
type GenericInterface interface {
	//int 代表int类型
	//~int 代表int类型及派生类型
	int | string | []byte
	Type() string
}

// GenericMap 定义泛型map
type GenericMap[G GenericInterface] map[string]G

// GenericStruct 定义泛型结构体
type GenericStruct[G GenericInterface] struct {
	//使用泛型map
	genericMap GenericMap[G]
}

// GenericFunction 定义泛型函数
type GenericFunction[G GenericInterface] func(G)

// NewGenericStruct 创建泛型结构体
func NewGenericStruct[G GenericInterface](v G) GenericStruct[G] {
	switch v.Type() {
	case String:
	case Int:
		//...
	}

	return GenericStruct[G]{}
}

// GenericStructFunctionImpl 结构体的泛型函数
func (s *GenericStruct[G]) GenericStructFunctionImpl(g G) {
}
