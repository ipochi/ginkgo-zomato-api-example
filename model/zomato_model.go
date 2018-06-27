package model

// City Struct
type City struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CountryID string `json:"country_id"`
	Country   string `json:"country_name"`
}

type Cities struct {
	Cities []City
}

// Category of Restaurants
type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Categories struct {
	Categories []Category
}

type Restaurant struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	AvgCostForTwo float64  `json:"average_cost_for_two"`
	Currency      string   `json:"currency"`
	Location      Location `json:"location"`
	//Extra         map[string]interface{}
}

type Restaurants struct {
	Restaurants []Restaurant
	//Extra map[string]interface{}
}

type Location struct {
	Address   string `json:"address"`
	Locality  string `json:"locality"`
	City      string `json:"city"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Zipcode   string `json:"zipcode"`
	CountryID string `json:"country_id"`
}
