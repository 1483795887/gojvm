package classfile

//ConstantPool 表示常量池
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

func (client ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := client[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (client ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := client.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := client.getUtf8(ntInfo.nameIndex)
	_type := client.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (client ConstantPool) getUtf8(index uint16) string {
	utf8Info := client.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

func (client ConstantPool) getClassName(index uint16) string {
	classInfo := client.getConstantInfo(index).(*ConstantClassInfo)
	return client.getUtf8(classInfo.nameIndex)
}
