module github.com/sausheong/go-recipes/modules

go 1.14

require (
    github.com/sausheong/go-recipes/modules/app v0.0.0
    github.com/gorilla/mux v1.8.0
)

replace github.com/sausheong/go-recipes/modules/app => ./app