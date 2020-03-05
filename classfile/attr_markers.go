package classfile

// 废弃的类
type DeprecatedAttribute struct {
	MarkerAttribute
}

// 编译器生成的类
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {
}

/**
由于这两个属性都没有数据，所以readInfo（）方法是空的。
*/
func (m *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
