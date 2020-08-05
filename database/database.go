package database

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "encoding/json"
    "encoding/csv"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)

// An array of Buyers
type Buyers struct {
    Buyers []Buyer `json:"buyers"`
}

type Buyer struct {
    Name   string `json:"name"`
    Age    int    `json:"age"`
    ID     string `json:"id"`
}

type Product struct {
    Name   string `json:"name"`
    Price    int    `json:"price"`
    ID     string `json:"id"`
}

type Transaction struct {
    ID   string `json:"name"`
    BuyerID    string `json:"buyerid"`
    IP  string `json:"ip"`
    Device string `json:"device"`
}

type BuyerProfile struct {
    TransactionID   string `json:"transactionid"`
    Device string `json:"device"`
    Products string `json:"products"`
    SameIP string `json:"sameip"`
}

func Database() {
    // Init database connection
    fmt.Println("----------------------------------")
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
	_, err = db.Exec("CREATE database IF NOT EXISTS BuyerDB")
    if err != nil {
    fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully created database...")
    }
    // Use database
    _, err = db.Exec("USE BuyerDB")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully selected database...")
    }
	// Create Buyer table
    _, err = db.Exec(`
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
    _, err = db.Exec(`
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
    _, err = db.Exec(`
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
    _, err = db.Exec(`
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
    _, err = db.Exec(`
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
    jsonFile, err := os.Open("responses/buyers.json")
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
        _, err = db.Exec("INSERT IGNORE into Buyer VALUES (" + data + ");")
    }
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully added buyers data...")
    }
    // Add products data
    csvFile, err := os.Open("responses/products.csv")
    if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
    csvProducts, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
        fmt.Println(err)
    }
    for _, line := range csvProducts {
        data := "'" + line[0] + "'" + "," + "'" + line[1] + "'" + "," + line[2]
        _, err = db.Exec("INSERT IGNORE into Product VALUES (" + data + ");")
    }
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully added products data...")
    }
    // Add transactions data
    csvFile, err = os.Open("responses/transactions.csv")
    if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
    csvTransactions, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
        fmt.Println(err)
    }
    for _, line := range csvTransactions {
        data := "'" + line[0] + "'" + "," + "'" + line[1] + "'" + "," + "'" + line[2] + "'" + "," + "'" + line[3] + "'"
        _, err = db.Exec("INSERT IGNORE into Transaction VALUES (" + data + ");")
    }
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully added transactions data...")
    }
    // Add buyer mapping with transactions
    for i := 0; i < len(buyers.Buyers); i++ {
        buyer_id := buyers.Buyers[i].ID
        for _, line := range csvTransactions {
            id_on_transaction := line[1]
            if buyer_id == id_on_transaction {
                transaction_id := line[0]
                data := "'" + buyer_id + "','" +  transaction_id + "'"
                _, err = db.Exec("INSERT IGNORE into TransactionBuyerMapping VALUES (" + data + ");")
            }
        }
    }
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully added buyer mapping data...")
    }
    // Add product mapping with transactions
    for _, p_line := range csvProducts {
        product_id := p_line[0]
        for _, line := range csvTransactions {
            ids_on_transaction := line[4]
            if strings.Contains(ids_on_transaction, product_id) {
                transaction_id := line[0]
                data := "'" + product_id + "','" +  transaction_id + "'"
                _, err = db.Exec("INSERT IGNORE into TransactionProductMapping VALUES (" + data + ");")
            }
        }
    }
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully added product mapping data...")
    }
}

func QueryBuyers() (export_buyers string) {
    // Init database connection
    fmt.Println("----------------------------------")
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
    // Use database
    _, err = db.Exec("USE BuyerDB")
    if err != nil {
        fmt.Println(err.Error())
    }
    // Query buyers
    rows, err := db.Query("SELECT * FROM Buyer")
    if err != nil {
        fmt.Println(err.Error())
    }
    for rows.Next() {
        var id string
        var name string
        var age int
        err = rows.Scan(&id, &name, &age)
        actual := Buyer {Name: name, Age: age, ID: id}
        json_format, _ := json.Marshal(actual)
        export_buyers += string(json_format) + ","
    }
    export_buyers = "[" + export_buyers + "]"
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("-------------------------")
        fmt.Println("Successfully buyer writing on endpoint...")
    }
    return
}

func QueryProducts() (export_products string) {
    // Init database connection
    fmt.Println("----------------------------------")
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
    // Use database
    _, err = db.Exec("USE BuyerDB")
    if err != nil {
        fmt.Println(err.Error())
    }
    // Query products
    rows, err := db.Query("SELECT * FROM Product")
    if err != nil {
        fmt.Println(err.Error())
    }
    for rows.Next() {
        var id string
        var name string
        var price int
        err = rows.Scan(&id, &name, &price)
        actual := Product {Name: name, Price: price, ID: id}
        json_format, _ := json.Marshal(actual)
        export_products += string(json_format) + ","
    }
    export_products = "[" + export_products + "]"
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("-------------------------")
        fmt.Println("Successfully product writing on endpoint...")
    }
    return
}

func QueryTransactions() (export_transactions string) {
    // Init database connection
    fmt.Println("----------------------------------")
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
    // Use database
    _, err = db.Exec("USE BuyerDB")
    if err != nil {
        fmt.Println(err.Error())
    }
    // Query transactions
    rows, err := db.Query("SELECT * FROM Transaction")
    if err != nil {
        fmt.Println(err.Error())
    }
    for rows.Next() {
        var id string
        var buyerid string
        var ip string
        var device string
        err = rows.Scan(&id, &buyerid, &ip, &device)
        actual := Transaction {ID: id, BuyerID: buyerid, Device: device, IP: ip}
        json_format, _ := json.Marshal(actual)
        export_transactions += string(json_format) + ","
    }
    export_transactions = "[" + export_transactions + "]"
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("-------------------------")
        fmt.Println("Successfully transaction writing on endpoint...")
    }
    return
}

func SearchBuyer(id_buyer string) (export_buyers string) {
    // Init database connection
    fmt.Println("----------------------------------")
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
    // Use database
    _, err = db.Exec("USE BuyerDB")
    if err != nil {
        fmt.Println(err.Error())
    }
    // Query buyer transactions
    rows, err := db.Query("SELECT * FROM Buyer WHERE id LIKE \"%" + id_buyer + "%\";")
    if err != nil {
        fmt.Println(err.Error())
    }
    for rows.Next() {
        var id string
        var name string
        var age int
        err = rows.Scan(&id, &name, &age)
        actual := Buyer {Name: name, Age: age, ID: id}
        json_format, _ := json.Marshal(actual)
        export_buyers += string(json_format) + ","
    }
    export_buyers = "[" + export_buyers + "]"
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("-------------------------")
        fmt.Println("Successfully buyer search on endpoint...")
    }
    return
}

func TransactionBuyer(id_buyer string) (export_transactions string) {
    // Init database connection
    fmt.Println("----------------------------------")
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
    // Use database
    _, err = db.Exec("USE BuyerDB")
    if err != nil {
        fmt.Println(err.Error())
    }
    // Query buyer transactions
    rows, err := db.Query("SELECT * FROM Buyer JOIN TransactionBuyerMapping on" +
    " Buyer.id = TransactionBuyerMapping.buyer_id JOIN Transaction T on TransactionBuyerMapping.transaction_id = T.id" +
    " WHERE Buyer.id = \"" + id_buyer + "\";")
    if err != nil {
        fmt.Println(err.Error())
    }
    for rows.Next() {
        var ignore string
        var transaction_id string
        var device string
        var ip string
        var buyer_id string
        err = rows.Scan(&buyer_id, &ignore, &ignore, &ignore, &transaction_id, &ignore, &ignore, &ip, &device)
        // Export transaction data
        actual := Transaction {ID: transaction_id, Device: device, IP: ip, BuyerID: buyer_id}
        json_format, _ := json.Marshal(actual)
        export_transactions += string(json_format) + ","
    }
    export_transactions = "[" + export_transactions + "]"
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("-------------------------")
        fmt.Println("Successfully profile transactions writing on endpoint...")
    }
    return
}

func SameIPBuyer(transaction_ip string) (export_sameip string){
    // Init database connection
    fmt.Println("----------------------------------")
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
    // Use database
    _, err = db.Exec("USE BuyerDB")
    if err != nil {
        fmt.Println(err.Error())
    }
    // Query products on transaction
    rows, err := db.Query("SELECT * FROM Buyer JOIN TransactionBuyerMapping on" +
    " Buyer.id = TransactionBuyerMapping.buyer_id JOIN Transaction T on" +
    " TransactionBuyerMapping.transaction_id = T.id WHERE T.ip = \"" + transaction_ip + "\";")
    if err != nil {
        fmt.Println(err.Error())
    }
    for rows.Next() {
        var id string
        var name string
        var age int
        var ignore string
        err = rows.Scan(&id, &name, &age, &ignore, &ignore, &ignore, &ignore, &ignore, &ignore)
        actual := Buyer {Name: name, Age: age, ID: id}
        json_format, _ := json.Marshal(actual)
        export_sameip += string(json_format) + ","
    }
    export_sameip = "[" + export_sameip + "]"
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("-------------------------")
        fmt.Println("Successfully same IP buyers writing on endpoint...")
    }
    return
}


func TransactionProducts(transaction_id string) (export_products string){
    // Init database connection
    fmt.Println("----------------------------------")
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
    // Use database
    _, err = db.Exec("USE BuyerDB")
    if err != nil {
        fmt.Println(err.Error())
    }
    // Query products on transaction
    rows, err := db.Query("SELECT * FROM Product JOIN TransactionProductMapping on" +
    " Product.id = TransactionProductMapping.product_id JOIN Transaction T on" +
    " TransactionProductMapping.transaction_id = T.id WHERE T.id = \"" + transaction_id + "\";")
    if err != nil {
        fmt.Println(err.Error())
    }
    for rows.Next() {
        var id string
        var name string
        var price int
        var ignore string
        err = rows.Scan(&id, &name, &price, &ignore, &ignore, &ignore, &ignore, &ignore, &ignore)
        actual := Product {Name: name, Price: price, ID: id}
        json_format, _ := json.Marshal(actual)
        export_products += string(json_format) + ","
    }
    export_products = "[" + export_products + "]"
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("-------------------------")
        fmt.Println("Successfully product writing on endpoint...")
    }
    return
}
