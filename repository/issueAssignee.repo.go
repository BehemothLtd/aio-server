package repository

import (
	"aio-server/models"
	"context"

	"gorm.io/gorm"
)

type IssueAssigneeRepository struct {
	Repository
}

func NewIssueAssigneeRepository(c *context.Context, db *gorm.DB) *IssueAssigneeRepository {
	return &IssueAssigneeRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

func (r *IssueAssigneeRepository) CountByProjectAssignee(projectAssignee models.ProjectAssignee) int64 {
	var count int64

	r.db.Model(&models.IssueAssignee{}).Where("user_id = ?", projectAssignee.UserId).Where("issue_id IN (?)", r.db.Table("issues").Select("id").Where("project_id = ?", projectAssignee.ProjectId)).Count(&count)
	return count
}
