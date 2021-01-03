# ForestFireOracle
"Predicts" forest fires based off of data that only sort of works because my partner and I didn't have resources access to realistic data sets. But a proof-enough of concept :)

So this isn't an entirely accurate tool. Partly because data required for accurate forest fire prediction isn't accessible to two college students who
aren't gonna spend money on that, since this is originally just a class project. Also a little hacky and not guaranteed to work because little details
got tripped up and we just had to duct tape everything together in order to be able to present it.

So this is an incomplete tool, but its main purpose of creation was proof of concept for general "have a web app that takes a request and gathers
a bunch of info from a database and then responds the prediction" type of project. I indend to make the code more modular and iron out the hacky parts,
and hopefully extend its functionality.

More likely than not, the Weather API that this thing was hooked up to you isn't working. I got a cheap subscription, although if this project were seriously
revamped, I'd get a reliable Weather API subscription to the service this project was using.

React frontent with Golang backend. Golang backend chosen so the big data set can be analysed and have data gathered in parallel, to get it done quicker.
Concurrent programming better for reading large sets of data.
