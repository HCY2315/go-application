package casbin

import (
	"sync/atomic"
)

type Logger struct {
	enable int32
}

func (l *Logger) EnableLog(enable bool) {
	i := 0
	if enable {
		i = 1
	}
	atomic.StoreInt32(&(l.enable), int32(i))
}

func (l *Logger) IsEnabled() bool {
	return atomic.LoadInt32(&(l.enable)) != 0
}

// TODO: 丰富以下方法

// LogModel log info related to model.
func (l *Logger) LogModel(model [][]string) {

}

// LogEnforce log info related to enforce.
func (l *Logger) LogEnforce(matcher string, request []interface{}, result bool, explains [][]string) {

}

// LogRole log info related to role.
func (l *Logger) LogRole(roles []string) {

}

// LogPolicy log info related to policy.
func (l *Logger) LogPolicy(policy map[string][][]string) {

}
