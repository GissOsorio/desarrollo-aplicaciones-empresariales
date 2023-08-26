package db

import (
    "database/sql"
    "fmt"
    "log"

	"back/pkg/mocks"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 50000
    user     = "develop"
    password = "develop"
    dbname   = "develop"
)

func Connect() *sql.DB {
    connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", connInfo)
    if err != nil {
        log.Fatal(err)
    }
    err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("Successfully connected to db!")
    return db
}

func CloseConnection(db *sql.DB) {
    defer db.Close()
}

func CreateTableUsers(db *sql.DB) {
    var exists bool
    if err := db.QueryRow("SELECT EXISTS (SELECT FROM pg_tables WHERE  schemaname = 'public' AND tablename = 'lists' );").Scan(&exists); err != nil {
        fmt.Println("failed to execute query", err)
        return
    }
    if !exists {
        users, err := db.Query("CREATE TABLE users (id VARCHAR(100) PRIMARY KEY, date timestamp, name VARCHAR(50) NOT NULL);")
        if err != nil {
            fmt.Println("failed to execute query", err)
            return
        }
        fmt.Println("Table created successfully", users)

        for _, user := range mocks.Users {
            queryStmt := `INSERT INTO users (id,date,name) VALUES ($1, $2, $3) RETURNING id;`

            err := db.QueryRow(queryStmt, &user.Id, &user.Date, &user.Name).Scan(&user.Id)
            if err != nil {
                log.Println("failed to execute query", err)
                return
            }
        }
        fmt.Println("Mock Users included in Table", users)
    } else {
        fmt.Println("Table 'users' already exists ")
    }

}

func CreateTableLists(db *sql.DB) {
    var exists bool
    if err := db.QueryRow("SELECT EXISTS (SELECT FROM pg_tables WHERE  schemaname = 'public' AND tablename = 'lists' );").Scan(&exists); err != nil {
        fmt.Println("failed to execute query", err)
        return
    }
    if !exists {
        lists, err := db.Query("CREATE TABLE lists (id UUID PRIMARY KEY, date timestamp, userId VARCHAR(100) NOT NULL,name VARCHAR(50) NOT NULL, status VARCHAR(50) NOT NULL);")
        if err != nil {
            fmt.Println("failed to execute query", err)
            return
        }
        fmt.Println("Table created successfully", lists)

        for _, list := range mocks.Lists {
            queryStmt := `INSERT INTO lists (id,date,userId,name,status) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

            err := db.QueryRow(queryStmt, &list.Id, &list.Date, &list.UserId, &list.Name, &list.Status).Scan(&list.Id)
            if err != nil {
                log.Println("failed to execute query", err)
                return
            }
        }
        fmt.Println("Mock Lists included in Table", lists)
    } else {
        fmt.Println("Table 'lists' already exists ")
    }

}

func CreateTableElements(db *sql.DB) {
    var exists bool
    if err := db.QueryRow("SELECT EXISTS (SELECT FROM pg_tables WHERE  schemaname = 'public' AND tablename = 'elements' );").Scan(&exists); err != nil {
        fmt.Println("failed to execute query", err)
        return
    }
    if !exists {
        results, err := db.Query("CREATE TABLE elements (id UUID PRIMARY KEY, date timestamp, listId  UUID REFERENCES lists(id) ON DELETE CASCADE, content VARCHAR(50) NOT NULL, status VARCHAR(50) NOT NULL);")
        if err != nil {
            fmt.Println("failed to execute query", err)
            return
        }
        fmt.Println("Table created successfully", results)

        for _, element := range mocks.Elements {
            queryStmt := `INSERT INTO elements (id,date,listId,content,status) VALUES ($1, $2, $3, $4, $5) RETURNING id;`

            err := db.QueryRow(queryStmt, &element.Id, &element.Date, &element.ListId, &element.Content, &element.Status).Scan(&element.Id)
            if err != nil {
                log.Println("failed to execute query", err)
                return
            }
        }
        fmt.Println("Mock Elements included in Table", results)
    } else {
        fmt.Println("Table 'elements' already exists ")
    }

}