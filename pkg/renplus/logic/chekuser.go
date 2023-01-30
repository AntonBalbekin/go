package logic

import (
	"fmt"
	_ "log"
	"net/http"
	db "usr/api/pkg/renplus/sqlconect"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID                 int    `json:"ID"from"ID"`
	UF_TEST_UCM_ID     string `json:"UF_TEST_UCM_ID"from"UF_TEST_UCM_ID"`
	UF_TEST_FIELD_ONE  string `json:"UF_TEST_FIELD_ONE"from"UF_TEST_FIELD_ONE"`
	UF_TEST_FIELD_THRE string `json:"UF_TEST_FIELD_THRE"from"UF_TEST_FIELD_THRE"`
}

func ChekClient(cont *gin.Context) {
	userId := cont.Request.FormValue("userid")
	var sql string
	if userId != "" {
		sql = fmt.Sprintf(`select ID,UF_TEST_UCM_ID,UF_TEST_FIELD_ONE,UF_TEST_FIELD_THRE from test_client where UF_TEST_UCM_ID=%s order by UF_TEST_UCM_ID`, userId)
	} else {
		sql = "select ID,UF_TEST_UCM_ID,UF_TEST_FIELD_ONE,UF_TEST_FIELD_THRE from test_client"
	}
	rows, err := db.SqlDB.Query(sql)
	defer rows.Close()

	if err != nil {
		return
	}

	users := []User{}
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.ID, &user.UF_TEST_UCM_ID, &user.UF_TEST_FIELD_ONE, &user.UF_TEST_FIELD_THRE)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, user)

	}
	cont.JSON(http.StatusOK, gin.H{"users": users, "get": sql})
}
