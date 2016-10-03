# wallabag-stats

Draws a chart for unread and total articles in your [Wallabag](https://github.com/wallabag/wallabag/) instance.

Chart is only created when there at least two data sets, in which unread and total are both different compared with their previous value. Otherwise, if the delta is zero, the go-chart lib does not print a chart.


## Links to chart libs

- https://github.com/wcharczuk/go-chart
- http://bl.ocks.org/mbostock/3943967 from https://github.com/d3/d3/wiki/Gallery


## Project Status
### Go Report Card

[![Go Report Card Badge](https://goreportcard.com/badge/github.com/Strubbl/wallabag-stats)](https://goreportcard.com/report/github.com/Strubbl/wallabag-stats)


### Travis CI

[![Build Status](https://travis-ci.org/Strubbl/wallabag-stats.svg?branch=master)](https://travis-ci.org/Strubbl/wallabag-stats)
