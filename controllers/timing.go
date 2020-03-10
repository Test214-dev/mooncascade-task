package controllers

import (
	"errors"
	"github.com/Test214-dev/mooncascade-task/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TimingController struct {
	DB *gorm.DB
}

const (
	FinishLine     = "finish line"
	FinishCorridor = "finish corridor"
)

func (tc *TimingController) insertTiming(p *models.Timing) (string, *models.AppError) {
	var res models.Timing
	if err := tc.DB.Raw("INSERT INTO timing_points(timing_id, point_id, timestamp, chip_id) VALUES"+
		"(uuid_generate_v4(), $1, $2, $3) RETURNING timing_id",
		p.PointID, p.Timestamp, p.ChipID).Scan(&res).Error; err != nil {
		return "", &models.AppError{Code: 500, Error: err.Error()}
	}

	return res.TimingID, nil
}

func (tc *TimingController) getTiming(timingID string) (*models.Timing, *models.AppError) {
	result := models.Timing{}
	if err := tc.DB.Raw("SELECT * FROM timing_points t JOIN athletes a on t.chip_id = a.chip_id WHERE timing_id = $1", timingID).Scan(&result).Error; err != nil {
		code := 500
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = 404
		}
		return nil, &models.AppError{Code: code, Error: err.Error()}
	}

	return &result, nil
}

func (tc *TimingController) listTimings() ([]models.Timing, *models.AppError) {
	result := make([]models.Timing, 0)

	if err := tc.DB.Raw(`SELECT * FROM timing_points t JOIN athletes a ON t.chip_id = a.chip_id ORDER BY t.timestamp DESC`).Scan(&result).Error; err != nil {
		return nil, &models.AppError{Code: 500, Error: err.Error()}
	}

	return result, nil
}

// @Summary Add new timing
// @Accept json
// @Param	mommy	body	models.AddTimingRequest	true	"timing data"
// @Success 201 "Location header will contain generated timing id"
// @Failure 400 {object} models.AppError
// @Failure 500 {object} models.AppError
// @Router /api/timings/ [post]
func (tc *TimingController) HandleTimingPost(c *gin.Context) {
	pointRequest := models.AddTimingRequest{}
	if err := c.BindJSON(&pointRequest); err != nil {
		setError(c, &models.AppError{Code: 400, Error: err.Error()})
		return
	}

	if pointRequest.PointID != FinishLine && pointRequest.PointID != FinishCorridor {
		setError(c, &models.AppError{Code: 400, Error: `pointId must be one of "finish line", "finish corridor"`})
		return
	}

	point := models.Timing{ChipID: pointRequest.ChipID, Timestamp: pointRequest.Timestamp, PointID: pointRequest.PointID}
	id, err := tc.insertTiming(&point)
	if err != nil {
		setError(c, err)
		return
	}

	c.Header("Location", id)
	c.Status(201)
}

// @Summary Get a timing by ID
// @Param	id	path	string	true	"timing id"
// @Success 200 {object} models.Timing
// @Failure 404 {object} models.AppError
// @Failure 500 {object} models.AppError
// @Router /api/timings/:id [get]
func (tc *TimingController) HandleTimingGet(c *gin.Context) {
	pointID := c.Param("id")
	res, err := tc.getTiming(pointID)
	if err != nil {
		setError(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary List all timings
// @Success 200 {array} models.Timing
// @Failure 500 {object} models.AppError
// @Router /api/timings/ [get]
func (tc *TimingController) HandleTimingList(c *gin.Context) {
	res, err := tc.listTimings()
	if err != nil {
		setError(c, err)
		return
	}
	c.JSON(200, res)
}
