// Autonomo 1 Jorge Luis Rojas Robles - 2026

package service

import (
	"github.com/JorgeLuisRojasRobles/Autonomo-1/internal/adapter/epub"
	"github.com/JorgeLuisRojasRobles/Autonomo-1/internal/domain"
	"github.com/samber/lo"
	"github.com/samber/mo"
)

type LibraryService struct {
	repo domain.BookRepository
}

func NewLibraryService(repo domain.BookRepository) *LibraryService {
	return &LibraryService{repo: repo}
}

func (s *LibraryService) ImportBooks(paths []string) []domain.Book {

	// 1. Pipeline de Transformación
	results := lo.Map(paths, func(path string, _ int) mo.Result[domain.Book] {

		// A. Ejecutamos el Parser
		resMeta := epub.ParseMetadata(path)

		// B. Si falló, devolvemos el error inmediatamente (Vía Roja)
		if resMeta.IsError() {
			return mo.Err[domain.Book](resMeta.Error())
		}

		// C. Si tuvo éxito, extraemos la data y creamos el Libro (Vía Verde)
		meta := resMeta.MustGet()

		return domain.NewBook(
			domain.GenerateID(),
			meta.Title,
			meta.Author,
			100,
			path,
		)
	})

	// 2. Filtrado de Errores
	validBooks := lo.FilterMap(results, func(res mo.Result[domain.Book], _ int) (domain.Book, bool) {
		// Desempaquetamos manualmente para FilterMap
		val, err := res.Get()
		return val, err == nil
	})

	// 3. Persistencia (Usando la Interfaz)
	lo.ForEach(validBooks, func(book domain.Book, _ int) {
		s.repo.Save(book)
	})

	return validBooks
}

func (s *LibraryService) GetAllBooks() []domain.Book {
	return s.repo.ListAll()
}
