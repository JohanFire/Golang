package service

import (
	"context"
	"fmt"

	pb "go-router/internal/protobuf"
)

// ClassificationService implements the business logic for classification
type ClassificationService struct{}

// NewClassificationService creates a new instance of the classification service
func NewClassificationService() *ClassificationService {
	return &ClassificationService{}
}

// ClassifyService performs the actual classification business logic
func (s *ClassificationService) ClassifyService(ctx context.Context, request *pb.ClassificationRequest) (*pb.ClassificationResponse, error) {

	// Implement business logic

	response := &pb.ClassificationResponse{
		Response: fmt.Sprintf(
			"Classification completed for: %s, your appliance is: %s, applianceVIB: %s",
			request.UserName,
			request.ApplianceType,
			request.ApplianceVIB,
		),
	}

	return response, nil
}
