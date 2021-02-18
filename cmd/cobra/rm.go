package cobra

import (
	"fmt"

	"github.com/hi20160616/enit"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := enit.File(encodingKey, secretPath())
		switch {
		case len(args) == 0:
			fmt.Println("key is none.")
			return
		case len(args) != 1:
			fmt.Println("too many args.")
			return
		}
		key := args[0]
		if err := v.Remove(key); err != nil {
			fmt.Println("Secret remove error: ", err)
			return
		}
		fmt.Printf("Secret: %s remove successfully.\n", key)
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}
