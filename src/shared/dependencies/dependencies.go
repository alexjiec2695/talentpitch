package dependencies

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
	domainchallenge "talentpitch/src/modules/challenges/domain"
	persistenceChallenge "talentpitch/src/modules/challenges/infra/persistence"
	restchallenges "talentpitch/src/modules/challenges/infra/rest"
	domainuser "talentpitch/src/modules/users/domain"
	persistenceUser "talentpitch/src/modules/users/infra/persistence"
	restUser "talentpitch/src/modules/users/infra/rest"
	domainvideos "talentpitch/src/modules/videos/domain"
	persistenceVideo "talentpitch/src/modules/videos/infra/persistence"
	restVideo "talentpitch/src/modules/videos/infra/rest"
	"talentpitch/src/shared/persistence"
	"talentpitch/src/shared/rest"
)

func BuildMainDependencies() *gin.Engine {
	server := rest.NewServer()
	db := persistence.InitDB()

	userRepository := persistenceUser.NewUserRepository(db)
	useCaseUsers := domainuser.NewUseCase(userRepository)
	userController := restUser.NewController(useCaseUsers)

	videoRepository := persistenceVideo.NewVideosRepository(db)
	useCaseVideo := domainvideos.NewUseCase(videoRepository)
	videoController := restVideo.NewController(useCaseVideo)

	challengeRepository := persistenceChallenge.NewChallengesRepository(db)
	useCaseChallenge := domainchallenge.NewUseCase(challengeRepository)
	challengeController := restchallenges.NewController(useCaseChallenge)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		userRepository.MassiveCreate()
	}()

	go func() {
		defer wg.Done()
		videoRepository.MassiveCreate()
	}()

	go func() {
		defer wg.Done()
		challengeRepository.MassiveCreate()
	}()
	fmt.Println("Generating information.......")
	wg.Wait()

	restchallenges.Handler(server, challengeController)
	restVideo.Handler(server, videoController)
	restUser.Handler(server, userController)

	return server
}
