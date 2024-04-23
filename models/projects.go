package models

import (
	"aio-server/database"
	"aio-server/enums"
	"aio-server/pkg/constants"
	"slices"
	"time"

	"gorm.io/gorm"
)

type Project struct {
	Id                   int32
	Name                 string
	Code                 string
	Description          *string
	ProjectType          enums.ProjectType
	ProjectPriority      enums.ProjectPriority `gorm:"default:2"`
	State                enums.ProjectState    `gorm:"default: 1"`
	ActivedAt            *time.Time
	InactivedAt          *time.Time
	StartedAt            *time.Time
	EndedAt              *time.Time
	CreatedAt            time.Time
	UpdatedAt            time.Time
	SprintDuration       *int32
	ClientId             *int32
	Client               Client
	CurrentSprintId      *int32
	ProjectAssignees     []*ProjectAssignee
	ProjectIssueStatuses []*ProjectIssueStatus
	ProjectSprints       []ProjectSprint
	Issues               []Issue
	IssueStatuses        []IssueStatus `gorm:"many2many:project_issue_statuses;"`
	LockVersion          int32         `gorm:"default:1"`
	Logo                 *Attachment   `gorm:"polymorphic:Owner;polymorphicValue:Project"`
	Files                []*Attachment `gorm:"polymorphic:Owner;polymorphicValue:Project"`
}

func (p Project) HasEnoughProjectIssueStatuses() (bool, []string) {
	var requiredIssueStatusIds []int32

	if p.ProjectType == enums.ProjectTypeKanban {
		requiredIssueStatusIds = constants.RequiredIssueStatusIdsForKanbanProject()
	} else {
		requiredIssueStatusIds = constants.RequiredIssueStatusIdsForScrumProject()
	}

	requiredTitles := []string{}
	database.Db.Table("issue_statuses").Select("title").Where("id IN ?", requiredIssueStatusIds).Find(&requiredTitles)

	for _, issueStatusId := range requiredIssueStatusIds {
		if foundIdx := slices.IndexFunc(p.ProjectIssueStatuses, func(pis *ProjectIssueStatus) bool { return pis.IssueStatusId == issueStatusId }); foundIdx == -1 {
			return false, requiredTitles
		}
	}
	return true, requiredTitles
}

func (p *Project) BeforeUpdate(tx *gorm.DB) (err error) {
	if p.ProjectType == enums.ProjectTypeKanban {
		p.SprintDuration = nil
	}

	if tx.Statement.Changed() {
		tx.Statement.SetColumn("lock_version", p.LockVersion+1)

		if tx.Statement.Changed("state") {
			timeNow := time.Now()
			if p.State == enums.ProjectStateActive {
				tx.Statement.SetColumn("inactived_at", &timeNow)
				tx.Statement.SetColumn("actived_at", nil)
			} else {
				tx.Statement.SetColumn("actived_at", &timeNow)
				tx.Statement.SetColumn("inactived_at", nil)
			}
		}
	}

	return
}
