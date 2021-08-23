package api

import (
	"bytes"
	"encoding/json"
	"mse-gql/graph/model"
	"net/http"
)

var searchBaseURL = "http://localhost:8081"

func SearchRecipe(keyword string, isHalal, isVegetarian bool, page, take int) ([]*model.Recipe, error) {
	var err error
	var client = &http.Client{}
	var res []*model.Recipe

	req, err := json.Marshal(model.SearchParam{
		Keyword:      keyword,
		IsHalal:      isHalal,
		IsVegetarian: isVegetarian,
		Page:         page,
		Take:         take,
	})

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", searchBaseURL+"/search", bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
