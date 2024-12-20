package services

import "rideSharingApp/models"

type rideService struct {
	rideMapping map[string]*models.Ride
	userService IUserService
}

type SelectionStrategy struct {
	IsShortestRideTimeStrategy bool
	IsEarliestEndTimeStrategy  bool
}

type IRideService interface {
	CreateRide(driverID string, origin string, destination string, startTime int, journeyTime int, availableSeats int) (*models.Ride, error)
	GetAvailableRides(origin string, destination string) []models.Ride
	SelectRides(origin string, destination string, strategy SelectionStrategy) (models.Ride, error)
}

func NewRideService(userService IUserService) IRideService {
	return &rideService{
		rideMapping: make(map[string]*models.Ride),
		userService: userService,
	}
}

func (rs *rideService) CreateRide(driverID string, origin string, destination string, startTime int, journeyTime int, availableSeats int) (*models.Ride, error) {
	//validation for driverID
	//validation for origin and destination
	ride, err := models.NewRide(origin, destination, startTime, journeyTime, availableSeats, driverID)
	if err != nil {
		return nil, err
	}

	rs.rideMapping[ride.Id()] = ride
	return ride, nil
}

func (rs *rideService) GetAvailableRides(origin string, destination string) []models.Ride {
	var availableRides []models.Ride
	for _, ride := range rs.rideMapping {
		if ride.Origin() == origin && ride.Destination() == destination && ride.Status() == "pending" && ride.AvailableSeats() >= 1 {
			availableRides = append(availableRides, *ride)
		}
	}

	return availableRides
}

func (rs *rideService) SelectRides(origin string, destination string, strategy SelectionStrategy) (models.Ride, error) {
	availableRides := rs.GetAvailableRides(origin, destination)

	if strategy.IsShortestRideTimeStrategy {
		var shortestRideTime = 100000000
		shortestRide := models.Ride{}
		for _, ride := range availableRides {
			if (ride.EndTime() - ride.StartTime()) < shortestRideTime {
				shortestRideTime = ride.EndTime() - ride.StartTime()
				shortestRide = ride
			}
		}
		return shortestRide, nil
	} else if strategy.IsEarliestEndTimeStrategy {
		var shortestEndTime = 100000000
		shortestEndTimeRide := models.Ride{}
		for _, ride := range availableRides {
			if ride.EndTime() < shortestEndTime {
				shortestEndTime = ride.EndTime()
				shortestEndTimeRide = ride
			}
		}

		return shortestEndTimeRide, nil
	}

	return models.Ride{}, nil
}
