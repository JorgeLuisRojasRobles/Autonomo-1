package domain

import (
	"fmt"
	"time"

	"github.com/samber/mo"
)

// =========================================================
// 1. ESTRUCTURAS AUXILIARES (¡Esta era la que faltaba!)
// =========================================================
type EPUBMetadata struct {
	Title  string
	Author string
}

// =========================================================
// 2. MANEJO DE ERRORES (Unidad 3)
// =========================================================

type DomainError struct {
	Message string
	Code    int
}

func (e DomainError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func NewDomainError(msg string, code int) error {
	return DomainError{Message: msg, Code: code}
}

// =========================================================
// 3. INTERFACES (Unidad 3)
// =========================================================

type BookRepository interface {
	FindByID(id BookID) mo.Option[Book]
	Save(book Book) mo.Result[bool]
	ListAll() []Book
}

// =========================================================
// 4. DOMINIO PRINCIPAL
// =========================================================

type BookID string

type Book struct {
	ID          BookID
	Title       string
	Author      string
	Description mo.Option[string]
	PublishDate time.Time
	PageCount   int
	FilePath    string
	Format      string
}

func GenerateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// Constructor Seguro
func NewBook(id string, title string, author string, pages int, path string) mo.Result[Book] {
	if id == "" {
		return mo.Err[Book](NewDomainError("ID no puede estar vacío", 400))
	}
	if title == "" {
		return mo.Err[Book](NewDomainError("El título es obligatorio", 400))
	}
	if pages <= 0 {
		return mo.Err[Book](NewDomainError("El conteo de páginas debe ser positivo", 400))
	}

	return mo.Ok(Book{
		ID:          BookID(id),
		Title:       title,
		Author:      author,
		Description: mo.None[string](),
		PublishDate: time.Now(),
		PageCount:   pages,
		FilePath:    path,
		Format:      "EPUB",
	})
}

// Setter Funcional (Unidad 3)
func (b Book) WithTitle(newTitle string) Book {
	b.Title = newTitle
	return b
}

func (b Book) WithDescription(desc string) Book {
	b.Description = mo.Some(desc)
	return b
}
