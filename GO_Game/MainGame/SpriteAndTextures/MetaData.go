package SpriteAndTextures

type MetaData struct {
	name               string
	moveable           bool
	speed              int16
	health             int8
	hitpoints          int32
	max_hitpoints      int32
	hit_damage         int16
	position           []float64
	size               []float64
	current_image_path string
	animated           bool
}

func CreateMetaData(name string, moveable bool, speed int16, health int8, hitpoints int32, hit_damage int16, position []float64, scale []float64, image_path string, animated bool) MetaData {
	return MetaData{name, moveable, speed, health, hitpoints, hitpoints, hit_damage, position, scale, image_path, animated}
}

func CreateStaticMetaData(name string, health int8, hitpoints int32, hit_damage int16, position []float64, scale []float64, image_path string) MetaData {
	return MetaData{name, false, 0, health, hitpoints, hitpoints, hit_damage, position, scale, image_path, false}
}

func CreateDynamicMetaData(name string, speed int16, health int8, hitpoints int32, hit_damage int16, position []float64, scale []float64, image_path string) MetaData {
	return MetaData{name, true, speed, health, hitpoints, hitpoints, hit_damage, position, scale, image_path, true}
}

// Setters
func (s *MetaData) MovePosition(new_position []float64) {
	if s.moveable {
		s.position = new_position
	}
}

func (s *MetaData) MoveBy(dx, dy float64) {
	if s.moveable {
		s.position[0] += dx / s.size[0]
		s.position[1] += dy / s.size[1]
	}
}

func (s *MetaData) DecrementHitPoints(decrementer int32) {
	s.hitpoints -= decrementer
	if s.hitpoints < 0 {
		s.hitpoints = 0
	}
	s.health = int8((s.hitpoints / s.max_hitpoints) * 100)
}

func (s *MetaData) ChangeSize(new_size []float64) {
	s.size = new_size
}

func (s *MetaData) SetName(name string) {
	s.name = name
}

// Getters
func (s *MetaData) GetName() string {
	return s.name
}

func (s *MetaData) IsMoveable() bool {
	return s.moveable
}

func (s *MetaData) GetSpeed() int16 {
	return s.speed
}

func (s *MetaData) GetHealth() int8 {
	return s.health
}

func (s *MetaData) GetHitpoints() int32 {
	return s.hitpoints
}

func (s *MetaData) GetMaxHitpoints() int32 {
	return s.max_hitpoints
}

func (s *MetaData) GetHitDamage() int16 {
	return s.hit_damage
}

func (s *MetaData) GetPosition() []float64 {
	return append([]float64(nil), s.position...)
}

func (s *MetaData) GetSize() []float64 {
	return append([]float64(nil), s.size...)
}

func (s *MetaData) GetCurrentImagePath() string {
	return s.current_image_path
}

func (s *MetaData) IsAnimated() bool {
	return s.animated
}
