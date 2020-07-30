# gomodules2  
  
https://medium.com/@adiach3nko/package-management-with-go-modules-the-pragmatic-guide-c831b4eaaf31  

###  Initialize the module as a git repo  
NOTE: Alternatively, as a local module you can do go mod init gomodules2 - which will have home as current folder  
\$ go mod init github.com/coderdba-coding-org/golang2/various01/gomodules2  

###  Get modules that we think we need  
\$ go get github.com/coderdba-coding-org/gojunk@master  
```
go: finding github.com/coderdba-coding-org/gojunk master  
go: downloading github.com/coderdba-coding-org/gojunk v0.0.0-20200402072544-6b5e6658f6f0  
go: extracting github.com/coderdba-coding-org/gojunk v0.0.0-20200402072544-6b5e6658f6f0  

\$ ls -l  
-rw-r--r--  1 username  somedomain Users  653 Apr  2 12:56 README.md  
drwxr-xr-x  7 username  somedomain Users  224 Apr  2 12:56 bak  
-rw-------  1 username  somedomain Users  152 Apr  2 12:56 go.mod  
-rw-------  1 username  somedomain Users  249 Apr  2 12:56 go.sum  
```

### Find the downloaded module in $GOPATH which is as of now ~/go (the default)  
\$ cd ~/go  
\$ find . -name gojunk -print  
./pkg/mod/cache/download/github.com/coderdba-coding-org/gojunk  

\$ cat go.mod  

```
module github.com/coderdba-coding-org/golang2/various01/gomodules2  

go 1.12  

require github.com/coderdba-coding-org/gojunk v0.0.0-20200402072544-6b5e6658f6f0 // indirect  
```


\$ cat go.sum  
```
github.com/coderdba-coding-org/gojunk v0.0.0-20200402072544-6b5e6658f6f0 h1:flZwNkZxVTVRndx+5BudIkTpsoF8CK7cXQh+udBzMmk=  
github.com/coderdba-coding-org/gojunk v0.0.0-20200402072544-6b5e6658f6f0/go.mod h1:7iAJonrssGVS+QkuwAFsbC3Uzxa9Um5eYnUjj0gNzfM=  
```

### Create a program that uses the module gojunk  

\$ cat main.go  
```
package main

import (
       "github.com/coderdba-coding-org/gojunk"
)
```

### Tidy the module file  
NOTE: Note that the "//indirect" goes away as we are using the gojunk repo in main.go (or other program in this module)  

\$ go mod tidy  

\$ cat go.mod  
```
module github.com/coderdba-coding-org/golang2/various01/gomodules2  

go 1.12  

require github.com/coderdba-coding-org/gojunk v0.0.0-20200402072544-6b5e6658f6f0  
```

### Get the module into vendor folder within this module  
$ go mod vendor  

$ cd vendor  
$ ls -lR  
```

total 8
drwxr-xr-x  3 username  somedomain Users   96 Apr  2 13:10 github.com
-rw-r--r--  1 username  somedomain Users  113 Apr  2 13:10 modules.txt

./github.com:
total 0
drwxr-xr-x  3 username  somedomain Users  96 Apr  2 13:10 coderdba-coding-org

./github.com/coderdba-coding-org:
total 0
drwxr-xr-x  5 username  somedomain Users  160 Apr  2 13:10 gojunk

./github.com/coderdba-coding-org/gojunk:
total 24
-rw-r--r--  1 username  somedomain Users   8 Apr  2 13:10 README.md
-rw-r--r--  1 username  somedomain Users  43 Apr  2 13:10 git-push.sh
-rw-r--r--  1 username  somedomain Users  91 Apr  2 13:10 junk1.go
```

### Run the main program
\$ go run .
```
Hello from gojunk
```

### Get the latest version of the module - gojunk

- Get
\$ go get -u 
NOTE: No need to give the specific module name - will update all  


```
go: finding github.com/coderdba-coding-org/gojunk latest
go: downloading github.com/coderdba-coding-org/gojunk v0.0.0-20200402074708-955f14ff0fec
go: extracting github.com/coderdba-coding-org/gojunk v0.0.0-20200402074708-955f14ff0fec
```

- Update vendor folder
\$ go mod vendor

- Test by running the main program
NOTE: This gives an output with a line before and after the Hello line
\$ go run .
```

Hello from gojunk
```
