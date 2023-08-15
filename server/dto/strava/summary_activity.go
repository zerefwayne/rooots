package strava

import "time"

type LatLng []float64

type MetaAthlete struct {
	Id int64 `json:"id,omitempty"`
}

type PolylineMap struct {
	Id              string `json:"id,omitempty"`
	Polyline        string `json:"polyline,omitempty"`
	SummaryPolyline string `json:"summary_polyline,omitempty"`
}

type SummaryActivity struct {
	Id                 int64       `json:"id,omitempty"`
	Athlete            MetaAthlete `json:"athlete,omitempty"`
	Name               string      `json:"name,omitempty"`
	Distance           float64     `json:"distance,omitempty"`
	MovingTime         int64       `json:"moving_time,omitempty"`
	ElapsedTime        int64       `json:"elapsed_time,omitempty"`
	TotalElevationGain float64     `json:"total_elevation_gain,omitempty"`
	SportType          string      `json:"sport_type,omitempty"` // TODO Check enum implementation
	StartDateLocal     time.Time   `json:"start_date_local,omitempty"`
	Timezone           string      `json:"timezone,omitempty"`
	StartLatLng        LatLng      `json:"start_lat_lng,omitempty"`
	EndLatLng          LatLng      `json:"end_lat_lng,omitempty"`
	Map                PolylineMap `json:"map,omitempty"`
	Commute            bool        `json:"commute,omitempty"`
	Private            bool        `json:"private,omitempty"`
	WorkoutType        int64       `json:"workout_type,omitempty"`
	AverageSpeed       float64     `json:"average_speed,omitempty"`
	MaxSpeed           float64     `json:"max_speed,omitempty"`
}
