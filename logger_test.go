package logtamer

import (
	"fmt"
	"testing"
)

var logfile = "test.log"
var appname = "LogTest"

// Switch between log levels
// Level 0 No Logging at all
// Level 1 output all to screen, is also default
// Level 2 info, warn & error to screen, debug to file
// Level 3 warn & error to screen, debug and info to file
// Level 4 warn & error to file
// Level 5 error to file
// Level 6 warn & error to mail
// Level 7 error to mail

func TestDefault(t *testing.T) {
	var lev = "Default"
	fmt.Println("FMT Testing ", lev)
	Debug.Println("Level ", lev)
	Info.Println("Level ", lev)
	Warn.Println("Level ", lev)
	Error.Println("Level ", lev)
	fmt.Println("FMT end Testing ", lev)
}

func TestLevel0(t *testing.T) {
	var lev uint8 = 0
	fmt.Println("FMT changing Default Level ", lev)
	Default(appname, logfile)
	Level(lev)
	fmt.Println("FMT Testing ", lev)
	Debug.Println("Level ", lev)
	Info.Println("Level ", lev)
	Warn.Println("Level ", lev)
	Error.Println("Level ", lev)
	fmt.Println("FMT end Testing ", lev)
}

func TestLevel2(t *testing.T) {
	var lev uint8 = 2
	fmt.Println("FMT changing Default Level ", lev)
	Default(appname, logfile)
	Level(lev)
	fmt.Println("FMT Testing ", lev)
	Debug.Println("Level ", lev)
	Info.Println("Level ", lev)
	Warn.Println("Level ", lev)
	Error.Println("Level ", lev)
	fmt.Println("FMT end Testing ", lev)
}

func TestLevel3(t *testing.T) {
	var lev uint8 = 3
	fmt.Println("FMT changing Default Level ", lev)
	Default(appname, logfile)
	Level(lev)
	fmt.Println("FMT Testing ", lev)
	Debug.Println("Level ", lev)
	Info.Println("Level ", lev)
	Warn.Println("Level ", lev)
	Error.Println("Level ", lev)
	fmt.Println("FMT end Testing ", lev)
}
func TestLevel4(t *testing.T) {
	var lev uint8 = 4
	fmt.Println("FMT changing Default Level ", lev)
	Default(appname, logfile)
	Level(lev)
	fmt.Println("FMT Testing ", lev)
	Debug.Println("Level ", lev)
	Info.Println("Level ", lev)
	Warn.Println("Level ", lev)
	Error.Println("Level ", lev)
	fmt.Println("FMT end Testing ", lev)
}
func TestLevel5(t *testing.T) {
	var lev uint8 = 5
	fmt.Println("FMT changing Default Level ", lev)
	Default(appname, logfile)
	Level(lev)
	fmt.Println("FMT Testing ", lev)
	Debug.Println("Level ", lev)
	Info.Println("Level ", lev)
	Warn.Println("Level ", lev)
	Error.Println("Level ", lev)
	fmt.Println("FMT end Testing ", lev)
	FileClose()
}
