package repository

import (
	"ProblemMicro/proto"
	"context"
	"database/sql"
	"time"
)

type ProblemRepositoryServer interface {
	Create(ctx context.Context, problem *proto.Problem) (*proto.Problem, error)
	ReadByID(ctx context.Context, id int64) (*proto.Problem, error)
	ReadAll(ctx context.Context) ([]*proto.Problem, error)
	ReadByTypeID(ctx context.Context, typeID int32) ([]*proto.Problem, error)
	ReadByUserID(ctx context.Context, userID int64) ([]*proto.Problem, error)
	ReadBySolved(ctx context.Context, isSolved bool) ([]*proto.Problem, error)
	ReadByTimePeriod(ctx context.Context, start, end *proto.DateTime) ([]*proto.Problem, error)
	Update(ctx context.Context, problem *proto.Problem) (*proto.Problem, error)
	DeleteByID(ctx context.Context, id int64) error
	ReadTypeByID(ctx context.Context, id int32) (*proto.ProblemType, error)
	ReadTypeAll(ctx context.Context) ([]*proto.ProblemType, error)
	CreateSolution(ctx context.Context, problem *proto.Problem, solution *proto.Solution) (*proto.Solution, error)
	ReadSolution(ctx context.Context, problem *proto.Problem) (*proto.Solution, error)
}

type ProblemRepo struct {
	db *sql.DB
}

func NewProblemRepo(db *sql.DB) *ProblemRepo {
	return &ProblemRepo{db: db}
}

func (repo *ProblemRepo) Create(ctx context.Context, problem *proto.Problem) (*proto.Problem, error) {
	query := `INSERT INTO problems(user_id, type_Id, description, is_solved)
	VALUES($1, $2, $3, $4)
	RETURNING id, date_reported;`
	row := repo.db.QueryRowContext(ctx, query, problem.UserId, problem.Type.Id, problem.Description, problem.IsSolved)
	var currDate time.Time
	err := row.Scan(&problem.Id, &currDate)
	if err != nil {
		return problem, err
	}
	problem.ReportedAt = &proto.DateTime{
		Seconds: currDate.Unix(),
	}
	return problem, err
}

func (repo *ProblemRepo) ReadByID(ctx context.Context, id int64) (*proto.Problem, error) {
	query := `SELECT 
		probls.id, 
      	probls.user_id, 
       	probls.type_Id, 
       	types.name as typeName,
       	probls.description, 
       	probls.is_solved, 
       	probls.date_reported 
		FROM problems as probls 
		    LEFT JOIN problem_types as types 
		        ON probls.type_Id = types.id
		WHERE probls.id = $1;`
	row := repo.db.QueryRowContext(ctx, query, id)
	var problem proto.Problem
	var problemType proto.ProblemType
	var currDate time.Time
	err := row.Scan(&problem.Id, &problem.UserId, &problemType.Id,
		&problemType.Name, &problem.Description, &problem.IsSolved, &currDate)
	if err != nil {
		return &problem, err
	}
	problem.Type = &problemType
	problem.ReportedAt = &proto.DateTime{
		Seconds: currDate.Unix(),
	}
	return &problem, err
}

func (repo *ProblemRepo) readWithCondition(ctx context.Context, condition string, params ...interface{}) ([]*proto.Problem, error) {
	query := `SELECT 
		probls.id, 
      	probls.user_id, 
       	probls.type_Id, 
       	types.name as typeName,
       	probls.description, 
       	probls.is_solved, 
       	probls.date_reported 
		FROM problems as probls 
		    LEFT JOIN problem_types as types 
		        ON probls.type_Id = types.id
		`
	query += condition + `;`

	var result []*proto.Problem

	row, err := repo.db.QueryContext(ctx, query, params...)
	if err != nil {
		return result, err
	}
	defer row.Close()

	for row.Next() {
		var problem proto.Problem
		var problemType proto.ProblemType
		var currDate time.Time
		err := row.Scan(&problem.Id, &problem.UserId, &problemType.Id,
			&problemType.Name, &problem.Description, &problem.IsSolved, &currDate)
		if err != nil {
			return result, err
		}
		problem.Type = &problemType
		problem.ReportedAt = &proto.DateTime{
			Seconds: currDate.Unix(),
		}
		result = append(result, &problem)
	}
	return result, err
}

func (repo *ProblemRepo) ReadAll(ctx context.Context) ([]*proto.Problem, error) {
	return repo.readWithCondition(ctx, ``)
}

func (repo *ProblemRepo) ReadByTypeID(ctx context.Context, typeID int32) ([]*proto.Problem, error) {
	return repo.readWithCondition(ctx, `WHERE probls.type_Id = $1`, typeID)
}

func (repo *ProblemRepo) ReadByUserID(ctx context.Context, userID int64) ([]*proto.Problem, error) {
	return repo.readWithCondition(ctx, `WHERE probls.user_id = $1`, userID)
}

func (repo *ProblemRepo) ReadBySolved(ctx context.Context, isSolved bool) ([]*proto.Problem, error) {
	return repo.readWithCondition(ctx, `WHERE probls.is_solved = $1`, isSolved)
}

func (repo *ProblemRepo) ReadByTimePeriod(ctx context.Context, start, end *proto.DateTime) ([]*proto.Problem, error) {
	return repo.readWithCondition(ctx, `WHERE probls.date_reported >= $1 AND probls.date_reported <= $2`,
		time.Unix(start.Seconds, 0), time.Unix(end.Seconds, 0))
}

func (repo *ProblemRepo) Update(ctx context.Context, problem *proto.Problem) (*proto.Problem, error) {
	query := `UPDATE problems 
		SET 
      	user_id = $1, 
       	type_Id = $2, 
       	description = $3, 
       	is_solved = $4
		WHERE id = $5;`
	_, err := repo.db.ExecContext(ctx, query, problem.UserId, problem.Type.Id,
		problem.Description, problem.IsSolved, problem.Id)
	return problem, err
}

func (repo *ProblemRepo) DeleteByID(ctx context.Context, id int64) error {
	query := `DELETE FROM problems WHERE id = $1;`
	_, err := repo.db.ExecContext(ctx, query, id)
	return err
}

func (repo *ProblemRepo) ReadTypeByID(ctx context.Context, id int32) (*proto.ProblemType, error) {
	query := `SELECT 
		id, 
       	name
		FROM problem_types 
		WHERE id = $1;`
	row := repo.db.QueryRowContext(ctx, query, id)
	var problemType proto.ProblemType
	err := row.Scan(&problemType.Id, &problemType.Name)
	return &problemType, err
}

func (repo *ProblemRepo) ReadTypeAll(ctx context.Context) ([]*proto.ProblemType, error) {
	query := `SELECT 
		id, 
       	name
		FROM problem_types;`
	row, err := repo.db.QueryContext(ctx, query)
	var result []*proto.ProblemType
	if err != nil {
		return result, err
	}
	defer row.Close()
	for row.Next() {
		var problemType proto.ProblemType
		err := row.Scan(&problemType.Id, &problemType.Name)
		if err != nil {
			return result, err
		}
		result = append(result, &problemType)
	}
	return result, nil
}

func (repo *ProblemRepo) CreateSolution(ctx context.Context, problem *proto.Problem, solution *proto.Solution) (*proto.Solution, error) {
	query := `INSERT INTO 
		solutions(problem_id, description)
		VALUES($1, $2)
		RETURNING date_solved;`
	row := repo.db.QueryRowContext(ctx, query, problem.Id, solution.Description)
	var currDate time.Time
	err := row.Scan(&currDate)
	if err != nil {
		return solution, err
	}
	solution.SolvedAt = &proto.DateTime{
		Seconds: currDate.Unix(),
	}
	solution.Problem = problem
	return solution, err
}

func (repo *ProblemRepo) ReadSolution(ctx context.Context, problem *proto.Problem) (*proto.Solution, error) {
	query := `SELECT 
		description, 
      	date_solved
		FROM solutions 
		WHERE problem_id = $1;`
	row := repo.db.QueryRowContext(ctx, query, problem.Id)
	var solution proto.Solution
	var currDate time.Time
	err := row.Scan(&solution.Description, &currDate)
	if err != nil {
		return &solution, err
	}
	solution.Problem = problem
	solution.SolvedAt = &proto.DateTime{
		Seconds: currDate.Unix(),
	}
	return &solution, err
}
