package rpc

type ProductId struct {
	Id string `json:"id"`
}

type ProductItem struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ProductListRequest struct{}

type ProductListResponse []ProductItem

type ProductDownloadRequest ProductId

type ProductDownloadResponse StreamId

type ProductUploadRequest ProductId

type ProductUploadResponse StreamId

type ProductDeleteRequest ProductId

type ProductDeleteResponse struct{}
