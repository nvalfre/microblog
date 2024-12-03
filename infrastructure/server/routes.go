package server

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"microblog/application/usecases/tweet"
	"microblog/application/usecases/user"
	"microblog/infrastructure/adapters/cache"
	"microblog/infrastructure/adapters/persistence"
	"microblog/interface/controllers"
	"microblog/security/auth"
	"microblog/services"
)

// RegisterRoutes sets up the routes for the HTTP server.
func RegisterRoutes(router *gin.Engine, mongoClient *mongo.Client, redisCache *cache.RedisCache) {
	// Repositories
	userRepo := persistence.NewMongoUserTimelineRepository(mongoClient.Database("microblog").Collection("user_timeline"))
	tweetRepo := persistence.NewMongoTweetRepository(mongoClient.Database("microblog").Collection("tweets"))

	// Services
	userService := services.NewUserService(userRepo)
	tweetService := services.NewTweetService(tweetRepo)

	// Port use cases
	followUserUseCase := &user.FollowUserUseCase{UserRepo: userService}
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
		userApi := private.Group("/user")
		userApi.POST("/follow", userController.FollowUser)
		userApi.GET("/timeline", timelineController.GetTimeline)

		tweetApi := private.Group("/tweet")
		tweetApi.POST("/", tweetController.PublishTweet)
	}
}