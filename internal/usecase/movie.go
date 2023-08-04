package usecase

type MovieUseCase struct {
	movieRepo MovieElasticSearchRepo
}

// New -.
func NewMovie(m MovieElasticSearchRepo) *MovieUseCase {
	return &MovieUseCase{
		movieRepo: m,
	}
}
