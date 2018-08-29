package dbops

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func AddUserCredential(loginName string, pwd string) error  {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?,?)")
	if err!=nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err!=nil {
		return err
	}

	defer stmtIns.Close() // 性能消耗较大

	return nil
}

func GetUserCredential(loginName string) (string, error)  {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	// noRows 没有结果返回
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return err
	}
	stmtDel.Exec(loginName, pwd)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	defer stmtDel.Close()
	return nil
}