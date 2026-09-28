package spentcalories

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, errors.New("invalid data format")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, err
	}
	if steps <= 0 {
		return 0, "", 0, errors.New("invalid steps value")
	}
	activityType := parts[1]
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, err
	}
	if duration <= 0 {
		return 0, "", 0, errors.New("invalid duration value")
	}
	return steps, activityType, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	lenStep := height * stepLengthCoefficient
	return float64(lenStep*float64(steps)) / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	speed := distance(steps, height) / duration.Hours()
	return speed
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("invalid input values")
	}
	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := (weight * speed * minutes) / minInH
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("invalid input values")
	}
	speed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	calories := (weight * speed * minutes) / minInH
	return calories * walkingCaloriesCoefficient, nil
}
