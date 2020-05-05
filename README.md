# spar

Cloud Spanner DDL parser in Go base on syucream's work. New version support one kind of comment on it because we need to use this library to generate model json file automatically and the comments can be used as the description of fields. 

For now, only this kind of comments is accepted : 

``` sql
CREATE TABLE Singers (
  SingerId   INT64 NOT NULL, /* id of the singer */
  FirstName  STRING(1024),
  LastName   STRING(1024),
  SingerInfo BYTES(MAX),
  BirthDate  DATE,
) PRIMARY KEY (SingerId);
```

## Usage

- Pass your Spanner DDL to `parser.Parse()` and get parsed statements.
- See detail on cmd/spck/spck.go

## subtool

`spck` is a Spanner DDL syntax checker. It parses and returns exit code 0 if no error.

```
$ go get -u github.com/CuberL/spar/cmd/spck
$ spck ./examples/create_table.sql
```
