package fib

func Fibcci() func() int {
	a,b := 0,1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func PrintMsg() func() string  {
	return func() string{
		return "test"
	}
}