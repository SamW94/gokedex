module github.com/SamW94/gokedex

go 1.23.5

replace github.com/SamW94/gokedex/internal/pokeapi => ./internal/pokeapi

replace github.com/SamW94/gokedex/internal/pokecache => ./internal/pokecache

require github.com/SamW94/gokedex/internal/pokeapi v0.0.0

require github.com/SamW94/gokedex/internal/pokecache v0.0.0 // indirect
