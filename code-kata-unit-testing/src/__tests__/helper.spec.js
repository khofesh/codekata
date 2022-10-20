const { Helper } = require("../session-02/helper");

describe("#helper", () => {
  it("should greet", () => {
    const instance = new Helper();
    instance.getFullName = jest.fn().mockReturnValue("rio tinto");
    // jest.spyOn(instance, "getFullName");
    const result = instance.greet("json", "born");

    console.log(result);
    expect(result).toEqual("Hello rio tinto");
  });
});
