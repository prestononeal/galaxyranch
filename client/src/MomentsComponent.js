import React, { Component } from 'react'
import axios from 'axios';
class MomentsComponent extends Component {

    constructor() {
        super();
        this.state = {
            moments: 'pending'
        }
    }

    componentWillMount() {
        axios.get('api/moments')
            .then((response) => {
                this.setState(() => {
                    return { moments: response.data.moments }
                })
            })
            .catch(function (error) {
                console.log(error);
            });

    }

    render() {
        return <div>{this.state.moments}</div>;
    }
}

export default MomentsComponent; 