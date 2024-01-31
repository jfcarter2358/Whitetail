package dashboard

type Panel struct {
	Kind       string   `yaml:"kind" json:"kind"`
	Name       string   `yaml:"name" json:"name"`
	Source     string   `yaml:"source" json:"source"`
	XCoord     int      `yaml:"x_coord" json:"x_coord"`
	YCoord     int      `yaml:"y_coord" json:"y_coord"`
	RowSpan    int      `yaml:"row_span" json:"row_span"`
	ColSpan    int      `yaml:"col_span" json:"col_span"`
	Refresh    int      `yaml:"refresh" json:"refresh"`
	X          string   `yaml:"x" json:"x"`
	Ys         []string `yaml:"ys" json:"ys"`
	XAxisLabel string   `yaml:"x_axis_label" json:"x_axis_label"`
	YAxisLabel string   `yaml:"y_axis_label" json:"y_axis_label"`
	YLabels    []string `yaml:"y_labels" json:"y_labels"`
	Title      string   `yaml:"title" json:"title"`
	Colors     []string `yaml:"colors" json:"colors"`
	Width      int      `yaml:"width" json:"width"`
	Height     int      `yaml:"height" json:"height"`
	Callback   string   `yaml:"callback" json:"callback"`
	JS         string   `yaml:"js" json:"js"`
}

type Dashboard struct {
	Panels      []Panel `yaml:"panels" json:"panels"`
	AspectRatio float64 `yaml:"aspect_ratio" json:"aspect_ratio"`
}
