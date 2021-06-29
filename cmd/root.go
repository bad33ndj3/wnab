package cmd

import (
	"fmt"
	"path"
	"strings"

	"github.com/davecgh/go-spew/spew"

	"go.bmvs.io/ynab/api/category"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.bmvs.io/ynab"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = newRootCmd(ynab.NewClient(""))

func newRootCmd(ynabClient ynab.ClientServicer) *cobra.Command {
	return &cobra.Command{
		Use:   "wnab",
		Short: "We Need A Budget",
		Long: `We Need A Budget (WNAB) is a CLI tool that is meant to be used for couples that use YNAB.
This tool will sync the needed 'income' of a 'shared account' to a specific 'budget' in each partners personal YNAB budget.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			budgets, err := ynabClient.Budget().GetBudgets()
			if err != nil {
				return err
			}

			var categories []*category.GroupWithCategories
			for i := range budgets {
				curCategories, err := ynabClient.Category().GetCategories(budgets[i].ID)
				if err != nil {
					return err
				}
				categories = append(categories, curCategories...)
			}

			var person1, person2 category.Category
			for i := range categories {
				for j := range categories[i].Categories {
					if strings.Contains(categories[i].Categories[j].Name, "#A") {
						person1 = *categories[i].Categories[j]
					}
					if strings.Contains(categories[i].Categories[j].Name, "#B") {
						person2 = *categories[i].Categories[j]
					}
				}
			}

			spew.Dump(person1, person2)

			return nil
		},
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.wnab.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".wnab" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".wnab")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)
		configPath := path.Join(home, ".wnab.yaml")
		err = viper.SafeWriteConfigAs(configPath)
		cobra.CheckErr(err)
		fmt.Println("Created config file:", configPath)
	}
}
