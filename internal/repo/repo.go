package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/ozoncp/ocp-solution-api/internal/models"

	"github.com/jmoiron/sqlx"
)

const (
	solutionsTableName = "solutions"
	verdictsTableName  = "verdicts"
)

var (
	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

type Repo interface {
	AddSolution(ctx context.Context, issueId uint64) (*models.Solution, error)
	AddSolutions(ctx context.Context, issueIds []uint64) error
	RemoveSolution(ctx context.Context, solutionId uint64) error
	UpdateSolution(ctx context.Context, solution models.Solution) error
	ListSolutions(ctx context.Context, limit, offset uint64) ([]*models.Solution, error)
}

func NewRepo(db sqlx.DB) Repo {
	return &repo{db: db}
}

type repo struct {
	db sqlx.DB
}

func (r *repo) AddSolution(ctx context.Context, issueId uint64) (*models.Solution, error) {
	query := psql.Insert(solutionsTableName).
		Columns("issue_id").
		Values(issueId).
		Suffix(`RETURNING "id"`).
		RunWith(r.db)

	var solutionId uint64 = 0
	err := query.QueryRowContext(ctx).Scan(&solutionId)

	return models.NewSolution(solutionId, issueId), err
}

func (r *repo) AddSolutions(ctx context.Context, issueIds []uint64) error {
	query := psql.Insert(solutionsTableName).
		Columns("issue_id").
		Suffix(`RETURNING "id"`).
		RunWith(r.db)

	for _, issueId := range issueIds {
		query = query.Values(issueId)
	}

	_, err := query.ExecContext(ctx)
	return err
}

func (r *repo) RemoveSolution(ctx context.Context, solutionId uint64) error {
	query := psql.Delete(solutionsTableName).
		Where(sq.Eq{"id": solutionId}).
		RunWith(r.db)

	// TODO: remove verdict

	_, err := query.ExecContext(ctx)
	return err
}

func (r *repo) UpdateSolution(ctx context.Context, solution models.Solution) error {
	query := psql.Update(solutionsTableName).
		Set("issue_id", solution.IssueId()).
		Where(sq.Eq{"id": solution.Id()}).
		RunWith(r.db)

	// TODO: update verdict

	_, err := query.ExecContext(ctx)
	return err
}

func (r *repo) ListSolutions(ctx context.Context, limit, offset uint64) ([]*models.Solution, error) {
	query := psql.Select("id", "issueId").
		From(solutionsTableName).
		RunWith(r.db).
		Limit(limit).
		Offset(offset)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var solutions []*models.Solution
	for rows.Next() {
		var solutionId, issueId uint64
		if err := rows.Scan(&solutionId, &issueId); err != nil {
			continue
		}
		solutions = append(solutions, models.NewSolution(solutionId, issueId))
		// TODO: fetch verdicts
	}
	return solutions, nil
}
