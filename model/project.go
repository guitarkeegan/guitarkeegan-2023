package models

import "database/sql"

type Project struct {
	Id    int     `json:"id"`
	Title string  `json:"title"`
	Des1  string  `json:"des1"`
	Des2  *string `json:"des2"`
	Des3  *string `json:"des3"`
	Live  *string `json:"live"`
	Repo  *string `json:"repo"`
}

func GetProjects(count int) ([]Project, error) {

	rows, err := DB.Query("SELECT id, title, des1, des2, des3, live, repo FROM project LIMIT ?", count)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	projects := make([]Project, 0)

	for rows.Next() {
		project := Project{}
		err := rows.Scan(&project.Id, &project.Title, &project.Des1, &project.Des2, &project.Des3, &project.Live, &project.Repo)

		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return projects, nil

}

func GetProjectById(id string) (Project, error) {
	stmt, err := DB.Prepare("SELECT id, title, des1, des2, des3, live, repo FROM project WHERE id = ?")

	if err != nil {
		return Project{}, err
	}

	project := Project{}

	sqlErr := stmt.QueryRow(id).Scan(&project.Id, &project.Title, &project.Des1, &project.Des2, &project.Des3, &project.Live, &project.Repo)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Project{}, nil
		}
		return Project{}, sqlErr
	}

	return project, nil

}
