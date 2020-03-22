package services

func (d *DB) GetAllDirectors() ([]Director, error) {

	directors := []Director{}
	err := d.GoDB.Find(&directors).Error
	if err != nil {
		return nil, err
	}

	return directors, nil
}

func (d *DB) GetOneDirector(id int) (*Actor, error) {
	actor := Actor{}

	err := d.GoDB.Where("id = ?", id).Find(&actor).Error
	if err != nil {
		return nil, err
	}

	return &actor, nil
}
