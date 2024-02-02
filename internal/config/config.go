package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/samber/lo"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func NewConfig() *viper.Viper {
	configPath := filepath.Join(lo.Must(os.UserConfigDir()), "protocol.toml")
	v := viper.New()
	v.AutomaticEnv()
	v.AddConfigPath(configPath)

	return v
}

func ListenConfig(viper *viper.Viper) {
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println(in.Name)
	})
}
