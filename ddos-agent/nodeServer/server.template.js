import express from "express";
import chalk from "chalk";
import dotenv from "dotenv";

const app = express();
dotenv.config();
app.use(express.json());

const PORT = process.env.PORT || 5001;

let count = 0;

app.get("/count", (_, res) => {
  res.json({
    "Number of hits:": count,
  });
});

app.get("/home", (_, res) => {
  count++;
  console.log(count);
  res.send("THIS IS /HOME ROUTE");
});

app.get("/api/home", (_, res) => {
  res.json({
    msg: "THIS IS /HOME ROUTE",
  });
});

app.post("/api/intro", (req, res) => {
  const { name } = req.body;
  if (!name) {
    return res.status(400).json({ msg: "Name is required in request body" });
  }
  res.status(200).json({
    msg: `Hello ${name}`,
  });
});

app.post("/api/introquery", (req, res) => {
  const { name } = req.query;
  if (!name) {
    return res.status(400).json({ msg: "Name is required in query params" });
  }
  res.status(200).json({
    msg: `Hello ${name}`,
  });
});

app.listen(PORT, () => {
  const url = `http://localhost:${PORT}`;
  console.log(`Server running at: ${chalk.green(url)}`);
});
