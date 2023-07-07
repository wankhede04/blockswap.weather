package weatherservice

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wankhede04/blockswap.weather/weather-srv/db"
)

// ReportWeatherHandler creates weather report and update lastCall for member
func (s *WeatherService) ReportWeatherHandler(c *gin.Context) {
	report, ok := c.Get("report")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Report not found"})
		return
	}
	reportStr, ok := report.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid report format"})
		return
	}

	membership, ok := c.Get("membership")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	m, ok := membership.(db.Membership)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	weatherReport := db.WeatherReport{MembershipID: m.ID, Report: reportStr}

	// Acquire a database connection
	db, err := s.getDBConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer s.releaseDBConnection(db)

	// Save the weather report and update the last call time in a single transaction
	tx := db.Begin()
	if err := tx.Create(&weatherReport).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	m.LastCall = time.Now().Unix()
	if err := tx.Save(&m).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Weather reported successfully"})
}
