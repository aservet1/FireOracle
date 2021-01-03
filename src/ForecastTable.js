import React from 'react'
import ReactFlexyTable from "react-flexy-table"
import "react-flexy-table/dist/index.css"
import axios from 'axios'

class ForecastTable extends React.Component {

    constructor(props) {
        super(props)
        
        this.county = 'none';
        this.data = [
          { Day: 'n/a', FirePrediction: 'n/a'},
        ]

        this.handleCountyEnter = this.handleCountyEnter.bind(this)
        this.handleCountyBoxChange = this.handleCountyBoxChange.bind(this)
    }

    handleCountyBoxChange(e){
        this.county = e.target.value
    }

    handleCountyEnter(e) {
      e.preventDefault()
      axios.get(`/api/forestFireForecast/${this.county}`)
        .then(res => {
            console.log(res);
            console.log(res.data);
            console.log(res.data.fireforecast)
            let result = res.data.fireforecast
            let today = new Date()
            this.data = []
            for (let i = 0; i < result.length; i++)
              this.data.push({Day: (today.getMonth()+1)+"/"+(today.getDate()+(i)) , FirePrediction: result[i]} )
            this.setState({});
        }).catch(e => {
          console.log('error')
          console.log(e)
          let result = [{Day: "an error occurred", FirePrediction: false}]
          let today = new Date()
          this.data = []
          for (let i = 0; i < result.length; i++)
            this.data.push({Day: (today.getMonth()+1)+"/"+(today.getDate()+(i)) , FirePrediction: result[i]} )
          this.setState({});
        })
    }

    render() {
      return (
        <div className="ForecastTable">
           <h1>Forest Fire Predicting App</h1>
           <br></br>
           <br></br>
           <br></br>
           What is your county?
          <input
            type="text"
            onChange={this.handleCountyBoxChange}
          />
          <input 
            type="button"
            value="Submit"
            onClick={this.handleCountyEnter}
          /> 
        <div>
           <br></br>
           <br></br>
           <br></br>
          Forecast Table for {this.county}
          <ReactFlexyTable data={this.data}/>
        </div>
        </div>
      )
    }
}

export default ForecastTable;