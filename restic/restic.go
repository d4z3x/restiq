package restic

type ResticBackUp struct {
	MessageType    string   `json:"message_type"`
	SecondsElapsed int      `json:"seconds_elapsed"`
	PercentDone    float64  `json:"percent_done"`
	TotalFiles     int      `json:"total_files"`
	FilesDone      int      `json:"files_done"`
	TotalBytes     int      `json:"total_bytes"`
	BytesDone      int      `json:"bytes_done"`
	CurrentFiles   []string `json:"current_files,omitempty"`
}
