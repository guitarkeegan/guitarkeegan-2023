package models

import (
	"database/sql"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

type Blog struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	P1          *string    `json:"p1"`
	P2          *string    `json:"p2"`
	P3          *string    `json:"p3"`
	P4          *string    `json:"p4"`
	P5          *string    `json:"p5"`
	CS1         *string    `json:"cs1"`
	CS2         *string    `json:"cs2"`
	CS3         *string    `json:"cs3"`
	Repo        *string    `json:"repo"`
	CreatedAt   *time.Time `json:"createdAt"`
}

var DB *sql.DB

func CreateDBConnection() error {

	db, err := sql.Open("sqlite", "db.db")
	if err != nil {
		log.Fatal(err)
		// Handle the error
	}

	DB = db
	return nil

}

func GetPosts(count int) ([]Blog, error) {

	rows, err := DB.Query("SELECT * FROM blog")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := make([]Blog, 0)

	for rows.Next() {
		singlePost := Blog{}
		err = rows.Scan(&singlePost.ID, &singlePost.Title, &singlePost.Description, &singlePost.P1, &singlePost.P2, &singlePost.P3, &singlePost.P4, &singlePost.P5, &singlePost.CS1, &singlePost.CS2, &singlePost.CS3, &singlePost.Repo, &singlePost.CreatedAt)

		if err != nil {
			return nil, err
		}

		posts = append(posts, singlePost)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return posts, err
}

func GetPostById(id string) (Blog, error) {
	stmt, err := DB.Prepare("SELECT * FROM blog WHERE id = ?")

	if err != nil {
		return Blog{}, err
	}

	blog := Blog{}

	sqlErr := stmt.QueryRow(id).Scan(&blog.ID, &blog.Title, &blog.Description, &blog.P1, &blog.P2, &blog.P3, &blog.P4, &blog.P5, &blog.CS1, &blog.CS2, &blog.CS3, &blog.Repo, &blog.CreatedAt)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Blog{}, nil
		}
		return Blog{}, sqlErr
	}

	return blog, nil
}
