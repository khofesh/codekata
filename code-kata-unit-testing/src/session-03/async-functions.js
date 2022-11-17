const fetch = (url) => {
  return new Promise((resolve, reject) => {
    if (url) resolve("Hello Async");
    else reject("There is no URL");
  });
};

function addAsync(a, b, callback) {
  setTimeout(() => {
    const result = a + b;
    callback(result);
  }, 500);
}

module.exports = {
  fetch,
  addAsync,
};
