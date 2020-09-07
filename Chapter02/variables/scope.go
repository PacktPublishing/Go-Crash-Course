package main

var packageScope = 2

func main() {
	var functionScope = 7
	_ = functionScope
	{ // this is a block declaration
		var blockScope = 3
		_ = blockScope
	} // end of the block
}
