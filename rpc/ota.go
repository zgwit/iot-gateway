package rpc

type OtaVersionRequest struct{}

type OtaVersionResponse map[string]any

type OtaUploadRequest struct{}

type OtaUploadResponse StreamId

type OtaResultRequest struct {
	Version string `json:"version"`
}

type OtaResultResponse struct{}
