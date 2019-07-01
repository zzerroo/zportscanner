# Note

This project is used for test whether there is a process listen on spcifical ports

# Build

go build main.go

# Usage

The project supported two params:

-  -i, the ip adress or the file that include all ip（separate by \n）
- -pr, the port  or the port range that is separated by -

The usage like this:

```
./main -i 127.0.0.1 -pr 8000 //single ip,single port
./main -i 127.0.0.1 -pr 8000-8881 //single ip,multi port
./main -i ip.file -pr 8000 //multi port,single port
./main -i ip.file -pr 8000-8881 //multi port,multi port
```

The output like this:

```
./main -i 127.0.0.1 -pr 8000-8881
127.0.0.1:
	8081
192.168.0.1:
	8082
```

