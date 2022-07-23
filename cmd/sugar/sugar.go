package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/serialt/cli/pkg"
	"github.com/serialt/cli/sugar"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	APPVersion = "v0.2"
	BuildTime  = "20060102"
	GitCommit  = "ccccccccccccccc"

	configFile string
)

var rootCmd = &cobra.Command{
	Use:   "cli ",
	Short: "cli toolkit",
	Long:  `cli toolkit`,
	Run:   RunServer,
}

func RunServer(cmd *cobra.Command, args []string) {
	sugar.Log.Infof("config: %v", sugar.Config)
}

func initConfig() {
	homedir.Dir()
	configFile, _ = homedir.Expand(configFile)
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("failed to read config file: %v\n", err)
		return
	}
	// 初始化配置文件
	err := viper.Unmarshal(&sugar.Config)
	if err != nil {
		fmt.Println("配置文件转结构体失败")
		return
	}
	// 初始化日志
	sugar.Log = pkg.NewSugarLogger(sugar.Config.Log.LogLevel, sugar.Config.Log.LogFile, "", false)

}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "config.yaml", "config file path")
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(versionCmd)
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

// command: version
// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print the version of Gins",
	Long:  "print the version of Gins",
	Run:   DisplayVersion,
}

func DisplayVersion(cmd *cobra.Command, args []string) {
	fmt.Printf("APPVersion: %v  BuildTime: %v  GitCommit: %v\n",
		APPVersion,
		BuildTime,
		GitCommit)
}
