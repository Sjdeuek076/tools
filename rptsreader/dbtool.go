package rptsreader

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"gopkg.in/mgutz/dat.v1/sqlx-runner"

	_ "github.com/lib/pq"
	"github.com/mgutz/dat"
)

const (
	host     = 
	port     = 
	user     = 
	password = 
	dbname   = 
)

var DB *runner.DB

func init() {
	// create a normal database connection through database/sql
	//db, err := sql.Open("postgres", "dbname=trading user=trader password=trader01 host=localhost sslmode=disable")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// ensures the database can be pinged with an exponential backoff (15 min)
	runner.MustPing(db)
	// set to reasonable values for production
	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(16)
	// set this to enable interpolation
	dat.EnableInterpolation = true
	// set to check things like sessions closing.
	// Should be disabled in production/release builds.
	dat.Strict = false
	// Log any query over 10ms as warnings. (optional)
	runner.LogQueriesThreshold = 10 * time.Millisecond
	DB = runner.NewDB(db, "postgres")
}

func In2db(inline string, intbl string) {
	invalue := strings.Split(inline, ",")
	switch intbl {
	case "futures_n_options", "tmp": //7 value
		DB.
			Upsert(intbl).
			Columns("tdate", "f_contract_volume", "f_open_interest", "o_contract_volume", "o_open_interest", "fo_contract_volume", "fo_open_interest").
			Values(invalue[0], invalue[1], invalue[2], invalue[3], invalue[4], invalue[5], invalue[6]).
			Where("tdate=$1", invalue[0]).
			//Returning("username", "departname"). //the last is dot "."
			Returning("*"). //the same as the above statement
			//ToSQL()  //return values to sql and args. It can't co-exist with Exec()
			Exec() //execute the sql and args. It can't co-exist with ToSQL()

	case "hsi_options", "mini_hsi_options", "stock_options":
		DB.
			Upsert(intbl).
			Columns("tdate", "call_contract_volume", "put_contract_volume", "total_contract_volume", "call_open_interest", "put_open_interest", "total_open_interest").
			Values(invalue[0], invalue[1], invalue[2], invalue[3], invalue[4], invalue[5], invalue[6]).
			Where("tdate=$1", invalue[0]).
			//Returning("username", "departname"). //the last is dot "."
			Returning("*"). //the same as the above statement
			//ToSQL()  //return values to sql and args. It can't co-exist with Exec()
			Exec() //execute the sql and args. It can't co-exist with ToSQL()

	case "hsi_vix_futures", "hsi_futures", "mini_hsi_futures", "rmb_futures": //5 values
		DB.
			Upsert(intbl).
			Columns("tdate", "spot_month", "second_month", "contract_volume", "open_interest").
			Values(invalue[0], invalue[1], invalue[2], invalue[3], invalue[4]).
			Where("tdate=$1", invalue[0]).
			//Returning("username", "departname"). //the last is dot "."
			Returning("*"). //the same as the above statement
			//ToSQL()  //return values to sql and args. It can't co-exist with Exec()
			Exec() //execute the sql and args. It can't co-exist with ToSQL()

	case "stock_futures": //3 value
		DB.
			Upsert(intbl).
			Columns("tdate", "contract_volume", "open_interest").
			Values(invalue[0], invalue[1], invalue[2]).
			Where("tdate=$1", invalue[0]).
			//Returning("username", "departname"). //the last is dot "."
			Returning("*"). //the same as the above statement
			//ToSQL()  //return values to sql and args. It can't co-exist with Exec()
			Exec() //execute the sql and args. It can't co-exist with ToSQL()
	}
}

//quering
/*rows, err := db.Query("select code, tdate from stock_d limit 10")
Check(err)
for rows.Next() {
	var code string
	var tdate time.Time
	err = rows.Scan(&code, &tdate)
	Check(err)
	fmt.Println("code | tdate ")
	fmt.Printf("%3v | %8v\n", code, tdate)
}*/

/*
Example for "github.com/lib/pq"
dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	Check(err)
	defer db.Close()
	//inserting value
	//Note that PostgreSQL uses the $1, $2 format instead of the ? that MySQL uses, and it has a different DSN format in sql.Open.
	//Another thing is that the Postgres driver does not support sql.Result.LastInsertId(). So instead of this,
	stmt, err := db.Prepare("INSERT INTO (username,departname,created) VALUES($1,$2,$3);")
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	fmt.Println(res.LastInsertId())
*/
