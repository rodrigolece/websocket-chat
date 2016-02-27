var canvas = document.getElementById("canvas");
var ctx = canvas.getContext("2d");

var turnAngle = Math.PI/30;
var velFactor = 0.7;

var radiusParticle = 1/50;
var minD = 2*radiusParticle;
var lx = 1;
var ly = 1;

var gas = new Gas(lx - minD, ly - minD);


// La parte de la animación

var lastTime = null;
var Dt = 0;

function animate(time) {
    if (lastTime != null) {
        Dt = (time - lastTime) * 0.001;
    }
    lastTime = time;
    ctx.clearRect(0,0,canvas.width,canvas.height);
    gas.move(Dt);
    drawParticles(gas);
    requestAnimationFrame(animate);
}
requestAnimationFrame(animate);

function Gas(lx, ly) {
    this.numParticles = 0;
    this.particles = [];
    this.boxSize = { lx: lx, ly: ly };
    this.move = function(Dt) {
        for (var i = 0; i < this.numParticles; i++){
            this.particles[i].move(Dt);
        }
    }
    this.addParticle = function() {
        var index = this.particles.length;
        var pos = newPos(this);
        var part = new Particle(pos.x, pos.y, velFactor*(Math.random()*2-1),
            velFactor*(Math.random()*2-1), index);

        this.particles.push(part);
        this.numParticles += 1;

        return index
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

function Particle(posx, posy, velx, vely, index) {
    this.idx = index;
    this.pos = {x: posx, y: posy};
    this.vel = {vx: velx, vy: vely};
    this.move = moveParticle;
    this.changeVel = changeVel;
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

function changeVel(direction) {
    var angle = null;
    var factor = null;

    if (direction == "left") {
        // Cuidado! JS define los ángulos en sentido de las manecillas del reloj
        angle = -turnAngle;
    }
    if (direction == "right") {
        angle = turnAngle;
    }
    if (direction == "up") {
        factor = 1.02;
    }
    if (direction == "down") {
        factor = 0.98;
    }

    if (angle != null) {
        var c = Math.cos(angle);
        var s = Math.sin(angle);
        this.vel = {
            vx: c * this.vel.vx - s * this.vel.vy,
            vy: s * this.vel.vx + c * this.vel.vy
        }
    } else if (factor != null) {
        this.vel = {
            vx: this.vel.vx * factor,
            vy: this.vel.vy * factor
        }
    }
}
