/**
 * @Author: Sun
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 2022/6/15 16:54
 */

package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	IPAddr struct {
		Addr string `yaml:"Addr"`
		Port string `yaml:"Port"`
	} `yaml:"IPAddr"`
	PGConfig struct {
		Host     string `yaml:"Host"`
		DBPort   string `yaml:"DBPort"`
		DBName   string `yaml:"DBName"`
		UserName string `yaml:"UserName"`
		Password string `yaml:"Password"`
	} `yaml:"Postgresql"`
	RemoteAddr struct {
		Addr string `yaml:"Addr"`
		Port string `yaml:"Port"`
		URl  string `yaml:"URL"`
	} `yaml:"RemoteAddr"`
}

// NewConfig 解析配置文件
func NewConfig() *Config {
	execPath, _ := os.Getwd()
	filePath := filepath.Join(execPath, "config/config.yaml")
	fmt.Println(filePath)
	//yamlFile, err := ioutil.ReadFile("/Users/Sun/DTCloud-code/wg315/config/config.yaml")
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var c = new(Config)
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	fmt.Println(c)
	return c
}
