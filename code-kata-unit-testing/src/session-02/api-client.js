const axios = require("axios");
const helper = require("./helper");

const API = `${helper.getEnvs().API}/api`;

const getMovies = async () => {
  return axios
    .get(`${API}/movies`)
    .then((res) => {
      return res.data;
    })
    .catch((e) => {
      return [];
    });
};

const addMovie = async (name) => {
  return axios
    .post(`${API}/movies`, {
      name,
    })
    .then((res) => {
      return res.data;
    })
    .catch((e) => {
      return "something went wrong";
    });
};

/*
TODO: 
1. add unit testing to getMovies() function and mock axios.get call
  → test case when successfully response
  → test case when error in the response
2. add unit to addMovie() function  and mock axios.post call
  → test case when successfully response
  → test case when error in the response
*/

module.exports = { addMovie, getMovies };
