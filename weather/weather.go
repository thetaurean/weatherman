package weather

import "time"

type Weather struct {
	Temp      float32
	Humidity  float32
	Windspeed float32
	Created   time.Time
}

func (w Weather) Age() string {
	return time.Now().Sub(w.Created).String()
}
