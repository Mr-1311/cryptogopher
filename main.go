package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	Version     = "0.1.0-rc.1"
	Codename    = "Hidden"
	BuildStamp  = "0"
	BuildDate   time.Time
	BuildHost   = "unknown"
	BuildUser   = "unknown"
	IsRelease   bool
	IsCandidate bool
	IsBeta      bool
	LongVersion string
)

const (
	exitSuccess      = 0
	exitErrorGeneral = 1
	exitErrorParse   = 1
)

const (
	usage       = "cryptogopher [options]"
	usageDetail = `
TODO: Add Extra Usage Details
`
)

type RuntimeOptions struct {
	isEncryptMode bool
	isDecryptMode bool
	file          string
	quiet         bool
	showHelp      bool
	showVersion   bool
}

func setBuildMetadata() {
	exp := regexp.MustCompile(`^v\d+\.\d+\.\d+(-[a-z]+[\d\.]+)?$`)
	IsRelease = exp.MatchString(Version)
	IsCandidate = strings.Contains(Version, "-rc.")
	IsBeta = strings.Contains(Version, "-")

	stamp, _ := strconv.Atoi(BuildStamp)
	BuildDate = time.Unix(int64(stamp), 0)

	date := BuildDate.UTC().Format("2006-01-02 15:04:05 MST")
	LongVersion = fmt.Sprintf(`cryptogopher %s "%s" (%s %s-%s) %s@%s %s`, Version, Codename, runtime.Version(), runtime.GOOS, runtime.GOARCH, BuildUser, BuildHost, date)
}

func getDefaultRuntimeOptions() RuntimeOptions {
	options := RuntimeOptions{
		isEncryptMode: false,
		isDecryptMode: false,
		quiet:         false,
		file:          "",
		showHelp:      false,
		showVersion:   false,
	}

	return options
}

func parseCommandLineOptions() RuntimeOptions {
	options := getDefaultRuntimeOptions()

	flag.BoolVar(&options.showHelp, "h", false, "Show help, then exit")
	flag.BoolVar(&options.quiet, "q", false, "Be quiet")
	flag.BoolVar(&options.showVersion, "v", false, "Show version")
	flag.BoolVar(&options.isEncryptMode, "e", false, "Encrypt the given file")
	flag.BoolVar(&options.isDecryptMode, "d", false, "Decrypt the given file")
	flag.StringVar(&options.file, "f", options.file, "The file to be processed")

	flag.Usage = usageFor(flag.CommandLine, usage, usageDetail)
	flag.Parse()

	if len(flag.Args()) > 0 {
		fmt.Println("Error: Unrecognized option!")
		//flag.Usage()
		os.Exit(exitErrorParse)
	}

	return options
}

func main() {
	setBuildMetadata()

	options := parseCommandLineOptions()

	if options.showHelp {
		flag.Usage()
		os.Exit(exitSuccess)
	}

	if options.showVersion {
		fmt.Println(LongVersion)
		os.Exit(exitSuccess)
	}

	cryptogopherMain(options)
}

func cryptogopherMain(options RuntimeOptions) {
	if !options.isEncryptMode && !options.isDecryptMode {
		fmt.Println("Chose the main operation: [-e] or [-d]")
		os.Exit(exitErrorParse)
		return
	}

	if options.isEncryptMode && options.isDecryptMode {
		fmt.Println("Main operations can't set both at same time': [-e] and [-d]")
		os.Exit(exitErrorParse)
		return
	}

	if _, err := os.Stat(options.file); os.IsNotExist(err) {
		fmt.Println("You must enter a valid file name. Choose the file with: [-f]")
		os.Exit(exitErrorGeneral)
	} else {
		if options.isEncryptMode {
			CryptoHandler(options.file, 0)
			os.Exit(exitSuccess)
		}

		if options.isDecryptMode {
			CryptoHandler(options.file, 1)
			os.Exit(exitSuccess)
		}
	}

	fmt.Println("Reached end of the application..!")
}
