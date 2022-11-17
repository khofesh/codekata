const axios = require("axios");

const { getMovies, addMovie } = require("./api-client");

jest.mock("axios");
jest.mock("../session-02/helper.js", () => ({
  getEnvs: jest.fn().mockReturnValue({ API: "url" }),
}));

describe("# getMovies", () => {
  it("should return users list when axios response is successfully", async () => {
    // Your code here...
    // arrange
    const body = {
      data: [
        { id: 1, name: "movie 1" },
        { id: 2, name: "movie 2" },
      ],
    };
    axios.get.mockResolvedValueOnce(body);

    // act
    const result = await getMovies();

    console.log(result);

    // assert
    expect(axios.get).toHaveBeenCalledWith("url/api/movies");
    expect(result).toEqual(body.data);
  });

  it("should return error  when there is one error in the endpoint request", async () => {
    // Your code here...
  });
});
