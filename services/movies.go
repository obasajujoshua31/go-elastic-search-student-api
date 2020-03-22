package services

func (d *DB) GetAllMovies() ([]Movie, error) {

	movies := []Movie{}
	err := d.GoDB.Find(&movies).Error
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (d *DB) GetOneMovie(id int) (*Movie, error) {
	movie := Movie{}

	err := d.GoDB.Where("id = ?", id).Find(&movie).Error
	if err != nil {
		return nil, err
	}

	return &movie, nil
}
