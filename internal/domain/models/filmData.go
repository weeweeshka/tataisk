package models

type FilmData struct {
	Id           int32    `validate:"-"`
	Title        string   `validate:"required"`
	YearOfProd   uint32   `validate:"required,min=1900"`
	Imdb         float32  `validate:"required,min=0,max=10"`
	Description  string   `validate:"required"`
	Country      []string `validate:"omitempty"`
	Genre        []string `validate:"required"`
	FilmDirector string   `validate:"required"`
	Screenwriter string   `validate:"required"`
	Budget       int64    `validate:"required"`
	Collection   int64    `validate:"required"`
}

type FilmDataWithID struct {
	Id           int32    `validate:"required"`
	Title        string   `validate:"required"`
	YearOfProd   uint32   `validate:"required,min=1900"`
	Imdb         float32  `validate:"required,min=0,max=10"`
	Description  string   `validate:"required"`
	Country      []string `validate:"omitempty"`
	Genre        []string `validate:"required"`
	FilmDirector string   `validate:"required"`
	Screenwriter string   `validate:"required"`
	Budget       int64    `validate:"required"`
	Collection   int64    `validate:"required"`
}
