
console.log("Start")
//https://www.w3schools.com/html/html5_canvas.asp
function start() {
    var canvas = document.getElementById('canvas');
    if (!canvas) {
        alert('Error: Cannot find the canvas element!');
        return;
    }

    // 1. Let canvas fit its parents
    canvasParent = canvas.parentElement;
    width = canvasParent.offsetWidth;
    height = canvasParent.offsetHeight;    
    canvas.width = width;
    canvas.height = height;


    // 2. Check whether browser support canvas or not
    if (!canvas.getContext) {
        alert('Error: Canvas context does not exist!');
        return;
    }

    var ctx = canvas.getContext('2d');

    // Draw Rect
    ctx.fillStyle = "#3d3";
    ctx.fillRect(0, 0, 100, 100);

    // Draw line
    ctx.strokeStyle="#FF0000";
    ctx.lineWidth = 2;
    ctx.moveTo(0,0);
    ctx.lineTo(200,300);
    ctx.stroke()
}
start()




