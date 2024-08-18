package rpc

const (
	FS_LIST uint8 = iota
	FS_STATE
	FS_DOWNLOAD
	FS_UPLOAD
	FS_REMOVE
	FS_MOVE
	FS_MKDIR
	FS_FORMAT
)

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

type FsSearchRequest FsPath

type FsSearchResponse []FsItem

type FsDownloadRequest FsPath

type FsDownloadResponse StreamId

type FsUploadRequest FsPath

type FsUploadResponse StreamId

type FsDeleteRequest FsPath

type FsDeleteResponse struct{}

type FsMoveRequest FsPathMove

type FsMoveResponse struct{}

type FsMakeDirectoryRequest FsPath

type FsMakeDirectoryResponse struct{}

type FsFormatRequest struct {
	Disk string `json:"disk"`
	Type string `json:"type"`
}

type FsFormatResponse struct{}
