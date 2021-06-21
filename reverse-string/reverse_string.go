package reverse


func Reverse(s string) (rev string) {
	for _, c := range s {
		rev = string(c) + rev
	}
	return 
}
