package cobra

import (
	"fmt"

	"github.com/hi20160616/enit"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a secret in your secret storage",
	Run: func(cmd *cobra.Command, args []string) {
		v := enit.File(encodingKey, secretPath())
		switch len(args) {
		case 0:
			fmt.Println("key and value is none.")
			return
		case 1:
			fmt.Println("value is none")
			return
		}
		key, value := args[0], args[1]
		if err := v.Set(key, value); err != nil {
			fmt.Println("value set error: ", err)
			return
		}
		fmt.Println("Value set successfully.")
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
