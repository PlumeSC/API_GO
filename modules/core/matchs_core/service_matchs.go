package matchscore

type MatchsService interface {
	GetMatchs() error
}

type matchsService struct {
	repo MatchsRepository
}

func NewMatchService(repo MatchsRepository) *matchsService {
	return &matchsService{repo: repo}
}

func (s matchsService) GetMatchs() error {

	return nil
}
