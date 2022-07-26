package helpers

/*
Get the pointer position and a size from result of ExportedFunction("F").Call().
ex: the string pointer position (in memory) and the length of the string
*/
func GetPackedPtrPositionAndSize(result []uint64) (ptrPos uint32, size uint32) {
	return uint32(result[0] >> 32), uint32(result[0])
}
