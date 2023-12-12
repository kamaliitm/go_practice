package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type VehicleType int

const (
	CAR   VehicleType = iota
	BIKE              = iota
	TRUCK             = iota
)

const (
	ERROR_INVALID_TICKET = "invalid ticket id"
)

type Vehicle struct {
	vehicleType        VehicleType
	registrationNumber string
	color              string
}

type Slot struct {
	vehicleType VehicleType
	ticketId    string
}

type Floor struct {
	slots []Slot
}

type ParkingLot interface {
	// addFloors([]Floor) error
	// addSlot(floorIndex int32, vehicleType VehicleType) error
	parkVehicle(vehicle Vehicle) (string, error)
	unparkVehicle(ticketId string) error
	showFreeSlotCountPerFloor(vehicleType VehicleType)
	showFreeSlotsPerFloor(vehicleType VehicleType)
	showOccupiedSlotsPerFloor(vehicleType VehicleType)
}

type ParkingLotImpl struct {
	id     string
	floors []Floor
}

func NewParkingLot(id string, numFloors int32, numSlotsPerFloor int32) ParkingLot {
	parkingLot := ParkingLotImpl{
		id:     id,
		floors: []Floor{},
	}

	for i := 0; i < int(numFloors); i++ {
		floor := Floor{
			slots: []Slot{},
		}

		floor.slots[0] = Slot{vehicleType: TRUCK}
		floor.slots[1] = Slot{vehicleType: BIKE}
		floor.slots[2] = Slot{vehicleType: BIKE}
		for j := 3; j < int(numSlotsPerFloor); j++ {
			floor.slots[j] = Slot{vehicleType: CAR}
		}

		parkingLot.floors = append(parkingLot.floors, floor)
	}

	return &parkingLot
}

func (p *ParkingLotImpl) parkVehicle(vehicle Vehicle) (string, error) {
	if p == nil || len(p.floors) == 0 {
		return "", errors.New("create parking lot first")
	}

	for f, floor := range p.floors {
		for s, slot := range floor.slots {
			if slot.vehicleType == vehicle.vehicleType && slot.ticketId != "" {
				slot.ticketId = fmt.Sprintf("%s_%d_%d", p.id, f+1, s+1)
				p.floors[f].slots[s] = slot
				return slot.ticketId, nil
			}
		}
	}

	return "", errors.New("parking full")
}

func (p *ParkingLotImpl) unparkVehicle(ticketId string) error {
	if p == nil || len(p.floors) == 0 {
		return errors.New("create parking lot first")
	}

	ticketInfoTokens := strings.Split(ticketId, "_")
	if len(ticketInfoTokens) < 3 {
		return errors.New(ERROR_INVALID_TICKET)
	}

	floorNo, err := strconv.Atoi(ticketInfoTokens[1])
	if err != nil || floorNo <= 0 || floorNo >= len(p.floors) {
		return errors.New(ERROR_INVALID_TICKET)
	}

	slotNo, err2 := strconv.Atoi(ticketInfoTokens[2])
	if err2 != nil || slotNo <= 0 || slotNo >= len(p.floors[floorNo-1].slots) {
		return errors.New(ERROR_INVALID_TICKET)
	}

	slot := p.floors[floorNo-1].slots[slotNo-1]
	slot.ticketId = ""
	p.floors[floorNo-1].slots[slotNo-1] = slot

	return nil
}

func (p *ParkingLotImpl) showFreeSlotCountPerFloor(vehicleType VehicleType) {

}

func (p *ParkingLotImpl) showFreeSlotsPerFloor(vehicleType VehicleType) {

}

func (p *ParkingLotImpl) showOccupiedSlotsPerFloor(vehicleType VehicleType) {

}
