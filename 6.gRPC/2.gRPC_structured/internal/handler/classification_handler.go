package handler

import (
	"context"
	"log"

	pb "go-router/internal/protobuf"
	"go-router/internal/service"
)

// ClassificationHandler implements the gRPC server interface
type ClassificationHandler struct {
	pb.UnimplementedClassificationServiceServer
	service *service.ClassificationService
}

// NewClassificationHandler creates a new instance of ClassificationHandler
func NewClassificationHandler(service *service.ClassificationService) *ClassificationHandler {
	return &ClassificationHandler{
		service: service,
	}
}

// Classify handles the gRPC Classify method
func (h *ClassificationHandler) Classify(ctx context.Context, request *pb.ClassificationRequest) (*pb.ClassificationResponse, error) {
	log.Printf("Received request: %v", request)

	// Delegate to the service layer
	return h.service.ClassifyService(ctx, request)
}
