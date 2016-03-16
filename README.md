# GPXGenerator
GO utility to process GPX files

## Overview
This utility allows to add time to the trackpoints in the GPX file (created from Google Map for example) in order to simulate trip advancement in time.

## Usage
1. Open Google Map and create desired trip using "Trip Directions" function
2. Open gmap2xml.sh script and replace values in FROM and TO with yours which can be copied from "Trip Directions"
3. Run gmap2xml.sh script in order to generate Google Map XML file
4. Download GPSBabel from http://www.gpsbabel.org/
5. Start GPSBabel and select generated in #3 file as input
6. Use this utility to adjust generated in #5 GPX in order to have time stamps in tracks.
 
## Run
$ go run main.go inputFile duration [outputFile]

where:
* inputFile - the input GPX file
* duration - the emulated route duration in minutes
* outputFile - the path to the output file (optional). If missed than name of input will be used by appending .out suffix

