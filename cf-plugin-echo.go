// adaptation of:
//    https://github.com/cloudfoundry/cli/blob/master/plugin_examples/echo.go

//------------------------------------------------------------------------------
// This is an example plugin where we use both arguments and flags. The plugin
// will echo all arguments passed to it. The flag -uppercase will upcase the
// arguments passed to the command.
//------------------------------------------------------------------------------
package main

//------------------------------------------------------------------------------
import (
  "flag"
  "fmt"
  "os"
  "strings"

  pluginAPI "github.com/cloudfoundry/cli/plugin"
)

//------------------------------------------------------------------------------
type Plugin struct {
  uppercase *bool
}

//------------------------------------------------------------------------------
func main() {
  pluginAPI.Start(new(Plugin))
}

//------------------------------------------------------------------------------
func (plugin *Plugin) Run(cliConn pluginAPI.CliConnection, args []string) {
  // Initialize flags
  argsOpts  := flag.NewFlagSet("echo", flag.ExitOnError)
  uppercase := argsOpts.Bool("uppercase", false, "displays all provided text in uppercase")

  // Parse starting from [1] because the [0]th element is the name of the command
  err := argsOpts.Parse(args[1:])
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  var result string
  for _, value := range argsOpts.Args() {
    if *uppercase {
      result += strings.ToUpper(value) + " "
    } else {
      result += value + " "
    }
  }

  fmt.Println(result)
}

//------------------------------------------------------------------------------
func (plugin *Plugin) GetMetadata() pluginAPI.PluginMetadata {
  return pluginAPI.PluginMetadata{
    Name: "echo",
    Version: pluginAPI.VersionType{
      Major: 0,
      Minor: 1,
      Build: 0,
    },
    Commands: []pluginAPI.Command{
      {
        Name:     "echo",
        Alias:    "echo-what-is-passed-in",
        HelpText: "Echo text passed into the command. To obtain more information use --help",
        UsageDetails: pluginAPI.Usage{
          Usage: "echo - print input arguments to screen\n   cf echo [-uppercase] text",
          Options: map[string]string{
            "uppercase": "If this param is passed, which ever word is passed to echo will be all capitals.",
          },
        },
      },
    },
  }
}
