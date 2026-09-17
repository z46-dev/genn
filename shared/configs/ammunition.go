package configs

var (
	Bullet = &Definition{
		ID:   1,
		Type: TypeBullet,
		Name: "Bullet",
	}

	Drone = &Definition{
		ID:   2,
		Type: TypeBullet,
		Name: "Drone",
	}

	Ammunition = []*Definition{
		Bullet,
		Drone,
	}
)
