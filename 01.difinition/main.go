package _1_difinition

// GenericInterface 定义泛型
type GenericInterface interface {
	//int 代表int类型
	//~int 代表int类型及派生类型
}

// GenericStruct 定义泛型结构体
type GenericStruct[G GenericInterface] struct {
}

// GenericFunction 定义泛型函数
type GenericFunction[G GenericInterface] func(G)

// NewGenericStruct 创建泛型结构体
func NewGenericStruct[G GenericInterface]() GenericStruct[G] {
	return GenericStruct[G]{}
}

// GenericStructFunctionImpl 结构体的泛型函数
func (s *GenericStruct[G]) GenericStructFunctionImpl(g G) {
}
