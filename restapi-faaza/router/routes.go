package router

import (
    "github.com/gin-gonic/gin"
    "path/to/your/controllers" // Import package yang berisi controller-controller Anda
)

func SetupRouter() *gin.Engine {
    // Inisialisasi router
    r := gin.Default()

    // Contoh endpoint untuk GET
    r.GET("/users/:userId", controllers.GetUserHandler)

    // Contoh endpoint untuk POST
    r.POST("/users/register", controllers.RegisterUserHandler)

    // Contoh endpoint untuk PUT
    r.PUT("/users/:userId", controllers.UpdateUserHandler)

    // Contoh endpoint untuk DELETE
    r.DELETE("/users/:userId", controllers.DeleteUserHandler)

    // Tambahkan rute lainnya sesuai kebutuhan Anda

    return r
}
