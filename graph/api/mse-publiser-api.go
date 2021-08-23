package api

import (
	"bytes"
	"encoding/json"
	"mse-gql/graph/model"
	"net/http"
	"strconv"
)

var baseURL = "http://localhost:8082"

func FetchRecipe(id int) (*model.Recipe, error) {
	var err error
	var client = &http.Client{}
	var data model.Recipe

	request, err := http.NewRequest("GET", baseURL+"/get?id="+strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func InsertRecipe(rec model.NewRecipe) (*model.InsertResult, error) {
	var err error
	var data model.InsertResult

	req, err := json.Marshal(rec)

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(baseURL+"/insert", "application/json", bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func UpdateRecipe(rec model.UpdateRecipe) (*model.UpdateDeleteResult, error) {
	var err error
	var client = &http.Client{}
	var data model.UpdateDeleteResult

	req, err := json.Marshal(rec)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("PUT", baseURL+"/update", bytes.NewBuffer(req))
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func DeleteRecipe(id int) (*model.UpdateDeleteResult, error) {
	var err error
	var client = &http.Client{}
	var data model.UpdateDeleteResult

	request, err := http.NewRequest("DELETE", baseURL+"/delete?id="+strconv.Itoa(id), nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
