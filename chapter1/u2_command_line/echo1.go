package u2_command_line

import "os"

func main() {
	// init variable numbers and sequence
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	println(s)
}
