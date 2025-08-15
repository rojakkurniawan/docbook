package services

import (
	"docbook/entity"
	"docbook/repository"
	"fmt"
)

type TimeslotService interface {
	CreateMultipleTimeslots(request *entity.CreateTimeslotRequest) error
	GetAllTimeslots() ([]entity.TimeslotResponse, error)
	GetTimeslotsWithFilter(filter *entity.TimeslotFilter) ([]entity.TimeslotResponse, error)
	GetTimeslotByID(id uint) (*entity.TimeslotResponse, error)
	UpdateTimeslot(id uint, timeslot *entity.TimeSlot) error
	DeleteTimeslot(id uint) error
}

type timeslotService struct {
	timeslotRepo repository.TimeslotRepository
}

func NewTimeslotService(timeslotRepo repository.TimeslotRepository) TimeslotService {
	return &timeslotService{timeslotRepo: timeslotRepo}
}

func (s *timeslotService) CreateMultipleTimeslots(request *entity.CreateTimeslotRequest) error {
	// Validasi doctor schedule exists
	_, err := s.timeslotRepo.GetDoctorScheduleByID(request.DoctorScheduleID)
	if err != nil {
		return err
	}

	// Buat multiple timeslots
	var timeslots []entity.TimeSlot
	for _, timeslotReq := range request.TimeSlots {
		timeslot := entity.TimeSlot{
			DoctorScheduleID: request.DoctorScheduleID,
			Date:             request.Date,
			StartTime:        timeslotReq.StartTime,
			EndTime:          timeslotReq.EndTime,
			IsAvailable:      timeslotReq.IsAvailable,
			IsBlocked:        timeslotReq.IsBlocked,
		}
		timeslots = append(timeslots, timeslot)
	}

	return s.timeslotRepo.CreateMultipleTimeslots(timeslots)
}

func (s *timeslotService) GetAllTimeslots() ([]entity.TimeslotResponse, error) {
	timeslots, err := s.timeslotRepo.GetAllTimeslots()
	if err != nil {
		return nil, err
	}

	// Group timeslots by doctor_schedule_id and date
	groupedTimeslots := make(map[string]entity.TimeslotResponse)

	for _, timeslot := range timeslots {
		key := fmt.Sprintf("%d-%s", timeslot.DoctorScheduleID, timeslot.Date)

		if response, exists := groupedTimeslots[key]; exists {
			// Add timeslot to existing group
			response.Timeslots = append(response.Timeslots, entity.TimeslotDetail{
				ID:          timeslot.ID,
				StartTime:   timeslot.StartTime,
				EndTime:     timeslot.EndTime,
				IsAvailable: timeslot.IsAvailable,
				IsBlocked:   timeslot.IsBlocked,
			})
			groupedTimeslots[key] = response
		} else {
			// Create new group
			groupedTimeslots[key] = entity.TimeslotResponse{
				DoctorScheduleID: timeslot.DoctorScheduleID,
				Date:             timeslot.Date,
				Timeslots: []entity.TimeslotDetail{
					{
						ID:          timeslot.ID,
						StartTime:   timeslot.StartTime,
						EndTime:     timeslot.EndTime,
						IsAvailable: timeslot.IsAvailable,
						IsBlocked:   timeslot.IsBlocked,
					},
				},
			}
		}
	}

	// Convert map to slice
	var result []entity.TimeslotResponse
	for _, response := range groupedTimeslots {
		result = append(result, response)
	}

	return result, nil
}

func (s *timeslotService) GetTimeslotsWithFilter(filter *entity.TimeslotFilter) ([]entity.TimeslotResponse, error) {
	timeslots, err := s.timeslotRepo.GetTimeslotsWithFilter(filter)
	if err != nil {
		return nil, err
	}

	// Group timeslots by doctor_schedule_id and date (same logic as GetAllTimeslots)
	groupedTimeslots := make(map[string]entity.TimeslotResponse)

	for _, timeslot := range timeslots {
		key := fmt.Sprintf("%d-%s", timeslot.DoctorScheduleID, timeslot.Date)

		if response, exists := groupedTimeslots[key]; exists {
			// Add timeslot to existing group
			response.Timeslots = append(response.Timeslots, entity.TimeslotDetail{
				ID:          timeslot.ID,
				StartTime:   timeslot.StartTime,
				EndTime:     timeslot.EndTime,
				IsAvailable: timeslot.IsAvailable,
				IsBlocked:   timeslot.IsBlocked,
			})
			groupedTimeslots[key] = response
		} else {
			// Create new group
			groupedTimeslots[key] = entity.TimeslotResponse{
				DoctorScheduleID: timeslot.DoctorScheduleID,
				Date:             timeslot.Date,
				Timeslots: []entity.TimeslotDetail{
					{
						ID:          timeslot.ID,
						StartTime:   timeslot.StartTime,
						EndTime:     timeslot.EndTime,
						IsAvailable: timeslot.IsAvailable,
						IsBlocked:   timeslot.IsBlocked,
					},
				},
			}
		}
	}

	// Convert map to slice
	var result []entity.TimeslotResponse
	for _, response := range groupedTimeslots {
		result = append(result, response)
	}

	return result, nil
}

func (s *timeslotService) GetTimeslotByID(id uint) (*entity.TimeslotResponse, error) {
	timeslot, err := s.timeslotRepo.GetTimeslotByID(id)
	if err != nil {
		return nil, err
	}

	result := &entity.TimeslotResponse{
		DoctorScheduleID: timeslot.DoctorScheduleID,
		Date:             timeslot.Date,
		Timeslots: []entity.TimeslotDetail{
			{
				ID:          timeslot.ID,
				StartTime:   timeslot.StartTime,
				EndTime:     timeslot.EndTime,
				IsAvailable: timeslot.IsAvailable,
				IsBlocked:   timeslot.IsBlocked,
			},
		},
	}

	return result, nil
}

func (s *timeslotService) UpdateTimeslot(id uint, timeslot *entity.TimeSlot) error {
	return s.timeslotRepo.UpdateTimeslot(id, timeslot)
}

func (s *timeslotService) DeleteTimeslot(id uint) error {
	return s.timeslotRepo.DeleteTimeslot(id)
}
