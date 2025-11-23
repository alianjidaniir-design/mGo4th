package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type configStructure struct {
	MacPass     string `mapStructure:"macos"`
	LinuxPass   string `mapStructure:"linux"`
	WindowsPass string `mapStructure:"windows"`
	PostHost    string `mapStructure:"postgres"`
	MySQLHost   string `mapStructure:"mysql"`
	MongoHost   string `mapStructure:"mongodb"`
}

var CONFIG = ".config.json"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default file", CONFIG)
	} else {
		CONFIG = os.Args[1]
	}
	viper.SetConfigType("json")
	viper.SetConfigFile(CONFIG)
	fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	if viper.IsSet("macos") {
		fmt.Println("macos:", viper.Get("macos"))
	} else {
		fmt.Println("macos not set")
	}

	if viper.IsSet("active") {
		value := viper.GetBool("active")
		if value {
			fmt.Println("Active")
			postgres := viper.Get("postgres")
			mySql := viper.Get("mysql")
			mongo := viper.Get("mongodb")
			fmt.Println("postgres:", postgres, "mySql:", mySql, "mongo:", mongo)
		}
	} else {
		fmt.Println("Active is not set!")
	}

	if !viper.IsSet("DoesNotExist") {
		fmt.Println("DoesNotExist")
	}
	var t configStructure
	err = viper.Unmarshal(&t)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

}
