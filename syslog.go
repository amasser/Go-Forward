package main

import (
	"bytes"
	"text/template"
	"time"
)

type severity uint8
type facility uint8
type priority uint8

type syslogMessage struct {
	Facility  facility
	Severity  severity
	Message   string
	Syslogtag string
	Hostname  string
	timestamp time.Time
}

const maxMsgLen = 2048

// From /usr/include/sys/syslog.h.
const (
	logEmerg severity = iota
	logAlert
	logCrit
	logErr
	logWarning
	logNotice
	logInfo
	logDebug
)

var severityMap = map[severity]string{
	logEmerg:   "EMERG",
	logAlert:   "ALERT",
	logCrit:    "CRIT",
	logErr:     "ERR",
	logWarning: "WARNING",
	logNotice:  "NOTICE",
	logInfo:    "INFO",
	logDebug:   "DEBUG",
}

// From /usr/include/sys/syslog.h.
const (
	logKern facility = iota
	logUser
	logMail
	logDaemon
	logAuth
	logSyslog
	logLpr
	logNews
	logUucp
	logClock
	logAuthpriv
	logFtp
	logNtp
	logLogaudit
	logLogalert
	logCron
	logLocal0
	logLocal1
	logLocal2
	logLocal3
	logLocal4
	logLocal5
	logLocal6
	logLocal7
)

var facilityMap = map[facility]string{
	logKern:     "KERN",
	logUser:     "USER",
	logMail:     "MAIL",
	logDaemon:   "DAEMON",
	logAuth:     "AUTH",
	logSyslog:   "SYSLOG",
	logLpr:      "LPR",
	logNews:     "NEWS",
	logUucp:     "UUCP",
	logClock:    "CLOCK",
	logAuthpriv: "AUTHPRIV",
	logFtp:      "FTP",
	logNtp:      "NTP",
	logLogaudit: "LOGAUDIT",
	logLogalert: "LOGALERT",
	logCron:     "CRON",
	logLocal0:   "LOCAL0",
	logLocal1:   "LOCAL1",
	logLocal2:   "LOCAL2",
	logLocal3:   "LOCAL3",
	logLocal4:   "LOCAL4",
	logLocal5:   "LOCAL5",
	logLocal6:   "LOCAL6",
	logLocal7:   "LOCAL7",
}

func (s severity) String() string {
	if val, ok := severityMap[s]; ok {
		return val
	}
	return "UNKNOWN"
}

func (f facility) String() string {
	if val, ok := facilityMap[f]; ok {
		return val
	}
	return "UNKNOWN"
}

func (p priority) decode() (facility, severity) {
	return facility(p / 8), severity(p % 8)
}

func (s syslogMessage) render(format string) (string, error) {
	tmpl, err := template.New("").Parse(format)
	if err != nil {
		return "", err
	}
	buf := bytes.NewBuffer([]byte{})
	err = tmpl.Execute(buf, s)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
