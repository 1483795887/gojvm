package classfile

//ConstantInfo 表示常量信息
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

//各种常量
const (
	CONSTANTUtf8               = 1
	CONSTANTInteger            = 3
	CONSTANTFloat              = 4
	CONSTANTLong               = 5
	CONSTANTDouble             = 6
	CONSTANTClass              = 7
	CONSTANTString             = 8
	CONSTANTFieldref           = 9
	CONSTANTMethodref          = 10
	CONSTANTInterfaceMethodref = 11
	CONSTANTNameAndType        = 12
	CONSTANTMethodHandle       = 15
	CONSTANTMethodType         = 16
	CONSTANTInvokeDynamic      = 18
)

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANTInteger:
		return &ConstantIntegerInfo{}
	case CONSTANTLong:
		return &ConstantLongInfo{}
	case CONSTANTFloat:
		return &ConstantFloatInfo{}
	case CONSTANTDouble:
		return &ConstantDoubleInfo{}
	case CONSTANTUtf8:
		return &ConstantUtf8Info{}
	case CONSTANTString:
		return &ConstantStringInfo{cp: cp}
	case CONSTANTClass:
		return &ConstantClassInfo{cp: cp}
	case CONSTANTFieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANTMethodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANTInterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANTNameAndType:
		return &ConstantNameAndTypeInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
