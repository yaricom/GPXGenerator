package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"strconv"
)

// The routine to amend GPX route definition file
type GPX struct {
	XMLName	xml.Name 	`xml:"http://www.topografix.com/GPX/1/1 gpx"`
	Meta Metadata 		`xml:"metadata"`
	Trk Track 			`xml:"trk"`
}

type Metadata struct {
	Time time.Time 		`xml:"time"`
}

type Track struct {
	Name string			`xml:"name"`
	Desc string			`xml:"desc"`
	Segment TrackSegment `xml:"trkseg"`
}

type TrackSegment struct {
	TrackPoints []TrackPoint	`xml:"trkpt"`
}

type TrackPoint struct {
	Latitude float64	`xml:"lat,attr,omitempty"`
	Longitude float64   `xml:"lon,attr,omitempty"`
	Name string         `xml:"name"`
	Time time.Time      `xml:"time"`
}

func main()  {
	if len(os.Args) < 3 {
		printHelp()
		os.Exit(0)
	}
	var outputFile string
	inputFile := os.Args[1]
	var duration, err = strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Wrong duration provided: %s\n", os.Args[2])
		printHelp()
		os.Exit(0)
	}

	if len(os.Args) > 3 {
		outputFile = os.Args[3]
	} else {
		outputFile = inputFile + ".out.xml"
	}
	fmt.Printf("Input: %s, duration: %d minutes, output: %s\n", inputFile, duration, outputFile)

	// read XML
	gpx := GPX{}
	xmlContent, _ := ioutil.ReadFile(inputFile)
	err = xml.Unmarshal(xmlContent, &gpx)
	if err != nil { panic(err) }

	meta := gpx.Meta
	fmt.Printf("Start time: %s\n", meta.Time)

	trkpnts := gpx.Trk.Segment.TrackPoints
	delta := int32(duration * 60 * 1000 / len(trkpnts))
	fmt.Printf("Trackpoints count: %d, time delta: %d milliseconds\n", len(trkpnts), delta)
	
	// set times
	var pointTime = meta.Time
	for i := range trkpnts  {
		point := &trkpnts[i]
		point.Time = pointTime
		pointTime = pointTime.Add(time.Duration(delta) * time.Millisecond)
	}

	for _, point := range trkpnts  {
		fmt.Println(point)
	}

	// write output
	data, err := xml.Marshal(gpx)
	if err == nil {
		err = ioutil.WriteFile(outputFile, data, 0644)
		if err != nil {
			fmt.Println("Failed to save output: " + err.Error() )
		} else {
			fmt.Println("\n\nOperation complete successfully!")
		}
	}
}

func printHelp()  {
	fmt.Println("Arguments:")
	fmt.Println("inputFile - the path to the input file")
	fmt.Println("duration - the emulated route duration in minutes")
	fmt.Println("outputFile - the path to the output file (optional). If missed than name of input will be used by appending .out suffix")
}