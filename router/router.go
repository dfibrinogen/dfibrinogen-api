package router

import (
	"github.com/dafian47/dfibrinogen-api/config"
	"github.com/dafian47/dfibrinogen-api/controller"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/secure"
)

func InitRouter(db *gorm.DB) *gin.Engine {

	secureMiddleware := secure.New(secure.Options{
		FrameDeny:          true,
		ContentTypeNosniff: true,
		BrowserXssFilter:   true,
		IsDevelopment:      config.IsProduction,
	})

	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {

			err := secureMiddleware.Process(c.Writer, c.Request)

			// If there was an error, do not continue.
			if err != nil {
				c.Abort()
				return
			}

			// Avoid header rewrite if response is a redirection.
			if status := c.Writer.Status(); status > 300 && status < 399 {
				c.Abort()
			}
		}
	}()

	// Set DebugMode if you want to enable Log on Rest Server ( Gin )
	// And set ReleaseMode if you want to deploy to Production
	if config.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(secureFunc)
	router.MaxMultipartMemory = 8 << 20
	router.Static("/image", "./resource/images")

	baseController := controller.BaseController{DB: db}

	authV1 := router.Group("/auth")
	{
		authV1.POST("/login", baseController.UserLogin)
		authV1.POST("/register", baseController.UserRegister)
	}

	apiV1 := router.Group("/api/v1")
	{
		userRoute := apiV1.Group("/users")
		{
			userRoute.GET("/", baseController.GetUserAll)
			userRoute.DELETE("/:id", baseController.DeleteUser)
		}

		profileRoute := apiV1.Group("/profiles")
		{
			profileRoute.GET("/:id", baseController.GetProfileByID)
			profileRoute.PUT("/:id", baseController.UpdateProfile)
		}

		categoryRoute := apiV1.Group("/categories")
		{
			categoryRoute.GET("/", baseController.GetCategoryAll)
			categoryRoute.GET("/:id", baseController.GetCategoryByID)
			categoryRoute.POST("/", baseController.AddCategory)
			categoryRoute.PUT("/:id", baseController.UpdateCategory)
			categoryRoute.DELETE("/:id", baseController.DeleteCategory)
		}

		postRoute := apiV1.Group("/posts")
		{
			postRoute.GET("/", baseController.GetPostAll)
			postRoute.GET("/:id", baseController.GetPostByID)
			postRoute.POST("/", baseController.AddPost)
			postRoute.PUT("/:id", baseController.UpdatePost)
			postRoute.DELETE("/:id", baseController.DeletePost)
		}
	}

	return router
}
