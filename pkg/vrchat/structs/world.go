package structs

type (
	World struct {
		LimitedWorld
		Visits         int    `json:"visits"`
		Featured       bool   `json:"featured"`
		Description    string `json:"description"`
		AssetURL       string `json:"assetUrl"`
		AssetURLObject struct {
		} `json:"assetUrlObject"`
		PluginURL       string `json:"pluginUrl"`
		PluginURLObject struct {
		} `json:"pluginUrlObject"`
		UnityPackageURL       string `json:"unityPackageUrl"`
		UnityPackageURLObject struct {
		} `json:"unityPackageUrlObject"`
		Namespace        string        `json:"namespace"`
		Version          int           `json:"version"`
		PreviewYoutubeID interface{}   `json:"previewYoutubeId"`
		PublicOccupants  int           `json:"publicOccupants"`
		PrivateOccupants int           `json:"privateOccupants"`
		Instances        []interface{} `json:"instances"`
	}

	LimitedWorld struct {
		ID                  string   `json:"id"`
		Name                string   `json:"name"`
		AuthorID            string   `json:"authorId"`
		AuthorName          string   `json:"authorName"`
		Capacity            int      `json:"capacity"`
		ImageURL            string   `json:"imageUrl"`
		ThumbnailImageURL   string   `json:"thumbnailImageUrl"`
		ReleaseStatus       string   `json:"releaseStatus"`
		Organization        string   `json:"organization"`
		Tags                []string `json:"tags"`
		Favorites           int      `json:"favorites"`
		CreatedAt           string   `json:"created_at"`
		UpdatedAt           string   `json:"updated_at"`
		PublicationDate     string   `json:"publicationDate"`
		LabsPublicationDate string   `json:"labsPublicationDate"`
		UnityPackages       []struct {
			ID             string `json:"id"`
			AssetURL       string `json:"assetUrl"`
			AssetURLObject struct {
			} `json:"assetUrlObject"`
			PluginURL       string `json:"pluginUrl"`
			PluginURLObject struct {
			} `json:"pluginUrlObject"`
			UnityVersion    string `json:"unityVersion"`
			UnitySortNumber int64  `json:"unitySortNumber"`
			AssetVersion    int    `json:"assetVersion"`
			Platform        string `json:"platform"`
			CreatedAt       string `json:"created_at"`
		} `json:"unityPackages"`
		Popularity int `json:"popularity"`
		Heat       int `json:"heat"`
		Occupants  int `json:"occupants"`
	}
)

func (w World) GetName() string {
	if w.Name == "" {
		return "Private"
	}
	return w.Name
}
