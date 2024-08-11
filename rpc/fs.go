package rpc

type FsPath struct {
	Path string `json:"path"`
}

type FsPathMove struct {
	Path string `json:"path"`
	Move string `json:"move"`
}

type FsItem struct {
	Name string `json:"name"`
	Dir  bool   `json:"dir,omitempty"`
	Size int64  `json:"size,omitempty"`
	Time int64  `json:"time,omitempty"`
}

type FsStream struct {
	Stream uint16 `json:"stream"`
}

type FsSearchRequest FsPath

type FsSearchResponse []FsItem

type FsDownloadRequest FsPath

type FsDownloadResponse FsStream

type FsUploadRequest FsPath

type FsUploadResponse FsStream

type FsDeleteRequest FsPath

type FsMoveRequest FsPathMove

type FsMakeDirectoryRequest FsPath

type FsFormatRequest struct {
	Disk string `json:"disk"`
	Type string `json:"type"`
}
