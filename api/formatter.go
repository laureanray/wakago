package api

import (
	"errors"
	"fmt"
)

type Format int64

const (
	OneLiner Format = iota
	TwoLiner
	Pretty
	Custom
)

func FormatGoal(goalData GoalData) string {
	oneliner := fmt.Sprintf("[%s - %s] [%f] %s",
		goalData.Title,
		goalData.Type,
		goalData.ImproveByPercent,
		goalData.AverageStatus,
	)

	return oneliner
}

// TODO: Make this reusable if we want to implement
// formatter for other API endpoints
func FormatGoalsOneliner(goals Goals) string {
	return ""
}

func FormatGoalsPretty(goals Goals) string {
	return ""
}

func FormatGoalsCustom(goals Goals, customFormat string) string {
	return ""
}

func FormatGoalsTwoLiner(goals Goals) string {
	var result string

	for _, g := range goals.Data {
		var current string
		result += fmt.Sprintf("\n\n%s - %s\n", g.Title, g.CumulativeStatus)
		today := g.ChartData[len(g.ChartData)-1]
		current = fmt.Sprintf("\n %s %s %s of %s", today.Range.Date, today.RangeStatus, today.ActualSecondsText, today.GoalSecondsText)
		result += result + current
	}

	return result
}

func FormatGoals(goals Goals, format Format, opts string) (string, error) {
	switch format {
	case OneLiner:
		return FormatGoalsOneliner(goals), nil
	case TwoLiner:
		return FormatGoalsTwoLiner(goals), nil
	case Pretty:
		return FormatGoalsPretty(goals), nil
	case Custom:
		if len(opts) == 0 {
			return "", errors.New("Opts not provided")
		}
		return FormatGoalsCustom(goals, opts), nil
	}

	return "", errors.New("Invalid format")
}
