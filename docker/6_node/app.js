var express = require('express')
    , app = express();

app.get("/api/hello", (req, res, next) => {
    console.log("get hello")
    res.json({"ln":"Zhang", "fn":"Ian"});
});

app.listen(5000, function () {
    // Extra processing after run
    console.log("Server running on port 3000")
});