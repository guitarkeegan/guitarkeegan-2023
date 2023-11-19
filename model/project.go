package models

type Project struct {
	Id    int     `json:"id"`
	Title string  `json:"title"`
	Des1  string  `json:"des1"`
	Des2  string  `json:"des2"`
	Des3  string  `json:"des3"`
	Live  *string `json:"live"`
	Repo  *string `json:"repo"`
}

func GetProjects() ([]Project, error) {

	rows, err := DB.Query("SELECT title, des1, des2, des3, live, repo FROM project")

	defer rows.Close()

	if err != nil {
		return nil, err
	}

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
