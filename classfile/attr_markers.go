package classfile

//DeprecatedAttribute @Deprecate
type DeprecatedAttribute struct {MarkerAttribute}
//SyntheticAttribute @Synthetic
type SyntheticAttribute struct{MarkerAttribute}
//MarkerAttribute 标记属性
type MarkerAttribute struct{}

func (client *MarkerAttribute) readInfo(reader *ClassReader){
	//什么都不敢  
}