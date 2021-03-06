package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"playlist-saver/controller/playlistCtrl"
	"playlist-saver/controller/playlistDetailCtrl"
	"playlist-saver/controller/searchCtrl"
	"playlist-saver/controller/statusCtrl"
	"playlist-saver/controller/tokenCtrl"
	"playlist-saver/controller/userCtrl"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserController     userCtrl.UserController
	SearchController   searchCtrl.SearchController
	PlaylistController playlistCtrl.PlaylistController
	DetailController   playlistDetailCtrl.PlaylistDetailController
	TokenController    tokenCtrl.TokenController
	StatusController   statusCtrl.StatusController
}

func (c1 *ControllerList) Registration(e *echo.Echo) {

	apiV1 := e.Group("/api/v1")

	//	USER THINGS
	apiV1.POST("/users/register", c1.UserController.Register)
	apiV1.POST("/users/login", c1.UserController.Login)
	apiV1.GET("/users", c1.UserController.GetAllUser, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.PUT("/users", c1.UserController.UpdateUser, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.PUT("/users/token", c1.UserController.AddTokenUser, middleware.JWTWithConfig(c1.JWTMiddleware))

	//	SEARCH
	apiV1.GET("/search", c1.SearchController.SearchYtByParam, middleware.JWTWithConfig(c1.JWTMiddleware))

	//PLAYLIST
	apiV1.POST("/playlists", c1.PlaylistController.CreatePlaylist, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.GET("/playlists", c1.PlaylistController.GetAllPlaylist, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.GET("/playlists/:id", c1.PlaylistController.GetPlaylistByUserId, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.DELETE("/playlists/:id", c1.PlaylistController.DeletePlaylist, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.PUT("/playlists/:id", c1.PlaylistController.EditPlaylist, middleware.JWTWithConfig(c1.JWTMiddleware))

	//ADD YOUTUBE TO PLAYLIST
	apiV1.POST("/playlist/:playlistId/data", c1.DetailController.AddYoutubeToPlaylist, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.DELETE("/playlists/:details/data", c1.DetailController.DeleteYoutubeDataFromPlaylist, middleware.JWTWithConfig(c1.JWTMiddleware))

	//ADMIN TOKEN
	apiV1.POST("/admins/gacha/token", c1.TokenController.AddToken, middleware.JWTWithConfig(c1.JWTMiddleware))
	apiV1.GET("/admins/gacha/token", c1.TokenController.GetToken, middleware.JWTWithConfig(c1.JWTMiddleware))

	//cron check
	apiV1.PUT("/checks/cron", c1.StatusController.CronCheckToken)

}
