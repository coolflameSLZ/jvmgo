package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func ReadMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {

	memberCount := reader.ReadUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = &MemberInfo{
			cp:              cp,
			accessFlags:     reader.ReadUint16(),
			nameIndex:       reader.ReadUint16(),
			descriptorIndex: reader.ReadUint16(),
			attributes:      ReadAttributes(reader, cp),
		}
	}
	return members
}

func (self *MemberInfo) Name() string {
	return self.cp.GetUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.GetUtf8(self.descriptorIndex)
}

func (m MemberInfo) AccessFlags() uint16 {
	return m.accessFlags
}
