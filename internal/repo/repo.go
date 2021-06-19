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

type Repo interface {
	AddSolution(ctx context.Context, solution models.Solution) (*models.Solution, error)
	AddSolutions(ctx context.Context, solutions []models.Solution) error
	RemoveSolution(ctx context.Context, solutionId uint64) error
	UpdateSolution(ctx context.Context, solution models.Solution) error
	ListSolutions(ctx context.Context, limit, offset uint64) ([]*models.Solution, error)

	AddVerdict(ctx context.Context, verdict models.Verdict) (*models.Verdict, error)
	AddVerdicts(ctx context.Context, verdicts []models.Verdict) error
	RemoveVerdict(ctx context.Context, solutionId uint64) error
	UpdateVerdict(ctx context.Context, verdict models.Verdict) error
	ListVerdicts(ctx context.Context, limit, offset uint64) ([]*models.Verdict, error)
}

func NewRepo(db sqlx.DB) Repo {
	return &repo{db: db}
}

type repo struct {
	db sqlx.DB
}

func (r *repo) AddSolution(ctx context.Context, solution models.Solution) (*models.Solution, error) {
	query := sq.Insert(solutionsTableName).
		Columns("issue_id").
		Values(solution.IssueId()).
		Suffix(`RETURNING "id"`).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	var solutionId uint64 = 0
	err := query.QueryRowContext(ctx).Scan(&solutionId)

	return models.NewSolution(solutionId, solution.IssueId()), err
}

func (r *repo) AddSolutions(ctx context.Context, solutions []models.Solution) error {
	query := sq.Insert(solutionsTableName).
		Columns("issue_id").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	for _, solution := range solutions {
		query = query.Values(solution.IssueId())
	}

	_, err := query.ExecContext(ctx)
	return err
}

func (r *repo) RemoveSolution(ctx context.Context, solutionId uint64) error {
	query := sq.Delete(solutionsTableName).
		Where(sq.Eq{"id": solutionId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	_, err := query.ExecContext(ctx)

	if err == nil {
		// remove corresponding verdict
		query = sq.Delete(verdictsTableName).
			Where(sq.Eq{"solution_id": solutionId}).
			RunWith(r.db).
			PlaceholderFormat(sq.Dollar)

		_, err = query.ExecContext(ctx)
	}

	return err
}

func (r *repo) UpdateSolution(ctx context.Context, solution models.Solution) error {
	query := sq.Update(solutionsTableName).
		Set("issue_id", solution.IssueId()).
		Where(sq.Eq{"id": solution.Id()}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	_, err := query.ExecContext(ctx)
	return err
}

func (r *repo) ListSolutions(ctx context.Context, limit, offset uint64) ([]*models.Solution, error) {
	query := sq.Select("id", "issue_id").
		From(solutionsTableName).
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar)

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
	}
	return solutions, nil
}

func (r *repo) AddVerdict(ctx context.Context, verdict models.Verdict) (*models.Verdict, error) {
	query := sq.Insert(verdictsTableName).
		Columns("solution_id").
		Values(verdict.SolutionId()).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	_, err := query.ExecContext(ctx)

	return &verdict, err
}

func (r *repo) AddVerdicts(ctx context.Context, verdicts []models.Verdict) error {
	query := sq.Insert(verdictsTableName).
		Columns("solution_id").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	for _, verdict := range verdicts {
		query = query.Values(verdict.SolutionId())
	}

	_, err := query.ExecContext(ctx)
	return err
}

func (r *repo) RemoveVerdict(ctx context.Context, solutionId uint64) error {
	query := sq.Delete(verdictsTableName).
		Where(sq.Eq{"solution_id": solutionId}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	_, err := query.ExecContext(ctx)

	return err
}

func (r *repo) UpdateVerdict(ctx context.Context, verdict models.Verdict) error {
	status, comment, userId, timestamp := verdict.Status()
	query := sq.Update(verdictsTableName).
		Set("status", status).
		Set("comment", comment).
		Set("user_id", userId).
		Set("timestamp", timestamp).
		Where(sq.Eq{"solution_id": verdict.SolutionId()}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	_, err := query.ExecContext(ctx)
	return err
}

func (r *repo) ListVerdicts(ctx context.Context, limit, offset uint64) ([]*models.Verdict, error) {
	query := sq.Select("solution_id", "user_id", "status", "comment", "timestamp").
		From(verdictsTableName).
		RunWith(r.db).
		Limit(limit).
		Offset(offset).
		PlaceholderFormat(sq.Dollar)

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var verdicts []*models.Verdict
	for rows.Next() {
		var solutionId, userId uint64
		var status models.Status
		var comment string
		var timestamp int64
		if err := rows.Scan(&solutionId, &userId, &status, &comment, &timestamp); err != nil {
			continue
		}
		verdict := models.NewVerdict(solutionId, userId, status, comment)
		verdict.ForceTimestamp(timestamp)
		verdicts = append(verdicts, verdict)
	}
	return verdicts, nil
}
