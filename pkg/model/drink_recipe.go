package model

//drink recipe struct
//do remember recipe name must be unique
type DrinkRecipe struct {
	RecipeName string
	Ingredients []*Ingredient
}
