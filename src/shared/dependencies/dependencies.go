package dependencies

import (
	"log"
	"os"
	"talentpitch/src/modules/users/domain"
	"talentpitch/src/modules/users/infra/persistence"
	r "talentpitch/src/modules/users/infra/rest"
	p "talentpitch/src/shared/persistence"
	"talentpitch/src/shared/rest"
)

func BuildMainDependencies() {
	server := rest.NewServer()
	db := p.InitDB()

	userRepository := persistence.NewUserRepository(db)
	useCaseUsers := domain.NewUseCase(userRepository)
	userController := r.NewController(useCaseUsers)

	r.Handler(server, userController)

	err := server.Run(os.Getenv("PORT"))

	if err != nil {
		log.Fatalln(err)
	}
}
