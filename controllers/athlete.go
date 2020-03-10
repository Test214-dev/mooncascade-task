package controllers

import (
	"errors"
	"github.com/Test214-dev/mooncascade-task/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type AthleteController struct {
	DB *gorm.DB
}

func (ac *AthleteController) insertAthlete(a *models.Athlete) (string, *models.AppError) {
	var res models.Athlete
	if err := ac.DB.Raw("INSERT INTO athletes(chip_id, start_number, full_name) VALUES"+
		"(uuid_generate_v4(), $1, $2) RETURNING chip_id",
		a.StartNumber, a.FullName).Scan(&res).Error; err != nil {
		return "", &models.AppError{Code: 500, Error: err.Error()}
	}

	return res.ChipID, nil
}

func (ac *AthleteController) getAthlete(chipID string) (*models.Athlete, *models.AppError) {
	result := models.Athlete{}
	if err := ac.DB.Raw("SELECT * FROM athletes WHERE chip_id = $1", chipID).Scan(&result).Error; err != nil {
		code := 500
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = 404
		}
		return nil, &models.AppError{Code: code, Error: err.Error()}
	}

	return &result, nil
}

func (ac *AthleteController) listAthletes() ([]models.Athlete, *models.AppError) {
	result := make([]models.Athlete, 0)

	if err := ac.DB.Raw("SELECT * FROM athletes").Scan(&result).Error; err != nil {
		return nil, &models.AppError{Code: 500, Error: err.Error()}
	}

	return result, nil
}

// @Summary Add a new athlete
// @Accept json
// @Param	body	body	models.Athlete	true	"athlete data"
// @Success 201 "Location header will contain generated point id"
// @Failure 400 {object} models.AppError
// @Failure 500 {object} models.AppError
// @Router /api/athletes/ [post]
func (ac *AthleteController) HandleAthletePost(c *gin.Context) {
	athlete := models.Athlete{}
	if err := c.BindJSON(&athlete); err != nil {
		setError(c, &models.AppError{Code: 400, Error: err.Error()})
		return
	}
	id, err := ac.insertAthlete(&athlete)
	if err != nil {
		setError(c, err)
		return
	}

	c.Header("Location", id)
	c.Status(201)
}

// @Summary Get an athlete by ID
// @Param	id	path	string	true	"athlete id"
// @Success 200 {object} models.Athlete
// @Failure 404 {object} models.AppError
// @Failure 500 {object} models.AppError
// @Router /api/athletes/:id [get]
func (ac *AthleteController) HandleAthleteGet(c *gin.Context) {
	chipID := c.Param("id")
	res, err := ac.getAthlete(chipID)
	if err != nil {
		setError(c, err)
		return
	}

	c.JSON(200, res)
}

// @Summary List all athletes
// @Success 200 {array} models.Athlete
// @Failure 500 {object} models.AppError
// @Router /api/athletes/ [get]
func (ac *AthleteController) HandleAthleteList(c *gin.Context) {
	res, err := ac.listAthletes()
	if err != nil {
		setError(c, err)
		return
	}

	c.JSON(200, res)
}
