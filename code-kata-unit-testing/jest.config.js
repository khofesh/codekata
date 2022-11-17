module.exports = {
  verbose: true,
  // setupFiles: ["<rootDir>/jest/setEnvVars.js"],
  collectCoverage: true,
  collectCoverageFrom: ["./src/**"],
  coverageThreshold: {
    global: {
      lines: 90,
    },
  },
};
