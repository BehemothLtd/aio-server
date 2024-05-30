package repository

import (
	"aio-server/enums"
	"aio-server/gql/inputs/insightInputs"
	"aio-server/models"
	"aio-server/pkg/helpers"
	"context"
	"strings"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	Repository
}

func NewProjectRepository(c *context.Context, db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		Repository: Repository{
			db:  db,
			ctx: c,
		},
	}
}

// Find finds a project by its attributes.
func (r *ProjectRepository) Find(project *models.Project) error {
	dbTables := r.db.Model(&models.Project{})

	return dbTables.Where(&project).First(&project).Error
}

// NameCont        *string
// DescriptionCont *string
// ProjectTypeEq   *string
// StateEq        *string

func (r *ProjectRepository) List(
	projects *[]*models.Project,
	paginateData *models.PaginationData,
	query insightInputs.ProjectsQueryInput,
) error {
	dbTables := r.db.Model(&models.Project{})

	return dbTables.Scopes(
		helpers.Paginate(dbTables.Scopes(
			r.nameCont(query.NameCont),
			r.DescriptionCont(query.DescriptionCont),
			r.projectTypeEq(query.ProjectTypeEq),
			r.stateEq(query.StateEq),
		), paginateData),
	).Order("id desc").Find(&projects).Error
}

func (r *ProjectRepository) nameCont(nameLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nameLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(projects.name) LIKE ?`, strings.ToLower("%"+*nameLike+"%")))
		}
	}
}

func (r *ProjectRepository) DescriptionCont(descriptionCont *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if descriptionCont == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(projects.description) LIKE ?`, strings.ToLower("%"+*descriptionCont+"%")))
		}
	}
}

func (r *ProjectRepository) projectTypeEq(projectTypeEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if projectTypeEq == nil {
			return db
		} else {
			projectTypeEqInInt, _ := enums.ParseProjectType(*projectTypeEq)

			return db.Where(gorm.Expr(`projects.project_type = ?`, projectTypeEqInInt))
		}
	}
}

func (r *ProjectRepository) stateEq(stateEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if stateEq == nil {
			return db
		} else {
			stateEqInInt, _ := enums.ParseProjectState(*stateEq)

			return db.Where(gorm.Expr(`projects.state = ?`, stateEqInInt))
		}
	}
}

func (r *ProjectRepository) Create(project *models.Project) error {
	return r.db.Model(&project).Preload("ProjectIssueStatuses.IssueStatus").Preload("ProjectAssignees").Create(&project).First(&project).Error
}

func (r *ProjectRepository) UpdateBasicInfo(project *models.Project, updates map[string]interface{}) error {
	if err := r.db.Model(&project).Select(append(helpers.GetKeys(updates), "LockVersion")).Updates(updates).Error; err != nil {
		return err
	}

	return r.db.Model(&models.Project{}).
		Where("id = ?", project.Id).
		Preload("Client").
		Preload("Logo", "name = 'logo'").Preload("Logo.AttachmentBlob").
		Preload("Files", "name = 'files'").Preload("Files.AttachmentBlob").
		Find(&project).Error
}

func (r *ProjectRepository) Update(project *models.Project, updateProject models.Project) error {
	if err := r.db.Model(&project).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&updateProject).Error; err != nil {
		return err
	}

	return r.db.Model(&project).Preload("ProjectIssueStatuses.IssueStatus").Preload("ProjectAssignees").Where("id = ?", project.Id).First(&project).Error
}

func (r *ProjectRepository) UpdateFiles(project *models.Project) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		if project.Logo != nil {
			if err := r.db.Model(&models.Project{Id: project.Id}).Unscoped().Where("name = 'logo'").Association("Logo").Unscoped().Clear(); err != nil {
				return err
			}
		}

		if len(project.Files) > 0 {
			if err := r.db.Model(&models.Project{Id: project.Id}).Unscoped().Where("name = 'files'").Association("Files").Unscoped().Clear(); err != nil {
				return err
			}
		}

		return r.db.Model(&project).Updates(&project).Error
	}); err != nil {
		return err
	}

	return r.db.Model(&project).Where("id = ?", project.Id).
		Preload("Logo", "name = 'logo'").Preload("Logo.AttachmentBlob").
		Preload("Files", "name = 'files'").Preload("Files.AttachmentBlob").
		First(&project).Error
}

func (r *ProjectRepository) ActiveHighPriorityProjects(projects *[]models.Project) error {
	return r.db.Model(&models.Project{}).
		Where(
			"state = ? and project_priority = ?",
			enums.ProjectStateActive, enums.ProjectPriorityHigh,
		).
		Select("id, name").
		Order("id desc").
		Scan(&projects).
		Error
}

func (r *ProjectRepository) All(projects *[]*models.Project) error {
	return r.db.Table("projects").Order("id ASC").Find(&projects).Error
}

func (r *ProjectRepository) ListProjectByUser(projects *[]*models.Project, userId int32) error {
	return r.db.Table("projects").Joins("LEFT JOIN project_assignees ON projects.id = project_assignees.project_id").
		Where("project_assignees.user_id = ?", userId).
		Group("projects.id").
		Find(&projects).
		Error
}

func (r *ProjectRepository) Delete(project *models.Project) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.ProjectAssignee{}, "project_id = ?", project.Id).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.ProjectIssueStatus{}, "project_id = ?", project.Id).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.ProjectSprint{}, "project_id = ?", project.Id).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.IssueAssignee{}, "id IN (SELECT id FROM issues WHERE project_id = ?)", project.Id).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.Issue{}, "project_id = ?", project.Id).Error; err != nil {
			return err
		}

		if err := tx.Delete(&models.Project{}, "id = ?", project.Id).Error; err != nil {
			return err
		}

		// return nil will commit the whole transaction
		return nil
	})

	return err
}

func (r *ProjectRepository) ActiveSprint(project *models.Project, projectSprint models.ProjectSprint) error {
	return r.db.Model(&project).Update("current_sprint_id", projectSprint.Id).Error
}

func (r *ProjectRepository) ChangeActiveSprint(project *models.Project, sprint *models.ProjectSprint, moveToSprint models.ProjectSprint) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&project).Update("current_sprint_id", moveToSprint.Id).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.Issue{}).
			Where("project_id = ? AND project_sprint_id = ?", project.Id, sprint.Id).
			Update("project_sprint_id", moveToSprint.Id).Error; err != nil {
			return err
		}

		if err := tx.Delete(&sprint).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}
