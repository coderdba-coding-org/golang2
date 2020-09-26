# EXAMPLES
https://oralytics.com/2019/05/20/connecting-go-lang-to-oracle-database/
https://oralytics.com/2019/07/08/golang-querying-records-from-oracle-database-using-goracle/
https://static.rainfocus.com/oracle/oow19/sess/1567058525476001cK8G/PF/DEV6708-Using-the-Go-Language-for-Efficient-Oracle-Database-Applications_1568841171132001jI7d.pdf

# DRIVER
https://gopkg.in/goracle.v2
https://godoc.org/gopkg.in/goracle.v2

# CLIENT LIBRARIES (Oracle Instant Client)
https://oracle.github.io/odpi/doc/installation.html#macos

# HOW TO (on MAC)
https://oralytics.com/2019/05/20/connecting-go-lang-to-oracle-database/

Install Oracle client - and add that to PATH
Example: /Users/username/opt/oracle/instantclient_19_3/

Copy the following to ~/lib
libclntsh.dylib --> which is same as libclntsh.dylib.19.1
libclntshcore.dylib.19.1
libnnz19.dylib

Install Oracle driver/api for Go
go get gopkg.in/goracle.v2


