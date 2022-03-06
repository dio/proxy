package xdsserver

type Logger struct{}

func (logger Logger) Debugf(format string, args ...interface{}) {}
func (logger Logger) Infof(format string, args ...interface{})  {}
func (logger Logger) Warnf(format string, args ...interface{})  {}
func (logger Logger) Errorf(format string, args ...interface{}) {}
