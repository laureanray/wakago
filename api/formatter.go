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
