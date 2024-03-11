package main

import (
	"fmt"

	"github.com/belastingdienst/opr-paas/internal/crypt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func encryptCmd() *cobra.Command {
	var publicKeyFile string
	var dataFile string
	var paasName string

	cmd := &cobra.Command{
		Use:   "encrypt [command options]",
		Short: "encrypt using public key and print results",
		Long:  `encrypt using public key and print results`,
		RunE: func(command *cobra.Command, args []string) error {
			if paasName == "" {
				return fmt.Errorf("a paas must be set with eith --paas or environment variabele PAAS_NAME")
			}
			if dataFile == "" {
				return crypt.EncryptFromStdin(publicKeyFile, paasName)
			} else {
				return crypt.EncryptFile(publicKeyFile, paasName, dataFile)
			}
		},
		Example: `crypttool encrypt --publicKeyFile "/tmp/pub" --fromFile "/tmp/decrypted" --paas my-paas`,
	}

	flags := cmd.Flags()
	flags.StringVar(&publicKeyFile, "publicKeyFile", "", "The file to read the public key from")
	viper.BindPFlag("publicKeyFile", flags.Lookup("publicKeyFile"))
	viper.BindEnv("publicKeyFile", "PAAS_PUBLIC_KEY_PATH")
	flags.StringVar(&dataFile, "dataFile", "", "The file to read the data to be encrypted from")
	viper.BindPFlag("dataFile", flags.Lookup("dataFile"))
	viper.BindEnv("dataFile", "PAAS_INPUT_FILE")
	flags.StringVar(&paasName, "paas", "", "The paas this data is to be encrypted for")
	viper.BindPFlag("paas", flags.Lookup("paas"))
	viper.BindEnv("paas", "PAAS_NAME")

	return cmd
}