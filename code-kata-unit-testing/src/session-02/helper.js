require("dotenv").config();

class Helper {
  //   constructor(firstName, lastName) {
  //     this.firstName = firstName;
  //     this.lastName = lastName;
  //   }

  greet(firstName, lastName) {
    return `Hello ${this.getFullName(firstName, lastName)}`;
  }

  getFullName(firstName, lastName) {
    return `${firstName} ${lastName}`;
  }

  getRandomNumber() {
    return Math.random();
  }

  getEnvs() {
    return {
      API: process.env.API,
    };
  }
}

/*
TODO: 
1. add unit testing to greet() function and mock getFullName fn
2. add unit to getRandomNumber() function and mock Math fn
*/

module.exports = {
  Helper,
};
