package kernel

type TArrayDrink struct {
	Data  uintptr
	Count int32
	Max   int32
}

func (arr *TArrayDrink) ReplaceData(newData uintptr) {
	arr.Data = newData
}

func (arr *TArrayDrink) ReadAtIndex(index int, d *Driver) uintptr {
	return d.Readvm(arr.Data+(uintptr(index)*8), 8)
}
