package api

import (
	"errors"
	"fmt"
	"strconv"
)

type Format string

const (
	OneLiner   Format = "oneliner"
	MultiLiner        = "multiliner"
	Pretty            = "pretty"
	Custom            = "custom"
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

var counter int = 0

func FormatStatusBar(statusBar StatusBar) (result string, err error) {
	result = fmt.Sprintf("%s %s", statusBar.Data.GrandTotal.Text, statusBar.Data.Languages[0].Text)
	return result, err
}


// TODO: Make this reusable if we want to implement
// formatter for other API endpoints
func FormatGoalsOneliner(goals Goals, opts any) (result string, err error) {
	if opts == nil {
		return "", errors.New(`Opts not provided this is required when format is set to 'oneliner'
Example: wakago get goals 0 oneliner`)
	}

	if len(goals.Data) == 0 {
		return "", errors.New("No goals set")
	}

	var idx int

	i, ok := opts.(string)

	idx, err = strconv.Atoi(i)
	if !ok {
		return "", errors.New("Invalid opts")
	}

	if idx > len(goals.Data)-1 {
		return "", errors.New("Invalid goal index")
	}

	goal := goals.Data[idx]

	//	for _, v := range goal.ChartData {
	//		fmt.Printf("\n%s %s", v.Range, v.ActualSecondsText)
	//	}

	today := goal.ChartData[len(goal.ChartData)-1]
	// This will always be the current date

	result = fmt.Sprintf("%s %s of %s", today.RangeStatus, today.ActualSecondsText, today.GoalSecondsText)
	return
}

func FormatGoalsPretty(goals Goals) string {
	return ""
}

func FormatGoalsCustom(goals Goals, opts any) (string, error) {
	if opts == nil {
		return "", errors.New("Opts not provided, ")
	}
	s, ok := opts.(string)
	if !ok {
		return "", errors.New("Invalid opts")
	}

	return s, nil
}

func FormatGoalsMultiLiner(goals Goals) string {
	var result string

	for _, g := range goals.Data {
		var current string
		result += fmt.Sprintf("\n\n%s - %s\n", g.Title, g.CumulativeStatus)
		today := g.ChartData[len(g.ChartData)-1]
		current = fmt.Sprintf("\n %s %s %s of %s", today.Range.Date, today.RangeStatus, today.ActualSecondsText, today.GoalSecondsText)
		result += current
	}

	return result
}

func FormatGoals(goals Goals, format Format, opts any) (string, error) {
	switch format {
	case OneLiner:
		return FormatGoalsOneliner(goals, opts)
	case MultiLiner:
		return FormatGoalsMultiLiner(goals), nil
	case Pretty:
		return FormatGoalsPretty(goals), nil
	case Custom:
		return FormatGoalsCustom(goals, opts)
	}

	return "", errors.New("Invalid format")
}
