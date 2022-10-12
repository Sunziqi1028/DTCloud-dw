/**
 * @Author: Sun
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2022/6/15 16:46
 */

package router

import (
	"github.com/gin-gonic/gin"
	"wg315/controller"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/device/WG024YX", controller.WG024YX{}.HandData)
	router.POST("")

	return router
}
