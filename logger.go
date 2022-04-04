package logtamer

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/gotamer/mail/send"
)

// Configuration
type LogCfg struct {
	AppName    string
	File       string
	SendMailTo string // Application Admin Email
	Smtp       struct {
		Hostname string // smtp.example.com
		Hostport int    // 587
		Username string // usally email address username@example.com
		Password string
	}
}

const FAKEHOSTNAME = "example.com"

var (
	o     = new(LogCfg)
	mail  *send.SmtpTamer
	Debug = *log.Default()
	Info  = *log.Default()
	Warn  = *log.Default()
	Error = *log.Default()
)

func init() {
	// Config file defaults
	o.File = "GoTamerLogger.log"
	o.SendMailTo = "admin@example.com"
	o.Smtp.Hostname = FAKEHOSTNAME
	o.Smtp.Hostport = 587
	o.Smtp.Username = "user@example.com"
	o.Smtp.Password = "Long-Complicated-Secret-Pass-Word"
	Debug.SetPrefix("DBG ")
	Info.SetPrefix("INF ")
	Warn.SetPrefix("WRN ")
	Error.SetPrefix("ERR ")
	Info.SetFlags(log.Ltime | log.Lshortfile)
	Debug.SetFlags(log.Ltime | log.Lshortfile)
	Warn.SetFlags(log.Ltime | log.Lshortfile)
	Error.SetFlags(log.Ltime | log.Lshortfile)
}

// filepath.Join(appPath, appName+".log")
func Default(appName, file string) {
	o.AppName = appName
	o.File = file
}

// Switch between log levels
// Level 0 No Logging at all
// Level 1 output all to screen, is also default
// Level 2 info, warn & error to screen, debug to file
// Level 3 warn & error to screen, debug and info to file
// Level 4 warn & error to file
// Level 5 error to file
// Level 6 warn & error to mail
// Level 7 error to mail
func Level(level uint8) {

	Debug = *log.Default()
	Info = *log.Default()
	Warn = *log.Default()
	Error = *log.Default()
	Debug.SetPrefix("DBG ")
	Info.SetPrefix("INF ")
	Warn.SetPrefix("WRN ")
	Error.SetPrefix("ERR ")
	Info.SetFlags(log.Ltime | log.Lshortfile)
	Debug.SetFlags(log.Ltime | log.Lshortfile)
	Warn.SetFlags(log.Ltime | log.Lshortfile)
	Error.SetFlags(log.Ltime | log.Lshortfile)

	file, err := os.OpenFile(o.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	switch level {
	case 0:
		// Level 0 No Logging at all
		Debug.SetOutput(ioutil.Discard)
		Info.SetOutput(ioutil.Discard)
		Warn.SetOutput(ioutil.Discard)
		Error.SetOutput(ioutil.Discard)
	case 1:
		// Level 1 output all to screen, is also default
	case 2:
		// Level 2 info, warn & error to screen, debug to file
		// Log all but debug to screen aka debug mode. Debug goes to file
		Debug.SetOutput(file)
	case 3:
		// Level 3 warn & error to screen, debug and info to file
		Info.SetOutput(file)
		Debug.SetOutput(file)
	case 4:
		// Level 4 warn & error to file
		Warn.SetOutput(file)
		Error.SetOutput(file)
	case 5:
		// Level 5 error to file
		Error.SetOutput(file)
	case 6:
		// Level 6 warn & error to mail
		if o.Smtp.Hostname == "" || o.Smtp.Hostname == FAKEHOSTNAME {
			log.Fatal("ERR smtp config not set")
		}
		// email errors to admin
		mail = send.NewSMTP(o.Smtp.Hostname, o.Smtp.Username, o.Smtp.Password)
		mail.Envelop.SetFrom(o.AppName, o.Smtp.Username)
		mail.Envelop.SetTo("", o.SendMailTo)
		mail.Envelop.Subject = "ERROR " + o.AppName

		Debug.SetOutput(ioutil.Discard)
		Info.SetOutput(ioutil.Discard)
		Warn.SetOutput(mail)
		Error.SetOutput(mail)
	case 7:
		// Level 7 error to mail
		if o.Smtp.Hostname == "" || o.Smtp.Hostname == FAKEHOSTNAME {
			log.Fatal("ERR smtp config not set")
		}
		// email errors to admin
		mail = send.NewSMTP(o.Smtp.Hostname, o.Smtp.Username, o.Smtp.Password)
		mail.Envelop.SetFrom(o.AppName, o.Smtp.Username)
		mail.Envelop.SetTo("", o.SendMailTo)
		mail.Envelop.Subject = "ERROR " + o.AppName

		Debug.SetOutput(ioutil.Discard)
		Info.SetOutput(ioutil.Discard)
		Warn.SetOutput(ioutil.Discard)
		Error.SetOutput(mail)

	default:
		// Same as Level 1, output all to screen
	}
}
