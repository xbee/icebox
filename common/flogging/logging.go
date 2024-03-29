/*
Copyright IBM Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package flogging

import (
	//"io"
	"os"
	//"regexp"
	//"strings"
	"sync"

	//"github.com/op/go-logging"

	"errors"
	"fmt"
	"github.com/rs/zerolog"
)

const (
	pkgLogID      = "flogging"
	defaultFormat = "%{color}%{time:2006-01-02 15:04:05.000 MST} [%{module}] %{shortfunc} -> %{level:.4s} %{id:03x}%{color:reset} %{message}"
	defaultLevel  = zerolog.InfoLevel
)

var (
	//logger zerolog.Logger

	defaultOutput *os.File

	modules map[string]string // Holds the map of all modules and their respective log level

	lock sync.RWMutex
	once sync.Once
)

func init() {
	//logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	//logger = logging.MustGetLoggerWithDefaultLevel(pkgLogID)
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "msg"

	Reset()
	initgrpclogger()
}

// Reset sets to logging to the defaults defined in this package.
func Reset() {
	modules = make(map[string]string)
	lock = sync.RWMutex{}

	defaultOutput = os.Stderr
	//log.Logger = log.Logger.Output(defaultOutput)
	//InitBackend(SetFormat(defaultFormat), defaultOutput)
	//InitFromSpec("")
}

// InitBackend sets up the logging backend based on
// the provided logging formatter and I/O writer.
//func InitBackend(formatter logging.Formatter, output io.Writer) {
//	backend := logging.NewLogBackend(output, "", 0)
//	backendFormatter := logging.NewBackendFormatter(backend, formatter)
//	logging.SetBackend(backendFormatter).SetLevel(defaultLevel, "")
//}

// DefaultLevel returns the fallback value for loggers to use if parsing fails.
func DefaultLevel() string {
	return defaultLevel.String()
}

// GetModuleLevel gets the current logging level for the specified module.
func GetModuleLevel(module string) (*string, error) {
	// logging.GetLevel() returns the logging level for the module, if defined.
	// Otherwise, it returns the default logging level, as set by
	// `flogging/logging.go`.
	//level := zerolog.GetLevel(module).String()
	//return level
	if value, ok := modules[module]; ok {
		return &value, nil
	} else {
		return nil, errors.New(fmt.Sprintf("module not found: %s", module))
	}

}

// SetModuleLevel sets the logging level for the modules that match the supplied
// regular expression. Can be used to dynamically change the log level for the
// module.
//func SetModuleLevel(moduleRegExp string, level string) (string, error) {
//	return setModuleLevel(moduleRegExp, level, true, false)
//}

//func setModuleLevel(moduleRegExp string, level string, isRegExp bool, revert bool) (string, error) {
//	var re *regexp.Regexp
//	logLevel, err := zerolog.LogLevel(level)
//	if err != nil {
//		logger.Warn().Msgf("Invalid logging level '%s' - ignored", level)
//	} else {
//		if !isRegExp || revert {
//			logging.SetLevel(logLevel, moduleRegExp)
//			logger.Debug().Msgf("Module '%s' logger enabled for log level '%s'", moduleRegExp, level)
//		} else {
//			re, err = regexp.Compile(moduleRegExp)
//			if err != nil {
//				logger.Warn().Msgf("Invalid regular expression: %s", moduleRegExp)
//				return "", err
//			}
//			lock.Lock()
//			defer lock.Unlock()
//			for module := range modules {
//				if re.MatchString(module) {
//					//logging.SetLevel(logging.Level(logLevel), module)
//					zerolog.SetGlobalLevel(zerolog.InfoLevel)
//					modules[module] = logLevel.String()
//					logger.Debug().Msgf("Module '%s' logger enabled for log level '%s'", module, logLevel)
//				}
//			}
//		}
//	}
//	return logLevel.String(), err
//}

// MustGetLoggerWithDefaultLevel is used in place of `logging.MustGetLoggerWithDefaultLevel` to allow us to
// store a map of all modules and submodules that have loggers in the system.
func MustGetLoggerWithDefaultLevel(module string) *zerolog.Logger {
	dftLvl := zerolog.GlobalLevel()
	l := zerolog.New(defaultOutput).
		Level(dftLvl).
		With().Str("m", module).
		Timestamp().Logger()

	lock.Lock()
	defer lock.Unlock()
	modules[module] = dftLvl.String()

	return &l
}

func MustGetLogger(module string, lvl zerolog.Level) *zerolog.Logger {
	l := zerolog.New(defaultOutput).
		Level(lvl).
		With().Str("m", module).
		Timestamp().Logger()

	lock.Lock()
	defer lock.Unlock()
	modules[module] = lvl.String()

	return &l
}

// InitFromSpec initializes the logging based on the supplied spec. It is
// exposed externally so that consumers of the flogging package may parse their
// own logging specification. The logging specification has the following form:
//		[<module>[,<module>...]=]<level>[:[<module>[,<module>...]=]<level>...]
//func InitFromSpec(spec string) string {
//	levelAll := defaultLevel
//	var err error
//
//	if spec != "" {
//		fields := strings.Split(spec, ":")
//		for _, field := range fields {
//			split := strings.Split(field, "=")
//			switch len(split) {
//			case 1:
//				if levelAll, err = logging.LogLevel(field); err != nil {
//					logger.Warn().Msgf("Logging level '%s' not recognized, defaulting to '%s': %s", field, defaultLevel, err)
//					levelAll = defaultLevel // need to reset cause original value was overwritten
//				}
//			case 2:
//				// <module>[,<module>...]=<level>
//				levelSingle, err := logging.LogLevel(split[1])
//				if err != nil {
//					logger.Warn().Msgf("Invalid logging level in '%s' ignored", field)
//					continue
//				}
//
//				if split[0] == "" {
//					logger.Warn().Msgf("Invalid logging override specification '%s' ignored - no module specified", field)
//				} else {
//					modules := strings.Split(split[0], ",")
//					for _, module := range modules {
//						logger.Debug().Msgf("Setting logging level for module '%s' to '%s'", module, levelSingle)
//						logging.SetLevel(levelSingle, module)
//					}
//				}
//			default:
//				logger.Warn().Msgf("Invalid logging override '%s' ignored - missing ':'?", field)
//			}
//		}
//	}
//
//	logging.SetLevel(levelAll, "") // set the logging level for all modules
//
//	// iterate through modules to reload their level in the modules map based on
//	// the new default level
//	for k := range modules {
//		MustGetLoggerWithDefaultLevel(k)
//	}
//	// register flogging logger in the modules map
//	MustGetLoggerWithDefaultLevel(pkgLogID)
//
//	return levelAll.String()
//}
