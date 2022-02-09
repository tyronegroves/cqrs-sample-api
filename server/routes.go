package server

func (s *server) LoadRoutes() {
	r := s.router.Group("/api")
	r.Use(s.authorizeUser())
	r.POST("/register_user", s.handleRegisterUser())
	r.POST("/release_movie", s.handleReleaseMovie())
	r.POST("/rate_movie", s.handleRateMovie())
	r.GET("/movie", s.handleGetMovie())
	r.GET("/recommendations", s.handleGetRecommendations())
	r.GET("/ratings", s.handleGetRatings())
	r.GET("/user_ratings", s.handleGetUserRatings())
}
