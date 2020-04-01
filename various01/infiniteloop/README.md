Folder Structure:  
https://github.com/golang-standards/project-layout  
https://eli.thegreenplace.net/2019/simple-go-project-layout-with-modules/  

Tutorial:  
https://golangbot.com/go-packages/  

Work:  
mkdir -p pkg --> to keep functions which can be public  
mkdir -p internal --> to keep functions which are internal (not public)  
mkdir -p pkg/functions --> for functions under package 'functions' - a kind of misnomer to use functions itself as functions package name!  
  
go mod init basic --> this will create a go.mod file and all import of fuctions will be relative to this  
main.go --> has import of basic/pkg/functions --> which has various go files with 'package functions' as the first line in the code  
(if you want to use github reference, then reference the git folder of the functions in main.go like www.github.com/coderdba-coding-org/golang2/various01/basic/pkg/functions" - and before compiling push this repo to git  
pkg/functions/func.go --> has first line as 'package functions' and defines some functions  
go build --> creates 'basic' as executable  
