const fs = require('fs');
const readline = require('readline');


console.log('world');


const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

rl.question('Enter File path to open ', (answer) => {  
fs.open(answer, 'rr', (err, fd) => {
  if (err) throw err;
  fs.close(fd, (err) => {
    if (err) throw err;
  });
});


  rl.close();
});


var x;

