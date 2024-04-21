package main

func main() {
	// defining multi const all in one:
	const (
		firstConstNormal  = 1
		secondConstNormal = 2
		thirdConstNormal  = 3
	)

	// if you have a incremental model const that starts with 1:
	const (
		firstConst   = iota + 1        // ----> 1
		secondConst                    // ----> 2
		thirdConst                     // ----> 3
		fourthConst  = thirdConst + 10 // ----> 13
		fifthConst                     // ----> 13 because we didn't use iota before and if we didn't assign any it will return last value
		sixthConst   = iota + 20       // ----> 25 it returns 5 because last step iota goes to 4 and now it is fifth iteration of iota
		seventhConst                   // ----> 26
		eighthConst                    // ----> 27
	)
}
