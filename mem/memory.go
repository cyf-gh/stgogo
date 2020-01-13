package st_memory

func Memset(a []byte, v byte ) {
	for i := range a {
		a[i] = v
	}
}

func CutAdapt( big []byte, small []byte ) []byte {
	for i := range small {
		small[i] = big[i]
	}
	return small
}