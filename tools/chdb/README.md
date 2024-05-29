# chdb-go
> https://github.com/chdb-io/chdb-go


## Install

### Install libchdb.so

```sh
sudo curl -sL https://lib.chdb.io | bash
```

### Install chdb-go

```sh

go install github.com/chdb-io/chdb-go@latest
```

## Example

### Simple mode

```sh
chdb-go 'SELECT * from file("sample1.parquet")'
```
output:
```
"first"," last"
"Jorge","Frank"
"Hunter","Moreno"
"Esther","Guzman"
"Dennis","Stephens"
"Nettie","Franklin"
"Stanley","Gibson"
"Eugenia","Greer"
...

```