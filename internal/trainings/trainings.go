package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return fmt.Errorf("Invalid format")
	}
	t.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid format datastring steps")
	}
	if t.Steps <= 0 {
		return fmt.Errorf("the number of steps must be greater than 0")
	}
	t.TrainingType = parts[1]
	t.Duration, err = time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("invalid format datastring duration")
	}
	if t.Duration <= 0 {
		return fmt.Errorf("duration must be greater than 0")
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var err error
	userCalories := 0.0
	distance := spentenergy.Distance(t.Steps, t.Height)
	averageSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	switch t.TrainingType {
	case "Бег":
		userCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		userCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		return "", fmt.Errorf("calorie calculation error: %w", err)
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %0.2f ч.\nДистанция: %0.2f км.\nСкорость: %0.2f км/ч\nСожгли калорий: %0.2f\n", t.TrainingType, t.Duration.Hours(), distance, averageSpeed, userCalories), nil
}
