package main

import (
	"github.com/hi20160616/enit/cmd/cobra"
)

func main() {
	cobra.RootCmd.Execute()
	// v := enit.File("my-fake-key", ".secrets")
	// err := v.Set("demo_key", "some crazy value")
	// if err != nil {
	//         panic(err)
	// }
	// plain, err := v.Get("demo_key")
	// if err != nil {
	//         panic(err)
	// }
	// fmt.Println("Plain:", plain)
}
