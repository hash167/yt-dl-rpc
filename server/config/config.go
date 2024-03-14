package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Config struct {
	CurrentLogFile  string
	LogPath         string `yaml:"log_path"`
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	DownloadPath    string `yaml:"downloadPath"`
	DownloaderPath  string `yaml:"downloaderPath"`
	QueueSize       int    `yaml:"queue_size"`
	SessionFilePath string `yaml:"session_file_path"`
}

var (
	instance     *Config
	instanceOnce sync.Once
)

func Instance() *Config {
	if instance == nil {
		instanceOnce.Do(func() {
			instance = &Config{}
		})
	}
	return instance
}

func (c *Config) LoadFile(filename string) error {
	fd, err := os.Open(filename)
	if err != nil {
		return err
	}

	if err := yaml.NewDecoder(fd).Decode(c); err != nil {
		return err
	}

	return nil
}

func (c *Config) String() string {
	return "host: " + c.Host + " port: " + string(rune(c.Port)) + " downloadPath: " + c.DownloadPath + " downloaderPath: " + c.DownloaderPath + " queueSize: " + string(rune(c.QueueSize)) + " sessionFilePath: " + c.SessionFilePath
}
