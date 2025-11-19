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
