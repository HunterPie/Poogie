package models

type VersionQueryRequest struct {
	Version string `form:"v" json:"v" binding:"required"`
}
