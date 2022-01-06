package service

import (
	"ProblemMicro/proto"
	"ProblemMicro/repository"
	"context"
	"database/sql"
)

type ProblemService struct {
	Repo *repository.ProblemRepo
	*proto.UnimplementedProblemServiceServer
}

func NewProblemService(repo *sql.DB) *ProblemService {
	return &ProblemService{
		Repo: repository.NewProblemRepo(repo),
	}
}

func (serv *ProblemService) AddNewProblem(ctx context.Context, problem *proto.Problem) (*proto.Response, error) {
	problemCreated, err := serv.Repo.Create(ctx, problem)
	return &proto.Response{
		Success: err == nil,
		Problem: problemCreated,
	}, err
}

func (serv *ProblemService) UpdateProblem(ctx context.Context, problem *proto.Problem) (*proto.Response, error) {
	problemUpdated, err := serv.Repo.Update(ctx, problem)
	return &proto.Response{
		Success: err == nil,
		Problem: problemUpdated,
	}, err
}

func (serv *ProblemService) DeleteProblem(ctx context.Context, problem *proto.Problem) (*proto.Response, error) {
	err := serv.Repo.DeleteByID(ctx, problem.Id)
	return &proto.Response{Success: err == nil}, err
}

func (serv *ProblemService) GetProblemByID(ctx context.Context, request *proto.ProblemRequest) (*proto.Response, error) {
	problem, err := serv.Repo.ReadByID(ctx, request.Id)
	return &proto.Response{
		Success: err == nil,
		Problem: problem,
	}, err
}

func (serv *ProblemService) GetAllProblems(ctx context.Context, request *proto.ProblemRequest) (*proto.Response, error) {
	_ = request
	problems, err := serv.Repo.ReadAll(ctx)
	return &proto.Response{
		Success:  err == nil,
		Problems: problems,
	}, err
}

func (serv *ProblemService) GetProblemsByTypeID(ctx context.Context, request *proto.ProblemRequest) (*proto.Response, error) {
	problems, err := serv.Repo.ReadByTypeID(ctx, request.TypeId)
	return &proto.Response{
		Success:  err == nil,
		Problems: problems,
	}, err
}

func (serv *ProblemService) GetProblemsByUserID(ctx context.Context, request *proto.ProblemRequest) (*proto.Response, error) {
	problems, err := serv.Repo.ReadByUserID(ctx, request.UserId)
	return &proto.Response{
		Success:  err == nil,
		Problems: problems,
	}, err
}

func (serv *ProblemService) GetProblemsBySolved(ctx context.Context, request *proto.ProblemRequest) (*proto.Response, error) {
	problems, err := serv.Repo.ReadBySolved(ctx, request.IsSolved)
	return &proto.Response{
		Success:  err == nil,
		Problems: problems,
	}, err
}

func (serv *ProblemService) GetProblemsByTimePeriod(ctx context.Context, request *proto.ProblemRequest) (*proto.Response, error) {
	problems, err := serv.Repo.ReadByTimePeriod(ctx, request.StartTime, request.EndTime)
	return &proto.Response{
		Success:  err == nil,
		Problems: problems,
	}, err
}

func (serv *ProblemService) GetProblemTypeByID(ctx context.Context, request *proto.ProblemRequest) (*proto.Response, error) {
	problemType, err := serv.Repo.ReadTypeByID(ctx, request.TypeId)
	return &proto.Response{
		Success:     err == nil,
		ProblemType: problemType,
	}, err
}

func (serv *ProblemService) GetAllProblemTypes(ctx context.Context, request *proto.ProblemRequest) (*proto.Response, error) {
	_ = request
	problemTypes, err := serv.Repo.ReadTypeAll(ctx)
	return &proto.Response{
		Success:      err == nil,
		ProblemTypes: problemTypes,
	}, err
}

func (serv *ProblemService) AddProblemSolution(ctx context.Context, request *proto.ProblemSolution) (*proto.Response, error) {
	var err error
	var solutionCreated *proto.Solution

	request.Problem, err = serv.Repo.ReadByID(ctx, request.Problem.Id)
	if err != nil {
		return &proto.Response{
			Success: false,
		}, err
	}

	solutionCreated, err = serv.Repo.CreateSolution(ctx, request.Problem, request.Solution)

	if err == nil {
		request.Problem.IsSolved = true
		_, err = serv.UpdateProblem(ctx, request.Problem)
	}

	return &proto.Response{
		Success:  err == nil,
		Solution: solutionCreated,
	}, err
}

func (serv *ProblemService) GetSolutionByProblem(ctx context.Context, request *proto.Problem) (*proto.Response, error) {
	solutionFound, err := serv.Repo.ReadSolution(ctx, request)
	return &proto.Response{
		Success:  err == nil,
		Solution: solutionFound,
	}, err
}
