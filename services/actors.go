package services

func (d *DB) GetAllActors() ([]Actor, error) {

	actors := []Actor{}
	err := d.GoDB.Find(&actors).Error
	if err != nil {
		return nil, err
	}

	return actors, nil
}

func (d *DB) GetOneActor(id int) (*Actor, error) {
	actor := Actor{}

	err := d.GoDB.Where("id = ?", id).Find(&actor).Error
	if err != nil {
		return nil, err
	}

	return &actor, nil
}
