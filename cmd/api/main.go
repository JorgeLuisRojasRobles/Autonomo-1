// Autonomo 1 Jorge Luis Rojas Robles - 2026
package main

import (
	"fmt"

	"github.com/JorgeLuisRojasRobles/Autonomo-1/internal/adapter/handler"
	"github.com/JorgeLuisRojasRobles/Autonomo-1/internal/adapter/repository" // Import nuevo
	"github.com/JorgeLuisRojasRobles/Autonomo-1/internal/service"
	"github.com/labstack/echo/v4"
)

func main() {
	// 1. Inicialización de Dependencias

	// A. Creamos el Repositorio
	repo := repository.NewInMemoryBookRepo()

	// B. Creamos el Servicio
	svc := service.NewLibraryService(repo)

	// C. Creamos el Handler
	bookHandler := handler.NewBookHandler(svc)

	// 2. Configuración del Servidor
	e := echo.New()

	// 3. Rutas
	e.POST("/books/import", bookHandler.Import)

	// Nueva ruta para ver qué tenemos guardado
	e.GET("/books", func(c echo.Context) error {
		books := svc.GetAllBooks()
		return c.JSON(200, books)
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Sistema Funcional de Gestión de Libros UIDE📚")
	})

	// 4. Arrancar
	fmt.Println("El servidor esta corriendo en el puerto 8080")
	e.Logger.Fatal(e.Start(":8080"))
}
