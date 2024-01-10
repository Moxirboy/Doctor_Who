package repository

func (r repo) CreatePhoto(id int, path []string) (err error) {
	query := `
	insert into photos(path ,owner_id) values($1,$2) 
	`
	for _, p := range path {
		_, err = r.db.Exec(query, p, id)
		if err != nil {
			r.Bot.SendErrorNotification(err)
			return err
		}
	}
	return nil
}
