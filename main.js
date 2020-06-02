const express = require("express");

const app = express();
const bodyparser = require("body-parser");

const port = process.env.PORT || 3000;

// middleware

app.use(bodyparser.json());
app.use(bodyparser.urlencoded({ extended: false }));

app.listen(port, () => {
   console.log(`running at port ${port}`);
});


app.get("/", (req, res) => {



   res.status(200).send("Hello World");


});

console.log('Server running at http://127.0.0.1:3000/');