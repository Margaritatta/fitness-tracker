package daysteps

import (
	errors
	"fmt"
	"strconv"
	"strings"
	"time"

	"yourmodule/spentenergy"
	"yourmodule/personaldata"
)

type DaySteps struct {
	Steps      int
	Duration   time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) == 2 {
		return errors.New("нecorrect...") 
	}
	...
