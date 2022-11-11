# Pepperstone


```
git clone  https://github.com/joycezemitchell/pepperstone.git
```


## Run binary
```
bin/pepperstonecsv
```


## Makefile commands

- `make build` build the binary
- `make linux` build the binary in linux format to be used in Docker container
- `make run` test run the app

##  Note

```
The script will read all csv files from csv/input directory.
Csv file with xxxx_ctt.csv naming convention will be process as Cash Transfer Type. 
The script will gererated a csv file and save it in /csv/output/ directory.

The script will be using a temporary database stored in domain/account/datasouce
Here are the data.    

    []Account{
		{
			Id:19800,
			Balance: 500,
			Currency: "AUD",
		},
		{
			Id:19801,
			Balance: 500,
			Currency: "AUD",
		},
		{
			Id:12350,
			Balance: 500,
			Currency: "AUD",
		},
		{
			Id:12356,
			Balance: 500,
			Currency: "AUD",
		},
	}

```

## Unit test
```
go test
```


## Docker environment
```
docker-compose up -d
```


