#!/bin/sh 
FROM="Perova Blvd, 23, Kyiv" 
TO="Antonovycha St, 62, Kyiv" 
wget -O - "http://maps.google.com/maps?q=$FROM to $TO&output=js" \
2&>/dev/null >google_map.js.xml
