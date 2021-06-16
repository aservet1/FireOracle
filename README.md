# ForestFireOracle

Proof-of-concept web app that predicts forest forest fires. Enter a location, will return predictions for the next 5 days telling whether or not a fire is likely to happen on that day, based on the weather forecast.

This project was made with one other teammate for a college class project, where we had to implement an idea with GoLang and concurrent programming. The concurrent programming comes in when we make predictions by pulling data from a record of [1.88 million recorded forest fires](https://www.kaggle.com/rtatman/188-million-us-wildfires). Searching for and pulling relevant data in parallel with Go-routines made predictions faster.

![image](https://user-images.githubusercontent.com/54599694/122151091-e7eec580-ce2c-11eb-8891-9520e2a71fb1.png)

## Technologies Used
- GoLang Backend
- React.js Frontend 
- SQLite Database

## Limitations
- Short time window. This was a class project, so we didn't have time for in depth research about how to make predictive software.
- Limited data access. While we pull real weather forecast data and compare it with real fire data, we could've paid for a more complex climate forecast, and for location-based historical weather data. Instead, we have a simple weather forecast and randomly generate placeholder historical data.

Because of the limitations, the predictions are not accurate. This project, however showcases our ability to build an application with GoLang and React, both in-demand technologies, and had this been a real-world project with access to real-world resources, it would have been an overall success.
