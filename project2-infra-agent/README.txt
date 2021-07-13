=====================
REFERENCES
=====================
Cobra for command-line CLI programming: 
- root.go is cobra's way of default program
- https://levelup.gitconnected.com/exploring-go-packages-cobra-fce6c4e331d6

Viper for config settings: 
- https://dev.to/techschoolguru/load-config-from-file-environment-variables-in-golang-with-viper-2j2d
- https://medium.com/@felipedutratine/manage-config-in-golang-to-get-variables-from-file-and-env-variables-33d876887152

Logging 
- Adding WARN, INFO, CRIT etc: https://www.honeybadger.io/blog/golang-logging/

=====================
SETUP
=====================

- Create github.com directory under ~/go directory
mkdir -p ~/go/src/github.com
- Link the work-directory bearing the org-name as a link under ~/go/src/github.com
ln -s ~/stuff/mystuff/coderdba-coding-org/ ~/go/src/github.com
--> Doing so is equivalent of working under a directory ~/go/src/github.com/coderdba-coding-org/golang2/project2-infra-agent

- In the project directory, initialize go module
go mod init www.github.com/coderdba-coding-org/golang2/project2-infra-agent

- Run the project
go run .
