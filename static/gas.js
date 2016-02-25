var canvas = document.getElementById("canvas");
var ctx = canvas.getContext("2d");

var numParticles = 1;
var radiusParticle = 1/50;
var minD = 2*radiusParticle;
var lx = 1;
var ly = 1;

var gas = new Gas(numParticles, lx - minD, ly - minD);
drawParticles(gas);

var lastTime = null;
var Dt = 0;

function animate(time) {
    if (lastTime != null) {
        Dt = (time - lastTime) * 0.001;
    }
    lastTime = time;
    ctx.clearRect(0,0,canvas.width,canvas.height)
    gas.move(Dt)
    drawParticles(gas);
    requestAnimationFrame(animate);
}
requestAnimationFrame(animate);

function Gas(numParticles, lx, ly) {
    this.numParticles = numParticles;
    this.particles = [];
    this.boxSize = {lx: lx, ly: ly};
    this.move = moveGas;

    for (var i = 0; i < numParticles; i++) {
        var pos = newPos(this);
        var part = new Particle(pos.x, pos.y, Math.random()*2-1,
        Math.random()*2-1, i);

        this.particles.push(part);
    }
}

function moveGas(Dt) {
    for (var i = 0; i < this.numParticles; i++){
        this.particles[i].move(Dt);
    }
}

function drawParticles(gas) {
    for (var i = 0; i < gas.numParticles; i++) {
        part = gas.particles[i];

        ctx.beginPath();
        ctx.arc(canvas.width * part.pos.x, canvas.height * part.pos.y,
            Math.min(canvas.width, canvas.height)*radiusParticle, 0, 2*Math.PI);
        ctx.stroke();
    }
}

function newPos(gas) {
    var n = gas.particles.length;
    do {
        var overlap = false;
        var x = gas.boxSize.lx * Math.random() + radiusParticle;
        var y = gas.boxSize.ly * Math.random() + radiusParticle;
        if (n == 0) {
            return { x: x, y: y };
        }
        for (var i = 0; i < n; i++){
            var part = gas.particles[i];
            var dx = part.pos.x - x;
            var dy = part.pos.y - y;
            var norm = Math.sqrt( dx*dx + dy*dy );
            if (norm <= minD) { overlap = true; }
        }
    } while (overlap) ;
    return { x: x, y: y };
}

function Particle(posx, posy, velx, vely, i) {
    this.pos = {x: posx, y: posy};
    this.vel = {vx: velx, vy: vely};
    this.move = moveParticle;
    this.index = i;
}

function moveParticle(Dt) {
    futurex = this.pos.x + this.vel.vx * Dt;
    futurey = this.pos.y + this.vel.vy * Dt;

    if (futurex + radiusParticle > lx || futurex - radiusParticle < 0) {
        this.vel.vx *= -1;
    }
    if (futurey + radiusParticle > ly || futurey - radiusParticle < 0) {
        this.vel.vy *= -1;
    }
    this.pos = {
        x: this.pos.x + this.vel.vx * Dt,
        y: this.pos.y + this.vel.vy * Dt
    };
}
