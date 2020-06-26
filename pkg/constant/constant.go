package constant

type CollectionName string
const (
	INGREDIENT_COLLECTION CollectionName = "ingredient"
	RECIPE_COLLECTION CollectionName = "recipe"
)

type OpType string

const (
	ADD OpType = "ADD"
	SUBTRACT OpType = "SUBTRACT"
)