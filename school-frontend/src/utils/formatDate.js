export const formatDate = (isoString) => {
  const date = new Date(isoString);
  return date.toLocaleDateString("en-KE", {
    year: "numeric",
    month: "short",
    day: "numeric",
  });
};
