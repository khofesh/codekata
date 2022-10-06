const {
  getRandomNumber,
  extractToken,
  inRange,
  stringToBoolean,
} = require("../session-01/exercise02");

describe("exercise02", () => {
  it("should return random number", () => {
    const max = 15;

    const result = getRandomNumber(max);

    expect(typeof result).toEqual("number");
  });

  it("should return return token", () => {
    const token = "bearer someRandomString";

    const result = extractToken(token);

    expect(result).toEqual("someRandomString");
  });

  it("should return return empty string", () => {
    const token = "";

    const result = extractToken(token);

    expect(result).toEqual("");
  });

  it("should return return empty string 2", () => {
    const token = "bearer";

    const result = extractToken(token);

    expect(result).toEqual("");
  });

  it("should return boolean", () => {
    const value = 5;
    const min = 0;
    const max = 10;

    const result = inRange(value, min, max);

    expect(result).toEqual(true);
  });

  it("should return a string", () => {
    const someRandomString = "true";

    const result = stringToBoolean(someRandomString);

    expect(result).toEqual(true);
  });

  it("should return a false", () => {
    const result = stringToBoolean();

    expect(result).toEqual(false);
  });
});
