# Go Dependency Exclusion Test

Test whether the go build command can exclude dependencies from the final binary.

```
go build -o deluxe-build -tags=deluxe main.go
```

```
go build -o normal-build main.go
```


```
$ ls -l
drwxr-xr-x   4 user  group      128     Apr 20 21:55 deluxe
-rwxr-xr-x   1 user  group  **2000786** Apr 20 22:03 deluxe-build
-rw-r--r--   1 user  group      247     Apr 20 21:57 go.mod
-rw-r--r--   1 user  group      865     Apr 20 21:57 go.sum
-rw-r--r--   1 user  group      105     Apr 20 22:01 main.go
-rwxr-xr-x   1 user  group  **1190722** Apr 20 22:03 normal-build
```