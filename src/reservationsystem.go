package main

import "errors"

type BookingStatus int

const (
	PENDING BookingStatus = iota
	BOOKED                = iota
)

type Station struct {
	id     string
	trains []int32
}

// type Seat struct {
// 	id int32
// 	// bookedStations [][]string
// }

type Coach struct {
	id       int32
	numSeats int32
}

type Train struct {
	id            int32
	name          string
	source        string
	destination   string
	route         []string // list of stations
	coaches       []Coach
	departureTime int64
	arrivalTime   int64
}

type TrainManager interface {
	getTrainInfo(trainID int32) (*Train, error)
}

type TrainManagerImpl struct {
	trainsInfo map[int32]*Train
}

func NewTrainManager() TrainManager {
	return &TrainManagerImpl{
		trainsInfo: map[int32]*Train{},
	}
}

func (tm *TrainManagerImpl) getTrainInfo(trainID int32) (*Train, error) {
	if train, ok := tm.trainsInfo[trainID]; ok {
		return train, nil
	}

	return nil, errors.New("Train info is not available. Please check the train ID.")
}

type RailUser struct {
	id    int64
	name  string
	email string
	phone int64
}

type RailSeatBookingInfo struct {
	id             int64
	userName       string
	coachNo        int32
	seatNo         int32
	sourceStn      string
	destinationStn string
}

type RailBooking struct {
	id        int64
	userId    int64 // user who booked this ticket(s)
	users     []RailUser
	seats     []RailSeatBookingInfo
	status    BookingStatus
	journeyId int64
}

type Journey struct {
	id              int64
	trainID         int32
	departureTime   int64
	arrivalTime     int64
	currentBookings []RailBooking
	bookTicketChan  chan *RailBooking
	trainManager    TrainManager
}

func NewJourney(trainID int32, departureTime int64, arrivalTime int64, trainManager TrainManager) *Journey {
	journey := &Journey{
		trainID:         trainID,
		departureTime:   departureTime,
		arrivalTime:     arrivalTime,
		currentBookings: []RailBooking{},
		bookTicketChan:  make(chan *RailBooking, 500),
		trainManager:    trainManager,
	}

	go journey.listenToBookingEvents(journey.bookTicketChan)

	return journey
}

func (j *Journey) listenToBookingEvents(bookTicketChan chan *RailBooking) {
	for {
		booking, ok := <-bookTicketChan
		if ok {
			j.bookTicket(booking)
		} else {
			break
		}
	}
}

func (j *Journey) bookTicket(booking *RailBooking) error {
	// trainInfo, err := j.trainManager.getTrainInfo(j.trainID)
	// if err != nil {
	// 	return err
	// }

	// alreadyBookedSeats := map[int32]map[int32]bool{}
	// for _, existingBooking := range j.currentBookings {

	// }

	return nil
}

type RailwaySystem interface {
	getAvailableSeats(src string, dest string, date string) map[int32]int32
	bookTickets(userId int64, users []RailUser, trainID int32, date string) (RailBooking, error)
}

type RailSystemImpl struct {
	trains        []Train
	stations      []Station
	users         []User
	bookingsState map[int64]map[string]Journey
}
