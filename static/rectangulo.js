var canvas = document.getElementById("canvas");
var ctx = canvas.getContext("2d");

function drawRect(x,y) {
    ctx.fillStyle = "#FF0000";
    ctx.fillRect(x,y,y+150,x+75);
}


var pos = 0;
var lastTime = null;
function animate(time) {
    if (lastTime != null) {
        pos += (time - lastTime) * 0.1;
    }
    lastTime = time;
    ctx.clearRect(0,0,canvas.width,canvas.height);
    drawRect(pos, pos);
    requestAnimationFrame(animate);
}
requestAnimationFrame(animate);
