package classfile

//MemberInfo 表示class中成员
type MemberInfo struct {
	cp             ConstantPool
	accessFlags    uint16
	nameIndex      uint16
	descritorIndex uint16
	attributes     []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:             cp,
		accessFlags:    reader.readUint16(),
		nameIndex:      reader.readUint16(),
		descritorIndex: reader.readUint16(),
		attributes:     readAttributes(reader, cp),
	}
}

//Name 返回成员名字
func (client *MemberInfo) Name() string{
	return client.cp.getUtf8(client.nameIndex)
}
