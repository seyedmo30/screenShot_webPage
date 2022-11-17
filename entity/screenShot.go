package entity

type ScreenShot struct {
	Id       uint64 `json:"id"`
	Url      string `json:"url"`
	FilePath string `json:"file_path"`
	Status   uint8  `json:"status"`
}
