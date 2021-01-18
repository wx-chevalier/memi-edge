'use strict';

const React = require('react');

class Greeting extends React.Component {
  render () {
    return React.createElement('div', null, 'hello ' + this.props.name);
  }
}

module.exports = Greeting;
