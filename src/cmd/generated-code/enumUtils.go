package swagger

func (category ExpenseCategory) String() string {
	switch category {
	case FOOD:
		return "FOOD"
	case ENTERTAINMENT:
		return "ENTERTAINMENT"
	case TRANSPORTATION:
		return "TRANSPORTATION"
	}
	return "DEFAULT"
}
