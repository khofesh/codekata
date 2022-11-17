const helper = require("./helper");

describe("# greet", () => {
  it("should return Hello Luis Puentes when Luis and Puentes are sent as parameters", () => {
    // arrange
    const firstName = "Luis";
    const lastName = "Puentes";
    const expected = `Hello ${firstName} ${lastName}`;
    helper.getFullName = jest.fn().mockReturnValue(`${firstName} ${lastName}`);

    // act
    const result = helper.greet(firstName, lastName);

    //assert
    expect(result).toBe(expected);
  });
});
