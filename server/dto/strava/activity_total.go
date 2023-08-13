package strava

type ActivityStats struct {
	BiggestRideDistance       float64 `json:"biggest_ride_distance,omitempty"`
	BiggestClimbElevationGain float64 `json:"biggest_climb_elevation_gain,omitempty"`
}
