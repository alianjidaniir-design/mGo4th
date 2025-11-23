package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func aliasNormalizeFunc(f *pflag.FlagSet, n string) pflag.NormalizedName {

	switch n {
	case "pass":
		n = "password"
		break
	case "ps":
		n = "password"
		break
	}

	return pflag.NormalizedName(n)
}

func main() {
	fmt.Println("Dana")
	pflag.StringP("name", "n", "Ali", "Name parameter")
	pflag.StringP("password", "p", "abc", "password for the user")
	pflag.CommandLine.SetNormalizeFunc(aliasNormalizeFunc)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return
	}
	name := viper.GetString("name")
	password := viper.GetString("password")
	fmt.Println(name, password)
	err = viper.BindEnv("GOMAXPROCS")
	if err != nil {
		return
	}
	val := viper.Get("GOMAXPROCS")
	if val != nil {
		fmt.Println("GOMAXPROCS", val)
		viper.Set("GOMAXPROCS", 16)

		val = viper.Get("GOMAXPROCS")

		fmt.Println("GOMAXPROCS:=", val)

	}

}
