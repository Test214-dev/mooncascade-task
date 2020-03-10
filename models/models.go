package models

import "time"

type Athlete struct {
	ChipID      string `json:"chipId"` // uuid
	FullName    string `json:"fullName" binding:"required"`
	StartNumber int    `json:"startNumber" binding:"required"`
}

type AddTimingRequest struct {
	ChipID    string    `json:"chipId" binding:"required"`  // uuid
	PointID   string    `json:"pointId" binding:"required"` // finish line or finish corridor
	Timestamp time.Time `json:"timestamp" binding:"required"`
}

type Timing struct {
	TimingID  string    `json:"timingId"` // uuid
	PointID   string    `json:"pointId"`  // finish line or finish corridor
	Timestamp time.Time `json:"timestamp"`
	ChipID    string    `json:"chipId"`
	Athlete   `json:"athlete"`
}

type AppError struct {
	Code  int    `json:"-"`
	Error string `json:"error"`
}
