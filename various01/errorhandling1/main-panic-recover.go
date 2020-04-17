package main

func recoverFunc() {

	r := recover()

	if r != nil {
		fmt.Println("revovered from ", r)
		fmt.Println("\n\nStacktrace follows:")
		debug.PrintStack()
	}

}

func somefunction(a, b int) {
	defer recoverFunc()

	if a == 0 {

		panic("runtime error: a cannot be 0")

	}
}

func someotherfunction() int {
	return 0
}

func main() {

        // this has recover - so it will allow main to continue to the next step
	somefunction(0, 2)

        // this point onwards will still work even if somefunction() panics - because it has recovery
	someotherfunction()
   
}
