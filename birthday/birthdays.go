package birthday

type Birthday struct {
	Month  string `json: m bson: m`
	Date   int8   `json: d bson: d`
	Name   string `json: name bson: name`
	Mobile int64  `josn: m bson: m`
}
