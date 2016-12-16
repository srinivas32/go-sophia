package sophia

// FieldType type of key or value in a row
type FieldType int

const (
	FieldTypeUInt8 FieldType = iota
	FieldTypeUInt16
	FieldTypeUInt32
	FieldTypeUInt64
	FieldTypeUInt8Rev
	FieldTypeUInt16Rev
	FieldTypeUInt32Rev
	FieldTypeUInt64Rev
	FieldTypeString
)

var fieldTypeNames = map[FieldType]string{
	FieldTypeUInt8:     "u8",
	FieldTypeUInt16:    "u16",
	FieldTypeUInt32:    "u32",
	FieldTypeUInt64:    "u64",
	FieldTypeUInt8Rev:  "u8rev",
	FieldTypeUInt16Rev: "u16rev",
	FieldTypeUInt32Rev: "u32rev",
	FieldTypeUInt64Rev: "u64rev",
	FieldTypeString:    "string",
}

func (t *FieldType) String() string {
	return fieldTypeNames[*t]
}
