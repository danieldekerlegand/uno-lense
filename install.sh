# install go

# install sqlite3
go get github.com/mattn/go-sqlite3

# set up database
sqlite3 lense.db
cat data/setup.sql | sqlite3 lense.db

# install docker api
go get github.com/docker/docker/client
