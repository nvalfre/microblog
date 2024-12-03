package server

import (
	"github.com/gin-gonic/gin"
	"microblog/application/usecases/tweet"
	"microblog/application/usecases/user"
	"microblog/infrastructure/adapters/cache"
	"microblog/infrastructure/adapters/persistence"
	"microblog/interface/controllers"
	"microblog/security/auth"
	"microblog/services"
)

// RegisterRoutes sets up the routes for the HTTP server.
func RegisterRoutes(router *gin.Engine, mongoUserTimelineRepository *persistence.MongoUserTimelineRepository, mongoTweetRepository *persistence.MongoTweetRepository, redisCache *cache.RedisCache) {
	// Services
	userService := services.NewUserService(mongoUserTimelineRepository)
	tweetService := services.NewTweetService(mongoTweetRepository)

	// Port use cases
	followUserUseCase := &user.FollowUserUseCase{UserService: userService}
	getTimelineUseCase := &tweet.GetTimelineUseCase{TweetService: tweetService, UserServuce: userService, Cache: redisCache}
	publishTweetUseCase := &tweet.PublishTweetUseCase{TweetService: tweetService}

	// Controllers
	userController := controllers.UserController{FollowUserUseCase: followUserUseCase}
	timelineController := controllers.TimelineController{GetTimelineUseCase: getTimelineUseCase}
	tweetController := controllers.TweetController{PublishTweetUseCase: publishTweetUseCase}
	authController := controllers.AuthController{}

	// Routes
	// Public Routes: No authentication required.
	public := router.Group("/")
	{
		public.GET("/generate_token", authController.GetToken)
	}

	// Private Routes: Require authentication.
	private := router.Group("/api")
	private.Use(auth.JwtSignedTokenMiddleware())
	{
		userApi := private.Group("/userCollection")
		userApi.POST("/follow", userController.FollowUser)
		userApi.GET("/timeline", timelineController.GetTimeline)

		tweetApi := private.Group("/tweet")
		tweetApi.POST("/", tweetController.PublishTweet)
	}
}
