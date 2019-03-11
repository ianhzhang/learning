npm init
npm install json-server --saveDev
mkdir public
download db.json 
download images/ to public
json-server --watch db.json

./node_modules/.bin/json-server --watch db.json -d 2000
 -d 2000: cause 2 seconds delay

  http://localhost:3000/images/alberto.png

  http://localhost:3000/dishes
  http://localhost:3000/promotions
  http://localhost:3000/leaders
  http://localhost:3000/feedback


add:
"start":"./node_modules/.bin/json-server --watch db.json -d 2000"

npm run start
npm start
