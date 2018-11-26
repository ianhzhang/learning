function SimpleWaveChart(options) {
    var canvas = document.getElementById('canvas');
    if (!canvas) {
        alert('Error: Cannot find the canvas element!');
        return;
    }

    // 1. Let canvas fit its parents
    var canvasParent = canvas.parentElement;
    W = canvasParent.offsetWidth;
    H = canvasParent.offsetHeight;
    H_2 = H/2;    
    
    canvas.width = W;
    canvas.height = H;
    [this.x1,this.y1] = SimpleWaveChart.prototype.convertCoor(0,0)

    this.ctx = canvas.getContext('2d');
    this.ctx.strokeStyle="green";
    this.ctx.lineWidth = 3;

}
SimpleWaveChart.prototype.convertCoor = function(x,y) {
    // We can access width and height in outer function
    new_x = x;
    new_y = H_2 - y*H_2
    return [new_x, new_y]
}

SimpleWaveChart.prototype.line = function(x,y) {

    var [x2,y2] = SimpleWaveChart.prototype.convertCoor(x,y);

    if (x2 < this.x1) {
        this.x1 = 0;
        this.y1 = 0;
    }

    this.ctx.clearRect(x2, 0, 40, H);

    // Draw line

    this.ctx.beginPath()
    this.ctx.arc(this.x1,this.y1,1.8,0,2*Math.PI);
    this.ctx.fillStyle="green";
    this.ctx.fill();

    this.ctx.beginPath();

    this.ctx.moveTo(this.x1,this.y1);
    this.ctx.lineTo(x2,y2);
    this.ctx.stroke();
    
    this.x1 = x2;
    this.y1 = y2;

    this.ctx.beginPath();
    this.ctx.arc(x2,y2,1.8,0,2*Math.PI);
    this.ctx.fillStyle="red";
    this.ctx.fill();

}

SimpleWaveChart.prototype.clearChart = function() {
    this.ctx.clearRect(0,0,W,H);
    this.x1 = 0;
    this.y1 = 0;
}
