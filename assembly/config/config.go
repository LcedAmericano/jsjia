package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/mantoucat/jsjia/utils/file"
	"github.com/spf13/viper"
)

var (
	cfg      *viper.Viper
	FilePath = "conf/application.yml"
	once     sync.Once
)

func init() {
	once.Do(func() {
		err := Init(FilePath)
		if err != nil {
			fmt.Println(fmt.Sprintf("err: %s", err))
		}
	})
}

func Init(filePath ...string) (err error) {
	if len(filePath) > 0 {
		FilePath = filePath[0]
	}

	cfg = viper.New()

	var cfgPath string
	cfgPath, err = configFilePath()
	if err != nil {
		return
	}

	cfg.SetConfigFile(cfgPath)
	cfg.SetConfigType("yaml")

	if err = cfg.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("config file not found")
		} else {
			return errors.New("read config file error")
		}
	}

	return
}

func configFilePath() (path string, err error) {
	var (
		appPath  string
		workPath string
	)
	if appPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}
	workPath, err = os.Getwd()
	if err != nil {
		return
	}

	if path = filepath.Join(appPath, FilePath); file.Exists(path) {
		return
	}
	if path = filepath.Join(workPath, FilePath); file.Exists(path) {
		return
	}

	return "", errors.New("no config file found")
}

// GetString 获取一个字符串配置项
func GetString(key string, def ...string) string {
	result := cfg.GetString(key)
	if result == "" && len(def) > 0 {
		return def[0]
	}
	return result
}

// GetInt 获取一个int配置项
func GetInt(key string, def ...int) (result int) {
	result = cfg.GetInt(key)
	if result == 0 && len(def) > 0 && !cfg.IsSet(key) {
		result = def[0]
	}
	return
}

// GetStringSlice 获取一个字符串切片
func GetStringSlice(key string) []string {
	return cfg.GetStringSlice(key)
}

// GetBool 获取布尔值
func GetBool(key string) bool {
	return cfg.GetBool(key)
}

// IsSet 是否设置了某个key
func IsSet(key string) bool {
	return cfg.IsSet(key)
}

// GetStringMap 获取map sring
func GetStringMap(key string) map[string]interface{} {
	return cfg.GetStringMap(key)
}

// Core ...
func Core() *viper.Viper {
	return cfg
}
