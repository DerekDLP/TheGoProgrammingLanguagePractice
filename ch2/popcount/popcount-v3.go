package popcount

func PopCountV3(x uint64) int  {
	// 使用循环替代单一表达式
	var res int
	var i uint64
	for ; i < uint64(len(pc)); i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return res
}