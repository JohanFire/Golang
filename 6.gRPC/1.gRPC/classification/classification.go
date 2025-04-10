package classification

import (
	"context"
	"fmt"
	"log"
	// pb "1.gRPC/classification" // in this case import cycle not allowed
)

type Server struct {
	UnimplementedClassificationServiceServer
	// pb.UnimplementedClassificationServiceServer
}

func (s *Server) Classify(ctx context.Context, request *ClassificationRequest) (*ClassificationResponse, error) {
	log.Printf("Received request: %v", request)
	// log.Printf("Name: %s, Appliance: %s, ApplianceVIB: %s", request.Name, request.Appliance, request.ApplianceVIB)

	// Simulate classification logic
	response := &ClassificationResponse{
		// Response: "Classified finished for: " + request.Name,
		Response: fmt.Sprintf(
			"Classified finished for: %s, your appliance is: %s, applianceVIB: %s",
			request.Name,
			request.Appliance,
			request.ApplianceVIB,
		),
	}
	return response, nil
}
