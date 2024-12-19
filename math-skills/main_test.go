package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd1 := exec.Command("go", "build", "main.go")
	build, err := cmd1.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(build)
	for i := 0; i < 10000; i++ {
		cmd := exec.Command("bin/math-skills")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		cmd1 := exec.Command("./main")
		output1, err := cmd1.CombinedOutput()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if string(output1) != string(output) {
			fmt.Println("Output does not match")
			fmt.Printf("o1:%s\ngo o2 :%s", output, output1)
			return
		}
		fmt.Printf("Test %d passed\n", i)
	}
	// Print the command's output
	fmt.Print("good\n")
}
