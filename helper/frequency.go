package helper

type FrequencyCustomValue struct {
	Value     int    `json:"value"`
	Frequency string `json:"frequency"`
}

// ConvertFrequencyToNumber converts frequency string to number of days
func ConvertFrequencyToNumber(frequency string, customValue *FrequencyCustomValue) int {
	switch frequency {
	case "Daily":
		return 1
	case "Weekly":
		return 7
	case "Monthly":
		return 30
	case "Annually":
		return 365
	case "Custom":
		if customValue != nil {
			return customValue.Value * ConvertCustomValueFrequency(customValue.Frequency)
		}
		return 0
	default:
		return 0
	}
}

// ConvertCustomValueFrequency converts custom frequency unit to days
func ConvertCustomValueFrequency(frequency string) int {
	switch frequency {
	case "Day":
		return 1
	case "Week":
		return 7
	case "Month":
		return 30
	case "Year":
		return 365
	default:
		return 0
	}
}
