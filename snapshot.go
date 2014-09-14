package image

// Snapshot is a DigitalOcean Snapshot/Backup.
type Snapshot struct {
	ID        int      `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Dist      string   `json:"distribution,omitempty"`
	Slug      string   `json:"slug,omitempty,omitempty"`
	Public    bool     `json:"public,omitempty"`
	Regions   []string `json:"regions,omitempty"`
	ActionIDs []int    `json:"action_ids,omitempty"`
}
