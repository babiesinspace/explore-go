# Junior Backend Engineer Homework
Thanks for taking the time to work through the homework. While we don't expect you to come in with a knowledge of Go, we do value willingness and ability to get up to speed on the language, which is why this homework is Go-based.

## Getting started
1. Install and setup Go
2. Install sqlite if it is not already on your machine
3. Fork and clone this repo in the following location `$GOPATH/src/github.com/Hinge/backend-hw-junior`. e.g.:
```
cd $GOPATH/src/github.com/Hinge
git clone git@github.com:Hinge/backend-hw-junior.git
```
4. `go get github.com/mattn/go-sqlite3` and `go get github.com/go-chi/chi` to install dependencies
5. `cd backend-hw-junior`
6. `go run main.go`

## The assignment
Think of this project as a simplified version of our Hinge backend. You'll be writing API endpoints that iOS and Android apps could use. The API you are building should communicate with JSON.

1. Add a new endpoint (`/likes`) for a user to get their incoming likes (e.g., people who have already liked them).
2. Add a new endpoint (`/users/{id}`) for a user to edit their profile.

### Bonus
3. Add a new endpoint (`/ratings`) for a user to take action on other users. There are two different rating types `like` and `remove`.
  - If you like a recommendation, they should see an incoming like from you
  - If you like an incoming like, that should create a match
  - If you remove a like, you should never see that like again
  - If you remove a match, you should never see that match again
4. Add authentication

## Notes
This project comes bundled with a sqlite database to give you some starter tables and to allow you to have a fully-functioning app without worrying about deploying and connecting to an external database. Please be sure to commit and push the `.db` file so we can run your code properly.

## Submitting your work
Push your changes up to your fork. Code quality and completeness are the primary criteria we'll be evaluating and we expect to be able to run your code locally. Other than that, feel free to wow us!