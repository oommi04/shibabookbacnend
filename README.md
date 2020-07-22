#### Setup project env

```sh
PORT=8080
```

#### How to run

```bash
$ go run main.go
```

#### How to build

```bash
$ go build
```

#### How to test

```bash
$ go test ./... -v
```

#### Folder Architecture

```sh
├── repository
│   ├── mongo
│   │   └──repository.go
│   └── interface.go
├── usecase
│   └── usecase1
│       └── usecase1Interface.go
├── domain
│   └── domain1
│       ├── errors.go
│       └── entity.go
├── drivers
│   └── driver1
│      └──driver.go
├── external
│   └── external1
│      └──external1.go
└── README.md
```

#### Layer


| Layer | Folder |
| ------ | ------ |
| Entities | domain |
| Use Case | usecase |
| Interface Adapters | repository/external |
| Frameworks and Drivers | drivers/handlers |


#### Entity
![alt text](https://github.com/oommi04/shibabookbackend/blob/[branch]/entity_schema.jpg?raw=true)

#### Sequnce Diagram
![alt text](https://github.com/oommi04/shibabookbackend/blob/[branch]/sequnce_diagram.png?raw=true)