var express = require('express')
app = express();

// for post

var bodyParser = require('body-parser');
app.use(bodyParser.urlencoded({
    extended: true
}));
app.use(bodyParser.json());




app.get("/hello1", function(req, res) {
    console.log("Hello1");
    res.send("HelloWorld1")
});



app.get("/hello2", function(req, res) {
    console.log("Hello 2");
    res.send("Hello World2")
});

app.get("/hello/:param", function(req, res) {
    var get_param = req.params.param;
    console.log("Hello" + get_param);
    res.send("Hello World " + get_param)
});


// Other Restful method
app.post('/createUser', function(req, res) {
    console.log("create user")
    console.log(req.body.fn);
    console.log(req.body.ln)
    res.send(JSON.stringify({"success": true}))
})

app.put('/updateUser', function(req, res) {
    console.log("Update user")
    console.log(req.body.fn);
    console.log(req.body.ln)
    res.send(JSON.stringify({"success": true}))
})

app.delete('/deleteUser/:id', function(req, res) {
    console.log("Delete user")
    console.log(req.params.id);

    res.send(JSON.stringify({"success": true}))
})


app.use(express.static('static_dir'))

app.listen(3000, function () {
    console.log('Listening on port 3000');
  });