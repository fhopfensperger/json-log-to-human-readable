/*
Copyright © 2020 Florian Hopfensperger <f.hopfensperger@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os"
)

var cfgFile string
var alternativeInput bool

var globalUsage = `A simple command line utility to transform one line json log message to a human readable output for example:

content test.json: { "level": "INFO", "timestamp": "2020-07-14T09:38:14.977Z", "message": "sample output" }
cat test.json | json-log-to-human-readable
tail -f test.json | json-log-to-human-readable
kubectl logs -f -n default pod-name-1 | json-log-to-human-readable`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "json-log-to-human-readable",
	Short: "Transforms json log message to a human readable output",
	Long:  globalUsage,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		return runCommand()
	},
	Args: NoArgs,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	rootCmd.Version = version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(globalUsage)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.json-log-to-human-readable.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().BoolVarP(&alternativeInput, "alternative", "a", false, "Spring Boot JSON input")
	rootCmd.SetVersionTemplate(`{{printf "%s\n" .Version}}`)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".json-log-to-human-readable" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".json-log-to-human-readable")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func NoArgs(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return errors.Errorf(
			"%q accepts no arguments\n\nUsage:  %s",
			cmd.CommandPath(),
			cmd.UseLine(),
		)
	}
	return nil
}

func runCommand() error {
	if isInputFromPipe() {
		return toHumanReadable(os.Stdin)
	}
	return errors.New("Input must be pipe")
}

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func toHumanReadable(r io.Reader) error {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	for scanner.Scan() {
		byteValue := scanner.Bytes()

		if alternativeInput {
			var logMessage AlternativeLogMessage
			err := json.Unmarshal(byteValue, &logMessage)
			if err != nil {
				fmt.Println(string(byteValue))
				continue
			}
			logMessage.print()

		} else {
			var logMessage LogMessage
			err := json.Unmarshal(byteValue, &logMessage)
			if err != nil {
				fmt.Println(string(byteValue))
				continue
			}
			logMessage.print()
		}

	}
	return nil
}
