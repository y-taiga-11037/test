package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type logFormat struct {
	TimestampFormat string
}

// Set the log format
func (f *logFormat) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	b.WriteByte('[')
	b.WriteString(strings.ToUpper(entry.Level.String()))
	b.WriteString("]:")
	b.WriteString(entry.Time.Format(f.TimestampFormat))

	b.WriteString(" [")
	b.WriteString(formatFilePath(entry.Caller.File))
	b.WriteString(":")
	fmt.Fprint(b, entry.Caller.Line)
	b.WriteString("] ")

	if entry.Message != "" {
		b.WriteString(" - ")
		b.WriteString(entry.Message)
	}

	if len(entry.Data) > 0 {
		b.WriteString(" || ")
	}
	for key, value := range entry.Data {
		b.WriteString(key)
		b.WriteByte('=')
		b.WriteByte('{')
		fmt.Fprint(b, value)
		b.WriteString("}, ")
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

// Runs when the package is loaded
func init() {
	logrus.SetReportCaller(true) // To handle executable files
	formatter := logFormat{}
	formatter.TimestampFormat = "2006-01-02 15:04:05" // Time Setting

	logrus.SetFormatter(&formatter)

	/* TODO:
	drone test does not recognize file path
	The log file and the environment setting file are read by the CLI
	Fix it when the CLI is installed


	// Configuring the log output file
	/*	f, err := openFile("/home/y-taiga/mdtd_bootcamp/tmp/api.log")
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.SetOutput(io.MultiWriter(os.Stdout, f)) */

	logrus.SetOutput(io.MultiWriter(os.Stdout))

	// Outputs a level higher than the set level
	logrus.SetLevel(logrus.DebugLevel)

}
