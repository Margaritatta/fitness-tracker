package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"yourmodule/personaldata"
	"yourmodule/spentenergy"
)

type DaySteps struct {
	Steps      int
	Duration   time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return errors.New("некорректный формат данных: ожидается количество шагов и длительность")
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("не удалось преобразовать количество шагов: %w", err)
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(data[1])
	if err != nil {
		return fmt.Errorf("не удалось преобразовать длительность: %w", err)
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories), nil
}
