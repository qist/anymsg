package cmd

import (
	"github.com/spf13/cobra"
	"os"

	//	"github.com/spf13/viper"

)

//var cfgFile string
var Config string


// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "anymsg",
	Short: "anymsg",
	Long:  `anymsg is any message can be sent`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(Config) == 0 {
			cmd.Help()
			os.Exit(-1)

		}

	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error{
	if err := RootCmd.Execute(); err != nil {
		return err
		//fmt.Println(err)
		//os.Exit(-1)
	}
	return nil
}

func init() {
	//	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	//	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.demo.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	RootCmd.Flags().StringVarP(&Config, "config", "f", "", "指定配置文件路径")
}
