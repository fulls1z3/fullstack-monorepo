## Home for Golang Services 
- **cmd:**
The command folder contains all the entrypoints for any application we have,
its split into folders each folder represents the application name, this folder
will always hold the `main` package and the `main` function. Inside the command
go file you will find that the CLI handlers are used through cobra library.

- **infra:**
The library folder contains all our infrastructure logic, 
each package holds code which responsible to access storage 
(Sql or NoSql DB, Redis, Search Engine) or 3d party API to get or persist data, 
the package `lib` is responsible to trigger the `infra`.

- **lib:**
The library folder contains all our applications logic, each package holds the
pure business logic of the domain, the package `srv` is responsible to trigger
the `lib` processors.

- **pkg:**
The pkg folder contains all shared code and libraries, e.g connection to DB, logging,
http or grpc server etc. The folder must not contain any service specific code.

- **srv:**
The srv folder contains all applications Server concrete implementation. E.g. gRPC-server interface implementation.
`cmd` is responsible to trigger the `srv` handler _(could be gRPC, HTTP, CLI MQ handler)_.
