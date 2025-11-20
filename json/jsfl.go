package json

import (
	"fmt"
	"os"

	"github.com/kr/pretty"
	"github.com/spf13/viper"
)

type configstructure struct {
	MacPass     string `mapstructure:"macos"`
	LinuxPass   string `mapstructure:"linux"`
	WindowsPass string `mapstructure:"windows"`
	PostHost    string `mapstructure:"postgres"`
	MySQLHost   string `mapstructure:"mysql"`
	MongoHost   string `mapstructure:"mongodb"`
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
	fmt.Printf("Using config file: %s\n", CONFIG)

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
	var t configstructure
	err := viper.Unmarshal(&t)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	pretty.Print(t)

}
