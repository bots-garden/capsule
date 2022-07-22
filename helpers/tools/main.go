package helpers

func PtrSizePairFromArray(s []uint64) (uint32, uint32) {
	ptr := uint32(s[0] >> 32)
	size := uint32(s[0])
	return ptr, size
}
