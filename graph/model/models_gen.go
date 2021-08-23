// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type InsertResult struct {
	Result string  `json:"result"`
	Data   *Recipe `json:"data"`
}

type NewRecipe struct {
	Name         string   `json:"name"`
	Ingredients  []string `json:"ingredients"`
	IsHalal      bool     `json:"isHalal"`
	IsVegetarian bool     `json:"isVegetarian"`
	Description  string   `json:"description"`
	Rating       float64  `json:"rating"`
}

type Recipe struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Ingredients  []string `json:"ingredients"`
	IsHalal      bool     `json:"isHalal"`
	IsVegetarian bool     `json:"isVegetarian"`
	Description  string   `json:"description"`
	Rating       float64  `json:"rating"`
}

type SearchParam struct {
	Keyword      string `json:"keyword"`
	IsHalal      bool   `json:"isHalal"`
	IsVegetarian bool   `json:"isVegetarian"`
	Page         int    `json:"page"`
	Take         int    `json:"take"`
}

type UpdateDeleteResult struct {
	Result string `json:"result"`
}

type UpdateRecipe struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Ingredients  []string `json:"ingredients"`
	IsHalal      bool     `json:"isHalal"`
	IsVegetarian bool     `json:"isVegetarian"`
	Description  string   `json:"description"`
	Rating       float64  `json:"rating"`
}
