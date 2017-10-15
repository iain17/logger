package logger

type Logger struct {
	module string
}

func New(module string) *Logger {
	return &Logger{
		module: module,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	Debug(l.module, v)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	Debugf(l.module+"	"+format, v)
}

func (l *Logger) Info(v ...interface{}) {
	Info(l.module, v)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	Infof(l.module+"	"+format, v)
}

func (l *Logger) Warning(v ...interface{}) {
	Warning(l.module, v)
}

func (l *Logger) Warningf(format string, v ...interface{}) {
	Warningf(l.module+"	"+format, v)
}

func (l *Logger) Error(v ...interface{}) {
	Error(l.module, v)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	Errorf(l.module+"	"+format, v)
}

func (l *Logger) Fatal(v ...interface{}) {
	Fatal(l.module, v)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	Fatalf(l.module+"	"+format, v)
}