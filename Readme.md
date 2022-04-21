## Running projects
open two terminal and cd to bellow paths:
`$ cd cmd/uss`
`$ cd cmd/es`
now first run Es and after that in Uss directory:
`$ go run .`

in Uss app you can input user and its segment and after that in Es app you can see estimate of each segment by enter segment name.

## Running projects with docker
`docker-compose up -d`
and its run two es and uss app. after running containers you should shell to them for using cli.

## make projects
in root directory of project run:
`make all`
of for building one of them run:
`make Uss/Es`

## Min-Heap
in fact min heap's are an array that keeping value in special order. by using min-heap we can access minimum value with O(1) time complexity but when a value want to inserted to the heap we need O(log(n)) time complexity.

it can help us to keep user ttl more easy and accessible. when we want to check expired users, we just need to check head of heap. if it was expired we should pop from heap until no expired user found.

## why Grpc?
in software with multiple part that interact with each other many solution can used to implement messaging between them. grpc is one of them that use http-2 and can provide urban or stream connection. Given that grpc combines with interfaces and free developers from consider routing of Rest-Api's, its a good choice for messaging between programs.

## Project structure
this project is based on Clean-Archtecture and directories retrieved from this repo:

https://github.com/golang-standards/project-layout

## note
In this project, we have tried to avoid additional libraries and focus on data structure's and clean architecture.