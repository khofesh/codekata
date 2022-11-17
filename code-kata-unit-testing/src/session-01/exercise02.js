export const getRandomNumber = (max) => {
  return Math.floor(Math.random() * max);
};

export const extractToken = (authToken) => {
  if (!authToken) {
    return "";
  }
  const parts = authToken.split(/\s+/);
  if (parts[0].toLowerCase() !== "bearer" || parts.length !== 2) {
    return "";
  }
  return parts[1];
};

export const inRange = (value, min, max) => {
  return value >= min && value <= max;
};

export const stringToBoolean = (str = "") => {
  return str.toLowerCase() === "true";
};
