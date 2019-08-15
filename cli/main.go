package main

/*
__________        .__                
\______   \_______|__| ______ _____  
 |     ___/\_  __ \  |/  ___//     \ 
 |    |     |  | \/  |\___ \|  Y Y  \
 |____|     |__|  |__/____  >__|_|  /
                          \/      \/ 

 ---==-=---=---===----=-=-=-=----=-=-=-=-----=-=-=*--=-=-=-=-------

 => File: ./cli/main.go
 => Package: v1
 => Version: 1.0.0
 => Author: Jonathan McAllister
 => Purpose: 
    Basic example for implementing a CLI in GO

-=-==-=---==-=-=-=-=-----=-=-=-=---=----------=-=-=-=-=-=-=-=----

*/


// Imports
// ----------------------------------
import (

  "fmt"
  "strings"
  "github.com/spf13/cobra"

)

// Global Variables
// ----------------------------------
var (

  BuildVersion string = ""
  GitCommitHash string = ""

)

// Main Entry Point
// ----------------------------------
func main() {

  // -- Get Versioning info
  version, githash, compiledate := GetVersion(`./public/html/version.txt`)

  // -- Display Header:
  fmt.Printf("\033[H\033[2J")
  fmt.Printf("__________        .__                       \n")         
  fmt.Printf("\\______   \\_______|__| ______ _____       \n")
  fmt.Printf("|     ___/\\_  __ \\  |/  ___//     \\      \n")
  fmt.Printf("|    |     |  | \\/  |\\___ \\|  Y Y  \\.   \n")
  fmt.Printf("|____|     |__|  |__/____  >__|_|  /.       \n")
  fmt.Printf(".                        \\/      \\/                                \n") 

  fmt.Printf("---==-=---=---===----=-=-=-=----=-=-=-=-----=-=-=*--=-=-=-=------- \n")
  fmt.Printf(" => Micro Service: main.go\n")
  fmt.Printf(" => Version: %s \n", version)
  fmt.Printf(" => Git Commit: %s \n", githash)
  fmt.Printf(" => Compilation Date: %s \n", compiledate)
  fmt.Printf(" => Runtime Mode: %s \n", config.ServerSetting.RunMode)
  fmt.Printf(" => Author: Jonathan McAllister \n")
  fmt.Printf(" => Purpose: Prism Micro Service \n\n")

  // -- Variable Declarations
  var echoTimes int
  var cmdPrint = &cobra.Command{

    Use:   "print [string to print]",
    Short: "Print anything to the screen",
    Long: `print is for printing anything back to the screen. For many years people have printed back to the screen.`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Print: " + strings.Join(args, " "))
    },
  }

  var cmdEcho = &cobra.Command{

    Use:   "echo [string to echo]",
    Short: "Echo anything to the screen",
    Long: `echo is for echoing anything back. Echo works a lot like print, except it has a child command.`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Print: " + strings.Join(args, " "))
    },
  }

  var cmdTimes = &cobra.Command{
    Use:   "times [string to echo]",
    Short: "Echo anything to the screen more times",
    Long: `echo things multiple times back to the user by providing a count and a string.`,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      for i := 0; i < echoTimes; i++ {
        fmt.Println("Echo: " + strings.Join(args, " "))
      }
    },
  }

  cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

  var rootCmd = &cobra.Command{Use: "app"}
  rootCmd.AddCommand(cmdPrint, cmdEcho)
  cmdEcho.AddCommand(cmdTimes)
  rootCmd.Execute()
  }

