package utils

import "smfbackend/models"

// ExtractProjectIds : extract project ids from project array
func ExtractProjectIds(projects []models.Project) []uint {
	ids := make([]uint, len(projects))
	for index, project := range projects {
		ids[index] = project.ID
	}
	return ids
}
