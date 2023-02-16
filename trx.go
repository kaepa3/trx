package trx

import (
	"encoding/xml"
	"os"
)

type Result struct {
	UnitTestResults []struct {
		TestName string `xml:"testName,attr"`
		Outcome  string `xml:"outcome,attr"`
	} `xml:"UnitTestResult"`
}

type ResultSummary struct {
	Conters struct {
		Total              string `xml:"total,attr"`
		Executed           string `xml:"executed,attr"`
		Failed             string `xml:"failed,attr"`
		Error              string `xml:"error,attr"`
		Timeout            string `xml:"timeout,attr"`
		Aborted            string `xml:"aborted,attr"`
		Inconclusive       string `xml:"Iinconclusivenc,attr"`
		PassedButRunAborte string `xml:"passedButRunAborte,attr"`
		NotRunnable        string `xml:"notRunnable,attr"`
		Disconnected       string `xml:"disconnected,attr"`
		Warning            string `xml:"warning,attr"`
		Completed          string `xml:"completed,attr"`
		InProgress         string `xml:"inProgress,attr"`
		Pending            string `xml:"pending,attr"`
	} `xml:"Counters"`
}

type Trx struct {
	Times         string   `xml:"Times,start"`
	Results       []Result `xml:"Results"`
	ResultSummary ResultSummary
}

func Parse(path string) (*Trx, error) {

	r, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	var rec Trx
	dec := xml.NewDecoder(r)
	err = dec.Decode(&rec)
	if err != nil {
		return nil, err
	}
	return &rec, err
}
