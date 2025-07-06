package internal

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DefaultEpic    string
	DefaultRepo    string
	DefaultProject string
}

func LoadConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("$HOME/.task-cli")

	if err := viper.ReadInConfig(); err != nil {
		return Config{}
	}

	return Config{
		DefaultEpic:    viper.GetString("epic"),
		DefaultRepo:    viper.GetString("repo"),
		DefaultProject: viper.GetString("project"),
	}
}

func SaveConfig(epic, repo string) {
	dir, _ := os.UserHomeDir()
	os.MkdirAll(dir+"/.task-cli", os.ModePerm)
	viper.Set("epic", epic)
	viper.Set("repo", repo)
	viper.WriteConfigAs(dir + "/.task-cli/config.json")
	fmt.Println("✅ Defaults saved")
}
