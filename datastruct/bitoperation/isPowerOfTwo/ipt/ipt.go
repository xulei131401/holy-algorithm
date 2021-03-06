package main

/**
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
0不是2的幂次方
2的1次-->2
2的2次-->2*2=4
2的3次-->2*2*2=8
2的4次-->2*2*2*2=16
2的5次-->2*2*2*2*2=32
结论：可以发现这些数，都是最高位为1，其他位为0
*/
func main() {

}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
