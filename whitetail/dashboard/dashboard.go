package dashboard

type Graph struct {
	Observer   string   `yaml:"observer" json:"observer"`
	Stream     string   `yaml:"stream" json:"stream"`
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
}

type Stream struct {
	Source  string `yaml:"source" json:"source"`
	Name    string `yaml:"name" json:"name"`
	XCoord  int    `yaml:"x_coord" json:"x_coord"`
	YCoord  int    `yaml:"y_coord" json:"y_coord"`
	RowSpan int    `yaml:"row_span" json:"row_span"`
	ColSpan int    `yaml:"col_span" json:"col_span"`
	Refresh int    `yaml:"refresh" json:"refresh"`
	Title   string `yaml:"title" json:"title"`
}

type Table struct {
	Source  string `yaml:"source" json:"source"`
	Name    string `yaml:"name" json:"name"`
	XCoord  int    `yaml:"x_coord" json:"x_coord"`
	YCoord  int    `yaml:"y_coord" json:"y_coord"`
	RowSpan int    `yaml:"row_span" json:"row_span"`
	ColSpan int    `yaml:"col_span" json:"col_span"`
	Refresh int    `yaml:"refresh" json:"refresh"`
	Title   string `yaml:"title" json:"title"`
}

type Layout struct {
	Graphs  []Graph  `yaml:"graphs" json:"graphs"`
	Streams []Stream `yaml:"streams" json:"streams"`
	Tables  []Table  `yaml:"tables" json:"tables"`
}

type Dashboard struct {
	Layout      Layout  `yaml:"layout" json:"layout"`
	AspectRatio float64 `yaml:"aspect_ratio" json:"aspect_ratio"`
}
