package initial

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

const configFileName = "config.json"

var (
	configDir         string
	configManagerIns  *ConfigManager
	configManagerOnce sync.Once
	downloadedDir     string
)

func makeDownloadDir() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err.Error())
	}
	downloadedDir = filepath.Join(userHomeDir, "Downloads")
}

func makeConfigDir() {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err.Error())
	}

	configDir = filepath.Join(userConfigDir, "esdumpweb")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		log.Fatal(err.Error())
	}
}

func init() {
	makeDownloadDir()
	makeConfigDir()
	makeLogDir()
}

type HostConfig struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	Addrs     map[string]HostConfig `json:"addrs"`
	AddrName  string                `json:"addr_name"`
	Index     string                `json:"index"`
	TimeField string                `json:"time_field"`
	Product   string                `json:"product"`
	SaveDir   string                `json:"save_dir"`
	Condition string                `json:"condition"`
}

type ConfigManager struct {
	Config         *Config
	configFile     *os.File
	configFilePath string
}

func WireConfigManager() *ConfigManager {
	configManagerOnce.Do(func() {
		configManagerIns = &ConfigManager{}
		configManagerIns.init()
	})
	return configManagerIns
}

func WireConfig() *Config {
	return WireConfigManager().Config
}

func (c *ConfigManager) init() {
	var err error
	// read config from file
	c.configFilePath = filepath.Join(configDir, configFileName)
	c.configFile, err = os.OpenFile(c.configFilePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("failed to initilize config manager. %v", err.Error()))
	}
	buf, err := io.ReadAll(c.configFile)
	if err != nil {
		panic(fmt.Sprintf("failed to initilize config manager. %v", err.Error()))
	}
	c.Config = new(Config)
	if len(buf) != 0 {
		err = json.Unmarshal(buf, &c.Config)
		if err != nil {
			fmt.Printf("failed to initilize config manager. %v\n", err.Error())
		}
	}

	if c.Config.SaveDir == "" {
		c.Config.SaveDir = downloadedDir
	}
	if c.Config.Addrs == nil {
		c.Config.Addrs = make(map[string]HostConfig)
	}
}

func (c *ConfigManager) Save() error {
	if c.configFile == nil {
		return nil
	}
	defer c.configFile.Close()
	return c.Flush()
}

func (c *ConfigManager) Flush() error {
	buf, err := json.MarshalIndent(c.Config, "", "  ")
	if err != nil {
		return err
	}
	_, err = c.configFile.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Printf("failed to flush config. %v\n", err.Error())
	}
	_, err = c.configFile.Write(buf)
	if err != nil {
		return err
	}
	return nil
}
