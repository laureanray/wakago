package api

import "fmt"

func FormatGoal(goalData GoalData) string {
	oneliner := fmt.Sprintf("[%s - %s] [%f] %s",
		goalData.Title,
		goalData.Type,
		goalData.ImproveByPercent,
		goalData.AverageStatus,
	)

	return oneliner
}

func FormatGoals(goals Goals) string {
	for _, g := range goals.Data {
		fmt.Printf("\n\n%s - %s\n", g.Title, g.CumulativeStatus)
		today := g.ChartData[len(g.ChartData)-1]
		current := fmt.Sprintf("\n %s %s %s of %s", today.Range.Date, today.RangeStatus, today.ActualSecondsText, today.GoalSecondsText)
		fmt.Print(current)
	}

	return ""
}
