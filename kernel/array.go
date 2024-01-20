package kernel

type TArray struct {
	Data  uintptr
	Count int32
	Max   int32
}

func (arr *TArray) ReplaceData(newData uintptr) {
	arr.Data = newData
}

func (arr *TArray) ReadAtIndex(index int, d *Driver) uintptr {
	return d.Readvm(arr.Data+(uintptr(index)*8), 8)
}
