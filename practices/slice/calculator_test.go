package practices

import "fmt"

func ExampleEval() {
	fmt.Println(Eval("5"))
	fmt.Println(Eval("5 + 1"))
	// Output:
	// 5
	// 6
}
