package basestation

import (
	"whitetail/analyst"
	"whitetail/datastore"
	"whitetail/operation"
)

func ReceiveData(n, o, d string) error {
	s := operation.Operations.Observers[o].Streams[n]
	a := operation.Operations.Analysts[s.Analyst]
	output, err := analyst.DoAnalyze(a, n, o, d, s.Arguments)
	if err != nil {
		return err
	}
	return datastore.AddData(output, n, o)
}

func GetData(s, o string, filter interface{}) ([]interface{}, error) {
	data, err := datastore.GetData(s, o, filter)
	if err != nil {
		return nil, err
	}
	return data, nil
}
