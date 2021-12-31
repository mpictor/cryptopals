package enctype

type EncType int

const (
	Unknown EncType = iota
	Ecb
	Cbc
)

func (e EncType) String() string {
	switch e {
	case Ecb:
		return "ECB"
	case Cbc:
		return "CBC"
	case Unknown:
		return "unknown"
	default:
		return "(undefined)"
	}
}

//looks for repetition. if found, assume ecb. block len must be 16.
func Detect_ecb_cbc_16(data []byte) EncType {
	const blockLen = 16
	type block [blockLen]byte
	blocks := make(map[block]interface{})
	for len(data) > 0 {
		var blk block
		copy(blk[:], data[:blockLen])
		data = data[blockLen:]
		_, ok := blocks[blk]
		if ok {
			return Ecb
		}
		blocks[blk] = nil
	}
	return Cbc
}
