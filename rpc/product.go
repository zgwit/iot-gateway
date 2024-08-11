package rpc

type ProductId struct {
	Id string `json:"id"`
}

type ProductItem struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ProductStream struct {
	Stream uint16 `json:"stream"`
}

type ProductListRequest struct{}

type ProductListResponse []ProductItem

type ProductDownloadRequest ProductId

type ProductDownloadResponse ProductStream

type ProductUploadRequest ProductId

type ProductUploadResponse ProductStream

type ProductDeleteRequest ProductId

type ProductDeleteResponse struct{}
