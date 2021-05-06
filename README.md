# LEARNING-GO :slightly_smiling_face:
  * Statically typed
  * Compiled programming language
### :one: BASIC COMMANDS
|       Commands      |             Description             |
|---------------------|-------------------------------------|
| make dev            | run dev server                      |
| go                  | list go commands                    |
| go build            | compile packages and dependencies   |
| go clean            | clean files generated by compilers  |
| go fmt              | format source code files            |
| go get              | install remote packages             |  
| go install          |  compiles all packages and generates files |
| go test             | test all your test files (*_test.go) | 
| go test ./...       |                                      |
| go test -cover ./...|                                     |
| godoc package_name  | look up the package documentation   |
| go fix              | upgrade code from an old version to a new version |
| go version          | get information about your version of Go |
| go env              | view environment variables about Go |
| go list             | list all installed packages         |
| go run              | compile temporary files and run the application |
| go mod init go_module | initialize the current directory as the root of the module |

### :two: TECHNOLOGIES USED
|    Tools    |     Description          |
|-------------|--------------------------|
|   Go        |    Programming Language  |
|   Gin       |    HTTP web frameweok    |
|   JWT       |      Authorization       |
|   GORM      | Full featured Go ORM library |
|   SQLite    | Small SQL database engine    |

### :three: PROJECT SETUP
  1. Clone the repo
  2. Create `.env` file in the root directory with the help of `.env.example` file
  3. Run app `make dev`

### :four: REFERENCES
1. [https://medium.com/technofunnel/understanding-golang-and-goroutines-72ac3c9a014d]
2. [https://medium.com/technofunnel/understanding-goroutine-go-channels-in-detail-9c5a28f08e0d]
3. [https://medium.com/technofunnel/error-handling-in-golang-with-panic-defer-and-recover-d77db7ae3875]
4. [https://www.golangprograms.com/go-language/struct.html]
5. [https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/01.3.html]
6. [https://towardsdatascience.com/building-restful-apis-in-golang-e3fe6e3f8f95]
7. [https://tutorialedge.net/golang/creating-restful-api-with-golang/]
8. [https://gorm.io/docs/]
9. [https://github.com/irahardianto/service-pattern-go]
10. [https://github.com/Alikhll/golang-developer-roadmap]
11. [https://tutorialedge.net/projects/chat-system-in-go-and-react/]
12. [https://medium.com/wesionary-team/dependency-injection-with-go-fx-b698a6585cf0]
13. [https://www.golangprograms.com/go-language.html]
14. [https://medium.com/easyread/today-i-learned-golang-live-reload-for-development-using-docker-compose-air-ecc688ee076]
15. [https://www.google.com/recaptcha/admin/create]
16. [https://devhints.io/go]

### Live reloading 
- reloads or refreshes the entire app when a file changes. For example, if you were four links deep into your navigation and saved a change, live reloading would restart the app and load the app back to the initial route.
  [https://techinscribed.com/5-ways-to-live-reloading-go-applications/]

### Hot reloading 
- only refreshes the files that were changed without losing the state of the app. For example, if you were four links deep into your navigation and saved a change to some styling, the state would not change, but the new styles would appear on the page without having to navigate back to the page you are on because you would still be on the same page.
