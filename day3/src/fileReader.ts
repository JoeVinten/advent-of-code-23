import fs from "fs";

const dataFilePath = new URL("data.txt", import.meta.url);
const data = fs.readFileSync(dataFilePath, "utf-8");

export { data };
