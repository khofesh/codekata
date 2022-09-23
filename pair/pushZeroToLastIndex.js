function pushAllZeroToLastIndex(num) {
  let nonZeros = [];
  let zeros = [];

  const len = num.length;

  for (let index = 0; index < num.length; index++) {
    // const element = num[index];
    if (num[index] === 0) {
      zeros.push(num[index]);
    } else {
      nonZeros.push(num[index]);
    }
  }

  return [...nonZeros, ...zeros];
}

(() => {
  const num = [1, 2, 5, 6, 7, 0, 0, 54];
  const num2 = [1, 2, 0, 0, 0, 3, 6];

  console.log("result1: ", pushAllZeroToLastIndex(num));
  console.log("result2: ", pushAllZeroToLastIndex(num2));
})();
