//go:generate go run _generate/dependencies/main.go
package main

import (
	"fmt"
	"os"
)

// //////////////////////////////////////////////////////////////

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Помилка:", err)
		os.Exit(1)
	}
}
