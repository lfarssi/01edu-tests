package main

import (
	"os/exec"
	"testing"
)

func TestLinearStats(t *testing.T) {
	cmd1 := exec.Command("go", "build", "-o", "main", "main.go")
	build, err := cmd1.CombinedOutput()
	if err != nil {
		t.Fatalf("Build failed: %v\nOutput: %s", err, string(build))
	}

	for i := 0; i < 1000; i++ {
		// Run the reference binary first
		output1, err := exec.Command("bin/linear-stats", "data.txt").CombinedOutput()
		if err != nil {
			t.Fatalf("Error running reference binary: %v\nOutput: %s", err, string(output1))
		}

		// Run your program (compiled main)
		output2, err := exec.Command("./main", "data.txt").CombinedOutput()
		if err != nil {
			t.Fatalf("Error running your program: %v\nOutput: %s", err, string(output2))
		}

		// Compare the outputs
		if string(output1) != string(output2) {
			t.Errorf("Test %d failed: outputs don't match\nRef:\n%s\nYour:\n%s", i, string(output1), string(output2))
			break
		}
	}
}
