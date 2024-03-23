package router

import "github.com/gin-gonic/gin"

type Router = *gin.Engine

type Context = *gin.Context

func NewRouter() Router {
	router := gin.Default()
	return router
}
