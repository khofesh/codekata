const { getMovies, addMovie } = require("./api-client");

console.log("getMovies:", getMovies);

async function callAPI() {
  //   const response = await getMovies();
  //   console.log("Res:", response);

  const response = await addMovie("Tom");
  console.log({ response });
}

callAPI();
