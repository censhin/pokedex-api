# pokedex-api
RESTful API to provide an interface to MongoDB for Pokedex data.

## setup
You'll need to have a Go environment set up properly with a $GOPATH
environment variable.
```
go install
cp config.json.default config.json
```
Populate `config.json` with port data for the server, and host and
port data for the database.

## run
```
$GOPATH/bin/pokedex-api
```
