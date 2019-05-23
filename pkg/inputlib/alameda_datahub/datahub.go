package main

import (
	"errors"
	"github.com/containers-ai/alameda/operator/datahub"
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/spf13/viper"
	"strings"
	"encoding/json"
	"fmt"
)

type inputLib struct {
	config  datahub.Config  `mapstructure:"datahub"`
	scope   *logUtil.Scope
}

func (i inputLib) Gather() error {
	return nil
}

func (i inputLib) LoadConfig(config string, scope *logUtil.Scope) error {
	i.scope = scope

	viper.SetEnvPrefix("InputLibDataHub")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	viper.SetConfigFile(config)
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("Read input library datahub configuration failed: " + err.Error()))
	}
	err = viper.Unmarshal(&i.config)
	if err != nil {
		panic(errors.New("Unmarshal input library datahub configuration failed: " + err.Error()))
	} else {
		if transmitterConfBin, err := json.MarshalIndent(i.config, "", "  "); err == nil {
			scope.Infof(fmt.Sprintf("Input library datahub configuration: %s", string(transmitterConfBin)))
		}
	}

	return nil
}

var InputLib inputLib
