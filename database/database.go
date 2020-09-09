package database

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	Name string `json:"name"`
	Age  int    `json:"age"`
	ID   string `json:"id"`
}

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	ID    string `json:"id"`
}

type Transaction struct {
	ID      string `json:"id"`
	BuyerID string `json:"buyerid"`
	IP      string `json:"ip"`
	Device  string `json:"device"`
}

type BuyerProfile struct {
	TransactionID string `json:"transactionid"`
	Device        string `json:"device"`
	Products      string `json:"products"`
	SameIP        string `json:"sameip"`
}

func Database() {
	// Init database connection
	fmt.Println("----------------------------------")
	fmt.Println("MySQL BuyerDB\n-------------------")
	db, err := sql.Open("mysql", "tester:@tcp(localhost:3306)/")
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
		buyerId := buyers.Buyers[i].ID
		for _, line := range csvTransactions {
			idOnTransaction := line[1]
			if buyerId == idOnTransaction {
				transactionId := line[0]
				data := "'" + buyerId + "','" + transactionId + "'"
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
	for _, pLine := range csvProducts {
		productId := pLine[0]
		for _, line := range csvTransactions {
			idsOnTransaction := line[4]
			if strings.Contains(idsOnTransaction, productId) {
				transactionId := line[0]
				data := "'" + productId + "','" + transactionId + "'"
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

func QueryBuyers() (exportBuyers string) {
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
		actual := Buyer{Name: name, Age: age, ID: id}
		jsonFormat, _ := json.Marshal(actual)
		exportBuyers += string(jsonFormat) + ","
	}
	exportBuyers = "[" + exportBuyers[:len(exportBuyers)-1] + "]"
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("-------------------------")
		fmt.Println("Successfully buyer writing on endpoint...")
	}
	return
}

func QueryProducts() (exportProducts string) {
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
		actual := Product{Name: name, Price: price, ID: id}
		jsonFormat, _ := json.Marshal(actual)
		exportProducts += string(jsonFormat) + ","
	}
	exportProducts = "[" + exportProducts[:len(exportProducts)-1] + "]"
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("-------------------------")
		fmt.Println("Successfully product writing on endpoint...")
	}
	return
}

func QueryTransactions() (exportTransactions string) {
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
		actual := Transaction{ID: id, BuyerID: buyerid, Device: device, IP: ip}
		jsonFormat, _ := json.Marshal(actual)
		exportTransactions += string(jsonFormat) + ","
	}
	exportTransactions = "[" + exportTransactions[:len(exportTransactions)-1] + "]"
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("-------------------------")
		fmt.Println("Successfully transaction writing on endpoint...")
	}
	return
}

func SearchBuyer(idBuyer string) (exportBuyers string) {
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
	rows, err := db.Query("SELECT * FROM Buyer WHERE id LIKE \"%" + idBuyer + "%\";")
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var id string
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		actual := Buyer{Name: name, Age: age, ID: id}
		jsonFormat, _ := json.Marshal(actual)
		exportBuyers += string(jsonFormat) + ","
	}
	// Bad query
	if len(exportBuyers) == 0 {
		exportBuyers = "[{\"name\": \"Not found\", \"age\": 0, \"id\": 0}]"
		return
	}
	// Good query
	exportBuyers = "[" + exportBuyers[:len(exportBuyers)-1] + "]"
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("-------------------------")
		fmt.Println("Successfully buyer search on endpoint...")
	}
	return
}

func TransactionBuyer(idBuyer string) (exportTransactions string) {
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
		" WHERE Buyer.id = \"" + idBuyer + "\";")
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var ignore string
		var transactionId string
		var device string
		var ip string
		var buyerId string
		err = rows.Scan(&buyerId, &ignore, &ignore, &ignore, &transactionId, &ignore, &ignore, &ip, &device)
		// Export transaction data
		actual := Transaction{ID: transactionId, Device: device, IP: ip, BuyerID: buyerId}
		jsonFormat, _ := json.Marshal(actual)
		exportTransactions += string(jsonFormat) + ","
	}
	exportTransactions = "[" + exportTransactions[:len(exportTransactions)-1] + "]"
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("-------------------------")
		fmt.Println("Successfully profile transactions writing on endpoint...")
	}
	return
}

func SameIPBuyer(transactionIp string) (exportSameip string) {
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
		" TransactionBuyerMapping.transaction_id = T.id WHERE T.ip = \"" + transactionIp + "\";")
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var id string
		var name string
		var age int
		var ignore string
		err = rows.Scan(&id, &name, &age, &ignore, &ignore, &ignore, &ignore, &ignore, &ignore)
		actual := Buyer{Name: name, Age: age, ID: id}
		jsonFormat, _ := json.Marshal(actual)
		exportSameip += string(jsonFormat) + ","
	}
	exportSameip = "[" + exportSameip[:len(exportSameip)-1] + "]"
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("-------------------------")
		fmt.Println("Successfully same IP buyers writing on endpoint...")
	}
	return
}

func TransactionProducts(transactionId string) (exportProducts string) {
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
		" TransactionProductMapping.transaction_id = T.id WHERE T.id = \"" + transactionId + "\";")
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var id string
		var name string
		var price int
		var ignore string
		err = rows.Scan(&id, &name, &price, &ignore, &ignore, &ignore, &ignore, &ignore, &ignore)
		actual := Product{Name: name, Price: price, ID: id}
		jsonFormat, _ := json.Marshal(actual)
		exportProducts += string(jsonFormat) + ","
	}
	exportProducts = "[" + exportProducts[:len(exportProducts)-1] + "]"
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("-------------------------")
		fmt.Println("Successfully product writing on endpoint...")
	}
	return
}
