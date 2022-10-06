const fizzBuzz = require("../session-01/exercise01");

describe("exercise01", () => {
  it("should return FizzBuzz", () => {
    const number = 15;

    const result = fizzBuzz(number);

    expect(result).toEqual("FizzBuzz");
  });

  it("should return Fizz", () => {
    const number = 9;

    const result = fizzBuzz(number);

    expect(result).toEqual("Fizz");
  });

  it("should return Buzz", () => {
    const number = 10;

    const result = fizzBuzz(number);

    expect(result).toEqual("Buzz");
  });

  it("should return the number", () => {
    const number = 34;

    const result = fizzBuzz(number);

    expect(result).toEqual(number);
  });
});
