# Cueball

This project aims to corrupt the data of all files in a defined folder (behind stages it deletes the content, but no the file).

## How to build?

This command will generate a binary file to execute wherever and whenever you want:

```shell
go build ./src/main.go
```

## How to use the `test` folder?

With the following commands you will be able to see the program run in action with some files to test:

```
# In the root directory let's run the following
cd test/
go build ../src/main.go
./main # Or ./main.exe
```

Then you will be able to see how the internal data was removed