import React from 'react';

class Command extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      items: []
    };
  }


  componentDidMount() {
    const commandName = window.location.pathname
    fetch(`http://localhost:3000/v1${commandName}`)
    .then(res => res.json())
    .then(
      (result) => {
        this.setState({
          isLoaded: true,
          items: result.command
        })
      },
      (error) => {
        this.setState({
          isLoaded: true,
          error
        });
      }
    )
  }

  render() {
    const { error, isLoaded, items } = this.state;
    //console.log({items})
    if (error) {
      return <div> Error: {error.message}</div>;
    } else if (!isLoaded) {
      return <div>Loading...</div>;
    } else {
      return  (
        <div>
          <p>
            Name: {items.name} <br></br>
            Text: {items.text}<br></br>
            Category: {items.category}<br></br>
            Level: {items.level}
          </p>
        </div>
      )
    }
  }
}

export default Command