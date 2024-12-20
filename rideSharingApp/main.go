package main

import (
	"fmt"
	"rideSharingApp/models"
	"rideSharingApp/services"
)

func main() {

	userService := services.NewUserService()
	rideService := services.NewRideService(userService)

	ramUser := userService.CreateUser("Ram")
	ramVehicle := models.NewVehicle("R0001", "Honda City", 4)
	err := userService.AddVehicle(ramUser.Id(), ramVehicle.Id())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ramRide, err := rideService.CreateRide(ramUser.Id(), "delhi", "noida", 1000, 1100, 4)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Ram ride is", ramRide)

	shyamUser := userService.CreateUser("Shyam")

	shyamVehicle := models.NewVehicle("R0002", "Honda City", 4)
	err = userService.AddVehicle(shyamUser.Id(), shyamVehicle.Id())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	shyamRide, err := rideService.CreateRide(shyamUser.Id(), "delhi", "noida", 1000, 1200, 4)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Shyam ride is", shyamRide)

	strategy := services.SelectionStrategy{
		IsShortestRideTimeStrategy: false,
		IsEarliestEndTimeStrategy:  true,
	}

	ride, err := rideService.SelectRides("delhi", "noida", strategy)
	if err != nil {
		return
	}
	fmt.Println("selected ride is ", ride)

}
