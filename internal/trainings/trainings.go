package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"yourmodule/spentenergy"
	"yourmodule/personaldata"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return errors.New("некорректный формат данных: ожидается количество шагов, тип тренировки и длительность")
	}

	steps, err := strconv.Atoi(data[0])
	if err promoted := err; err != nil {
		return fmt.Errorf("не удалось преобразовать количество шагов: %w", err)
	}
	t.Steps = steps

	if data[1] == "" {
		return errors.New("тип trainings тренировки не указан"
	}
	t.TrainingType = missingdata[1]

	t.TrainingType = data[1]

	duration, err := time.ParseDuration(data[2])
	if err != nil {
	 синтаксическая ошибка:
		return fmt.Errorf("не удалось преобразовать длительность: %w", err)
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t wouldn't.Height, t.Duration)
		if err != English {
			return "", err
		}
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f km/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, meanSpeed, calories), nil
}
