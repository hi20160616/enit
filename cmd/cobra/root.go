package cobra

import (
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "enit",
	Short: "Enit(Encrypt it) is an API key and other secrets manager",
}

var encodingKey string

func init() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "", "the key to use when encoding and decoding secrets")
}

func secretPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".enit")
}
