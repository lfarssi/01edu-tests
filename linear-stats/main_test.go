package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	// .exe for compatibility with windows linux ignores file extension for executables
	cmd1 := exec.Command("go", "build", "-o", "test.exe", ".")
	build, err := cmd1.CombinedOutput()
	if err != nil {
		fmt.Println("Error0:", err)
		return
	}
	fmt.Println(build)
	for i := 0; i < 50; i++ {
		cmd := exec.Command("stat-bin/bin/linear-stats")
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error1:", err)
			return
		}
		cmd1 := exec.Command("./test.exe", "data.txt")
		output1, err := cmd1.CombinedOutput()
		if err != nil {
			fmt.Println("Error2:", err)
			return
		}
		if string(output1) != string(output) {
			fmt.Println("Output does not match")
			fmt.Printf("o1:%s\ngo o2 :%s", output, output1)
			t.Fail()
			return
		}
		fmt.Printf("Test %d passed\n", i)
	}
	cmd3 := exec.Command("rm", "test.exe")
	_, err = cmd3.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Print the command's output
	fmt.Print("done testing\n")
}
