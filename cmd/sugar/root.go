package sugar

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/serialt/cli/pkg"
	"github.com/serialt/cli/t"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	APPVersion = "v0.2"
	BuildTime  = "20060102"
	GitCommit  = "ccccccccccccccc"
)

var rootCmd = &cobra.Command{
	Use:   "cli ",
	Short: "cli toolkit",
	Long:  `cli toolkit`,
	Run:   RunServer,
}

func RunServer(cmd *cobra.Command, args []string) {
	t.Sugar.Infof("config: %v", t.Config)
}

func initConfig() {
	homedir.Dir()
	if t.ConfigFile != "" {
		t.ConfigFile, _ = homedir.Expand(t.ConfigFile)
		viper.SetConfigFile(t.ConfigFile)
	} else {
		workDir, _ := os.Getwd()
		viper.SetConfigName("config.yaml")
		viper.AddConfigPath(workDir)
	}
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("failed to read config file: %v\n", err)
		return
	}
	// 初始化配置文件
	err := viper.Unmarshal(&t.Config)
	if err != nil {
		fmt.Println("配置文件转结构体失败")
		return
	}
	// 初始化日志
	t.Logger = pkg.NewLogger(t.Config.Log.LogLevel, t.Config.Log.LogFile)
	t.Sugar = t.Logger.Sugar()
	// 初始化数据库连接
	// t.DB = dao.NewDBConnect(&t.Config.Database)

}

func init() {
	rootCmd.PersistentFlags().StringVarP(&t.ConfigFile, "config", "c", "", "config file path")
	cobra.OnInitialize(initConfig)

}

func Execute() {
	rootCmd.Execute()

}
