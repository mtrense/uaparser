package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/mtrense/uaparser"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version = "none"
	commit  = "none"
	app     = &cobra.Command{
		Use:   "uaparser",
		Short: "Parser for User Agent",
		Run:   execute,
	}
	cmdVersion = &cobra.Command{
		Use:   "version",
		Short: "Show uaparsers version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Version: %s (ref: %s)\n", version, commit)
		},
	}
)

func init() {
	app.AddCommand(cmdVersion)
	viper.SetEnvPrefix("UAPARSER")
	viper.AutomaticEnv()
}

func main() {
	if err := app.Execute(); err != nil {
		panic(err)
	}
}

func execute(cmd *cobra.Command, args []string) {
	reader := csv.NewReader(os.Stdin)
	writer := csv.NewWriter(os.Stdout)
	writer.Write([]string{
		"Browser_Name",
		"Browser_Version",
		"Engine_Name",
		"Engine_Version",
		"Mobile",
		"Bot",
		"Mozilla",
		"Platform",
		"Localization",
		"OS_Name",
		"OS_FullName",
		"OS_Version",
	})
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		line := record[0]
		ua := uaparser.Parse(line)
		writer.Write([]string{
			ua.Browser.Name,
			ua.Browser.Version,
			ua.Engine.Name,
			ua.Engine.Version,
			strconv.FormatBool(ua.Mobile),
			strconv.FormatBool(ua.Bot),
			ua.Mozilla,
			ua.Platform,
			ua.Localization,
			ua.OS.Name,
			ua.OS.FullName,
			ua.OS.Version,
		})
	}
	writer.Flush()
}
