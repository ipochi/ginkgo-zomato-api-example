package api

import (
	"encoding/json"

	"github.com/ipochi/ginkgo-zomato-api-example/model"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/resty.v1"
)

var baseurl = "https://developers.zomato.com/api/v2.1/"

// GetRestaurantsInACity returns restuarants
func GetRestaurantsInACity(apiKey, cityID, entityType, count string) (int, []*model.Restaurant, error) {

	resp, err := resty.R().
		SetQueryParams(map[string]string{
			"entity_id":   cityID,
			"entity_type": entityType,
			"count":       count,
		}).
		//SetHeader("Accept", "application/json").
		SetHeader("user-key", apiKey).
		Get(baseurl + "/search")

	if err != nil {
		return resp.StatusCode(), nil, err
	}
	mar := map[string]interface{}{}
	err = json.Unmarshal(resp.Body(), &mar)

	if err != nil {
		return resp.StatusCode(), nil, err
	}

	if mar["restaurants"] == nil {
		return resp.StatusCode(), nil, err
	}

	restaurantsResponse := mar["restaurants"].([]interface{})
	var restaurants []*model.Restaurant
	for _, value := range restaurantsResponse {

		_restaurant := &model.Restaurant{}

		restMap := value.(map[string]interface{})
		val := restMap["restaurant"]

		_rest := val.(map[string]interface{})

		id := _rest["id"]
		_restaurant.ID = id.(string)

		loc := _rest["location"]
		locMap := loc.(map[string]interface{})
		mapstructure.Decode(locMap, &_restaurant.Location)

		name := _rest["name"]
		_restaurant.Name = name.(string)

		avgCostForTwo := _rest["average_cost_for_two"]
		_restaurant.AvgCostForTwo = avgCostForTwo.(float64)

		currency := _rest["currency"]
		_restaurant.Currency = currency.(string)

		restaurants = append(restaurants, _restaurant)

	}
	return resp.StatusCode(), restaurants, err
}
