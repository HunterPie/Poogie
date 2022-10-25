package models

import "github.com/Haato3o/poogie/core/persistence/patches"

type LatestVersionResponse struct {
	LatestVersion string `json:"latest_version"`
}

type AllPatchNotesResponse struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Banner      string `json:"banner"`
	Link        string `json:"link"`
}

func ToAllPatchNotesResponses(models []patches.Patch) []AllPatchNotesResponse {
	responses := make([]AllPatchNotesResponse, 0)

	for _, model := range models {
		responses = append(responses, AllPatchNotesResponse{
			Title:       model.Title,
			Description: model.Description,
			Banner:      model.Banner,
			Link:        model.Link,
		})
	}

	return responses
}
