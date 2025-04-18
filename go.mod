module devops-toolbelt

go 1.23

require (
    github.com/docker/docker v20.10.23+incompatible
    github.com/gorilla/mux v1.8.0
)

replace github.com/docker/distribution => github.com/docker/distribution v2.7.1+incompatible
