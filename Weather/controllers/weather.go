package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var apiKey = os.Getenv("WEATHER_API_KEY")

func GetWeather(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "City parameter is required"})
		return
	}

	url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s",
		strings.ToLower(city),
		apiKey,
	)

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch weather data"})
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"city":        result["name"],
		"temperature": result["main"].(map[string]interface{})["temp"],
		"condition":   result["weather"].([]interface{})[0].(map[string]interface{})["description"],
	})
}
