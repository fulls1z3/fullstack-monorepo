## Go services 
- **cmd:**
The `cmd` directory contains entry points - each folder representing the
service name and holding the `main` package.

- **lib:**
The `lib` directory contains application logic, while each package holding
the pure business logic of the domain.

- **pkg:**
The `pkg` directory contains all **shared** code and libraries, and must
not contain any service specific code.
