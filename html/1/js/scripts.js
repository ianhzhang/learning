
( function () {
    console.log("Start 112")      
    myChart = new SimpleWaveChart();
    // myChart.line(50,0.4);
    var x = 0;
    var cnt = 0;
    var intervalId = null;
    start = function() {
        intervalId = setInterval( function() {
            y = Math.random()*2-1;
            y = y * 0.8;
    
            myChart.line(x, y);
            x = x + 10;
            if (x > 980 ) {
                x = 0;
                cnt = cnt + 1
                   
            }
            if (cnt > 50) {
                window.clearInterval(intervalId);
                intervalId = null;
                x = 0;
                cnt = 0;
            }
        }, 30)
    }

    stop = function() {
        if (intervalId) {
            window.clearInterval(intervalId);
        }
        intervalId = null;
        x = 0;
        cnt = 0;
    }

    clearRegion = function clearRegion() {
        myChart.clearChart();
    }
}());
