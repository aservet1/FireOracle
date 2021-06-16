# ForestFireOracle

Proof-of-concept web app that predicts forest forest fires. Enter a location, will return predictions for the next 5 days telling whether or not a fire is likely to happen on that day, based on the weather forecast.

This project was made with one other teammate for a college class project, where we had to implement an idea with GoLang and concurrent programming. The concurrent programming comes in when we make predictions by pulling data from a record of [1.88 million recorded forest fires](https://www.kaggle.com/rtatman/188-million-us-wildfires). Searching for and pulling relevant data in parallel with Go-routines made predictions faster.

## Technologies Used
- GoLang Backend
- React.js Frontend 
- SQLite Database

## Disclaimer
Although this web app "predicts" forest fires, this should not be used as any reliable source of predictions. Since it was a relatively short-term class project, and the assignment was to implement something cool with GoLang, my partner and I didn't have an in depth background on what exactly the type of data one would need to make accurate predictions, and some of the data that would've been useful was behind a pay wall.

For example, we used a simple weather forecast API, but could've paid for a climate forecast with more in depth data. We also could've paid for access to historical weather data (to associate the forecast with weather at the time and place of a fire incident), but randomly generated the historical data instead.

The reason I pinned this project anyways is to experience with GoLang and React.js. I'm happy about how the project went, outside of the issues in data acquisition, and it's proof of concept building a project with contemporary and in demand technologies.
