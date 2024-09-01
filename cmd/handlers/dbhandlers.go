package handlers

import (
	"database/sql"
	"someweb/cmd/models"
)

var connstring string

func SetDBConn(conn string) {
	connstring = conn
}

func GetArticles() []models.Article {
	var posts []models.Article

	db, err := sql.Open("postgres", connstring)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `SELECT * FROM Posts order by posteddate DESC`
	rows, err := db.Query(sqlStatement) //db.Exec(sqlStatement, filename, filesize)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.Id, &article.PostedDate, &article.Title, &article.Content)
		if err != nil {
			panic(err)
		}
		posts = append(posts, article)
	}
	return posts
}
