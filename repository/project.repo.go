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
// ActiveEq        *string

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
			r.activeEq(query.ActiveEq),
		), paginateData),
	).Order("id desc").Find(&projects).Error
}

func (r *Repository) nameCont(nameLike *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if nameLike == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(projects.name) LIKE ?`, strings.ToLower("%"+*nameLike+"%")))
		}
	}
}

func (r *Repository) DescriptionCont(descriptionCont *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if descriptionCont == nil {
			return db
		} else {
			return db.Where(gorm.Expr(`lower(projects.description) LIKE ?`, strings.ToLower("%"+*descriptionCont+"%")))
		}
	}
}

func (r *Repository) projectTypeEq(projectTypeEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if projectTypeEq == nil {
			return db
		} else {
			projectTypeEqInInt, _ := enums.ParseProjectType(*projectTypeEq)

			return db.Where(gorm.Expr(`projects.project_type = ?`, projectTypeEqInInt))
		}
	}
}

func (r *Repository) activeEq(activeEq *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if activeEq == nil {
			return db
		} else {
			activeEqInInt, _ := enums.ParseProjectState(*activeEq)

			return db.Where(gorm.Expr(`projects.state = ?`, activeEqInInt))
		}
	}
}

func (r *ProjectRepository) Create(project *models.Project) error {
	return r.db.Model(&project).Preload("ProjectIssueStatuses.IssueStatus").Preload("ProjectAssignees").Create(&project).First(&project).Error
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
