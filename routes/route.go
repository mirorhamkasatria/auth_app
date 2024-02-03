package routes

import (
	con "github.com/auth_app/app/controllers"
	repo "github.com/auth_app/app/repositories"
	srv "github.com/auth_app/app/services"
)

func registerRouters(s *ApiServer) {
	ca := con.NewControllerUser(
		srv.NewServiceUser(
			repo.NewRepoUser(
				s.GormDB)),
		s.Validator)

	apiv1 := s.App.Group("/api/v1")

	users := apiv1.Group("/users")
	users.Post("/login", ca.LoginUser)
	users.Post("/register", ca.RegisterUser)

}
