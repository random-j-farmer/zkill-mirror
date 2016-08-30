// Package config handles the configuration of the app
//
// The following configuration files are tried:
// * $HOME/.ZKILL-MIRROR/zkill-mirror.toml: this allows us the override the current directory
// * current directory/zkill-mirror.toml
// * $HOME/.zkill-mirror/zkill-mirror.toml: user level
// * /etc/zkill-mirror/zkill-mirror.toml
//
// All configuration variables can be overriden by env vars in uppser case,
// with a prefix of ZKM_.
//
// An explanation of all the options is given in the example config file.
//
package config

import (
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// DBName gives the name of the bolt database
func DBName() string {
	return viper.GetString("db_name")
}

// BobsName gives the name of the blob store
func BobsName() string {
	return viper.GetString("bobs_name")
}

// PullDelay is the maximum time to wait before polling an idle zkillboard
func PullDelay() time.Duration {
	return viper.GetDuration("pull_delay")
}

// Port to listen on
func Port() int {
	return viper.GetInt("port")
}

// Verbose output
func Verbose() bool {
	return viper.GetBool("verbose")
}

// CacheTemplates - should templates be cached
func CacheTemplates() bool {
	return viper.GetBool("cache_templates")
}

func init() {
	viper.SetDefault("db_name", "zkill-mirror.bolt")
	viper.SetDefault("bobs_name", "zkill-mirror.bobs")
	viper.SetDefault("pull_delay", 5*time.Minute)
	viper.SetDefault("port", "8080")
	viper.SetDefault("verbose", false)
	viper.SetDefault("cache_templates", true)

	viper.SetConfigName("zkill-mirror")
	viper.SetConfigType("toml")
	viper.AddConfigPath(os.ExpandEnv("$HOME/.ZKILL-MIRROR"))
	viper.AddConfigPath(".")
	viper.AddConfigPath(os.ExpandEnv("$HOME/.zkill-mirror"))
	viper.AddConfigPath("/etc/zkill-mirror/")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(errors.Wrap(err, "config.init"))
	}

	viper.SetEnvPrefix("ZKM")
	viper.AutomaticEnv()

	if Verbose() {
		sep := "-------------------------------------"
		log.Println(sep)
		log.Println("Configuration:")
		log.Println("db_name:\t", DBName())
		log.Println("bobs_name:\t", BobsName())
		log.Println("port:\t", Port())
		log.Println("cache_templates:\t", CacheTemplates())
		log.Println("verbose:\t", Verbose())
		log.Println(sep)
	}
}
