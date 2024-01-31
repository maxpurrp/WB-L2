package model

type Event struct {
	ID   int    `json:"id"`
	Date string `json:"date"`
}

type Response struct {
	Events []Event `json:"events"`
}

type Config struct {
	Port string `mapstructure:"port"`
}
