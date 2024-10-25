package typescore

type FilteringParamsList struct {
	Offset         *uint64
	Limit          *uint64
	LikeFields     map[string]string
	OrSearchFields map[string]bool
}

type SortingType string

const (
	SortTypeASC  SortingType = "ASC"  // Сортировка по убыванию
	SortTypeDESC SortingType = "DESC" // Сортировка по возрастанию
)

var AllSortTypes = []SortingType{
	SortTypeASC,
	SortTypeDESC,
}
