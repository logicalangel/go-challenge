package services

// this service provide Send user segments ability to uss app
type EstimateService interface {
	SendUserSegment(username, segment string) error
}
