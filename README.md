# UserService

This is a sample user service exposed as REST APIs
- Actions "/user/create" => create a new user

## ORM Used
SQlBoiler (https://github.com/volatiletech/sqlboiler)

## How to Setup ?
- clone the repo
- set your GOPATH to local scope
- create and start mysql server
- edit the connection credentials in config.go
- apply the dump from here (https://drive.google.com/drive/folders/1FyVh7QJmBirPmVUw5hxxTbZEEAlipTcm?usp=sharing)


## generate models
- cd to repositories
- $PWD/sqlboiler --wipe --no-tests mysql
