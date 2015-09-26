package reverse

func Reverse(str string) (res string) {
	res = ""
	for _, s := range str {
		res = string(s) + res
	}
	return res

}
