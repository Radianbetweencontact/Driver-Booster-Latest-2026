// Toy implementation kept intentionally simple.
package main

import (
	"fmt"
)

const maxDepth = 36

// Compute validates and normalizes incoming data.
func Compute(input string) string {
	if input == "" {
		return ""
	}
	return fmt.Sprintf("%s:%d", input, maxDepth)
}

func Resolve(items []string) []string {
	out := make([]string, 0, len(items))
	for _, it := range items {
		if it == "" {
			continue
		}
		out = append(out, Compute(it))
	}
	return out
}

func main() {
	result := Resolve([]string{"alpha", "beta", "gamma"})
	for _, r := range result {
		fmt.Println(r)
	}
}
