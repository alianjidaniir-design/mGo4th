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
	pflag.StringP("name", "nm", "Ali", "Name parameter")
	pflag.StringP("password", "p", "abc", "password for the user")
	pflag.CommandLine.SetNormalizeFunc(aliasNormalizeFunc)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

}
