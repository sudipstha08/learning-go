# LEARNING-GO :slightly_smiling_face:
  * Statically typed
  * Compiled programming language
  
### :one: BASIC COMMANDS
|       commands      |             description             |
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
| go mod tidy | adds any missing modules necessary to build the current module's |

### :two: TECH STACKS
|    tools    |     description              |
|-------------|------------------------------|
|   `go`      |    programming language      |
|   `gin`     |    http web framework for go |
|   `jwt`     |      Authorization           |
|   `gorm`    | Full featured Go ORM library |
|   `sqlite`  | Small SQL database engine    |

### :three: PROJECT SETUP
  1. Clone the `repo`
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
17. [https://docs.sentry.io/platforms/go/]


### Fix
- Username and Password not accepted. Learn more at\n5.7.8  https://support.google.com/mail/?p=BadCredentials mt24sm3123297pjb.18 - gsmtp
```sh
Goto your gmail security settings and allow from less secured apps
```

### Sentry :book:
- Often we get asked what the difference is between Sentry and traditional logging. It makes a lot of sense when you're looking at something like a web application, where your errors typically go to a logfile on disk. When you start looking at mobile or desktop applications things start to become a bit more clear.

- Sentry supports every major language, framework, and library.
  
- You can get started for free. Pricing depends on the number of monthly events, transactions, and attachments that you send Sentry

- Sentry doesn’t impact a web site’s performance.Sentry is a listener/handler for errors that asynchronously sends out the error/event to Sentry.io. This is non-blocking. The error/event only goes out if this is an error.
Global handlers have almost no impact as well, as they are native APIs provided by the browsers.

- The best way to think about things is this: logging provides you with a trail of events. Often those events are errors, but many times they're simply informational. Sentry is fundamentally different because we focus on exceptions, or in other words, we capture application crashes.

- To summarize, Sentry works with your application logging infrastructure, often integrating directly. It does not replace the need for those logs, and it's also not a destination for things that aren't actionable errors or crashes.

- Difference between sentry and traditional logging 
  [https://sentry.io/vs/logging/]

- Sentry's Go SDK enables reporting messages, errors, and panics.
  
- Here we will see to integrate gin framework with sentry

- Sentry automatically assigns you a Data Source Name (DSN) when you create a project to start monitoring events in your app.

- A DSN tells a Sentry SDK where to send events so the events are associated with the correct project.

- If this value is not provided, SDKs will try to read it from the SENTRY_DSN environment variable, where applicable. This fallback does not apply in cases like a web browser, where the concept of environment variables does not exist.

- DSNs are safe to keep public because they only allow submission of new events and related event data; they do not allow read access to any information.

- Where to Find Your DSN
If you forget your DSN, view Settings -> Projects -> Client Keys (DSN) in sentry.io.

- The Parts of the DSN
The DSN configures the protocol, public key, server address, and project identifier. It is composed of the following parts:
```sh
{PROTOCOL}://{PUBLIC_KEY}:{SECRET_KEY}@{HOST}{PATH}/{PROJECT_ID}
```

### Live reloading 
- reloads or refreshes the entire app when a file changes. For example, if you were four links deep into your navigation and saved a change, live reloading would restart the app and load the app back to the initial route.
  [https://techinscribed.com/5-ways-to-live-reloading-go-applications/]

### Hot reloading 
- only refreshes the files that were changed without losing the state of the app. For example, if you were four links deep into your navigation and saved a change to some styling, the state would not change, but the new styles would appear on the page without having to navigate back to the page you are on because you would still be on the same page.