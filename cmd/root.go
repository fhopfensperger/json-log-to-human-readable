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
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var springBootInput bool
var uberZapInput bool
var dotnetInput bool

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
	Args: noArgs,
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
	rootCmd.PersistentFlags().BoolVarP(&dotnetInput, "dotnet", "d", false, ".NET JSON input")
	rootCmd.PersistentFlags().BoolVarP(&springBootInput, "springboot", "s", false, "Spring Boot JSON input")
	rootCmd.PersistentFlags().BoolVarP(&uberZapInput, "zap", "z", false, "Uber zap JSON Input")
	rootCmd.SetVersionTemplate(`{{printf "v%s\n" .Version}}`)
}

func noArgs(cmd *cobra.Command, args []string) error {
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
		return toHumanReadable(os.Stdin, os.Stdout)
	}

	return errors.New("Input must be pipe")
}

func isInputFromPipe() bool {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println(err)
	}

	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func toHumanReadable(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(bufio.NewReader(r))
	buf := make([]byte, 0, 64*1024) //nolint:gomnd // only used once
	// increase max buffer size to process large log messages
	scanner.Buffer(buf, 1024*1024) //nolint:gomnd // only used once

	for scanner.Scan() {
		byteValue := scanner.Bytes()

		switch {
		case uberZapInput:
			var logMessage GoZapLogMessage

			err := json.Unmarshal(byteValue, &logMessage)
			if err != nil {
				fmt.Fprintln(w, string(byteValue))
				continue
			}

			logMessage.transform(w)
		case springBootInput:
			var logMessage SpringBootLogMessage

			err := json.Unmarshal(byteValue, &logMessage)
			if err != nil {
				fmt.Fprintln(w, string(byteValue))
				continue
			}

			logMessage.transform(w)
		case dotnetInput:
			var logMessage DotNetLogMessage

			err := json.Unmarshal(byteValue, &logMessage)
			if err != nil {
				fmt.Fprintln(w, string(byteValue))
				continue
			}

			logMessage.transform(w)
		default:
			var logMessage QuarkusLogMessage

			err := json.Unmarshal(byteValue, &logMessage)
			if err != nil {
				fmt.Fprintln(w, string(byteValue))
				continue
			}

			logMessage.transform(w)
		}
	}

	return scanner.Err()
}
