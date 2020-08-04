package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
    "encoding/csv"
    "io/ioutil"
    "os"
    "strconv"
)

// an array of Buyers
type Buyers struct {
    Buyers []Buyer `json:"buyers"`
}

type Buyer struct {
    Name   string `json:"name"`
    Age    int    `json:"age"`
    ID     string `json:"id"`
}

func main() {
    // Init database connection
    fmt.Println("MySQL BuyerDB\n-------------------")
    db, err := sql.Open("mysql", "tester:@tcp(localhost:3306)/BuyerDB")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    // Check if there is connection
    err = db.Ping()
	if err != nil {
		fmt.Printf("Error connection: %v", err)
		return
	}
	// Create database
	_,err = db.Exec("CREATE database IF NOT EXISTS BuyerDB")
    if err != nil {
    fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully created database...")
    }
    // Use database
    _,err = db.Exec("USE BuyerDB")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully selected database...")
    }
	// Create Buyer table
    _,err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Buyer
        (
            id   varchar(12),
            name varchar(30),
            age  tinyint,
            PRIMARY KEY (id)
        );
    `)
    if err != nil {
    fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully created Buyer table...")
    }
    // Create Product table
    _,err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Product
        (
            id    varchar(12),
            name  varchar(100),
            price smallint,
            PRIMARY KEY (id)
        );
    `)
    if err != nil {
    fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully created Product table...")
    }
	// Create Transaction table
    _,err = db.Exec(`
        CREATE TABLE IF NOT EXISTS Transaction
        (
            id       varchar(40),
            buyer_id varchar(12),
            ip       varchar(20),
            device   varchar(12),
            PRIMARY KEY (id)
        );
    `)
    if err != nil {
    fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully created Transaction table...")
    }
    // Create mapping tables
    // Buyer with Transaction
    _,err = db.Exec(`
        CREATE TABLE IF NOT EXISTS TransactionBuyerMapping
        (
            buyer_id       varchar(12),
            transaction_id varchar(12),
            PRIMARY KEY (buyer_id, transaction_id),
            FOREIGN KEY (buyer_id) REFERENCES Buyer (id),
            FOREIGN KEY (transaction_id) REFERENCES Transaction (id)
        );
    `)
    if err != nil {
    fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully created TransactionBuyerMapping table...")
    }
    // Product with Transaction
    _,err = db.Exec(`
        CREATE TABLE IF NOT EXISTS TransactionProductMapping
        (
            product_id     varchar(12),
            transaction_id varchar(12),
            PRIMARY KEY (product_id, transaction_id),
            FOREIGN KEY (product_id) REFERENCES Product (id),
            FOREIGN KEY (transaction_id) REFERENCES Transaction (id)
        );
    `)
    if err != nil {
    fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully created TransactionProductMapping table...")
    }
    // Start adding data from response files
    fmt.Println("-----------------------")
    // Add buyers data
    jsonFile, err := os.Open("api-rest/responses/buyers.json")
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)
    var buyers Buyers
    json.Unmarshal(byteValue, &buyers)
    for i := 0; i < len(buyers.Buyers); i++ {
        id := buyers.Buyers[i].ID
        name := buyers.Buyers[i].Name
        age := strconv.FormatInt(int64(buyers.Buyers[i].Age), 10)
        data := "'" + id + "'" + "," + "'" + name + "'" + "," + age
        _,err = db.Exec("INSERT IGNORE into Buyer VALUES (" + data + ");")
    }
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully added buyers data...")
    }
    // Add products data
    csvFile, err := os.Open("api-rest/responses/products.csv")
    if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
    csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
        fmt.Println(err)
    }
    for _, line := range csvLines {
        data := "'" + line[0] + "'" + "," + "'" + line[1] + "'" + "," + line[2]
        _,err = db.Exec("INSERT IGNORE into Product VALUES (" + data + ");")
    }
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully added products data...")
    }
    // Add transactions data
    csvFile, err = os.Open("api-rest/responses/transactions.csv")
    if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
    csvLines, err = csv.NewReader(csvFile).ReadAll()
    if err != nil {
        fmt.Println(err)
    }
    for _, line := range csvLines {
        data := "'" + line[0] + "'" + "," + "'" + line[1] + "'" + "," + "'" + line[2] + "'"
        fmt.Println(data)
        _,err = db.Exec("INSERT IGNORE into Transaction VALUES (" + data + ");")
    }
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully added transactions data...")
    }
}