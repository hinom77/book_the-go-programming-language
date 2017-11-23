package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	var count uint64
	for i := 0; i < 64; i++ {
		count += uint64(pc[byte(x>>(uint(i)/8) & 1)])
	}

	return int(count)
}