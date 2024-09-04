package dependencies

import (
	"log"
	"os"
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

func BuildMainDependencies() {
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

	restchallenges.Handler(server, challengeController)
	restVideo.Handler(server, videoController)
	restUser.Handler(server, userController)

	err := server.Run(os.Getenv("PORT"))

	if err != nil {
		log.Fatalln(err)
	}
}
