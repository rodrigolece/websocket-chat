package main

// import (
//
// )

type gas struct {
    numParticles int
    particles []*particle

    register chan *particle
	unregister chan *particle
}

type particle struct {
    // El índice nos sirve para identificar las partículas en JS
    idx int
    pos map[string]float64
    vel map[string]float64
}


func newGas() *gas {
    return &gas{
        numParticles: 0,
        particles: make([]*particle, 0),
        register: make(chan *particle),
        unregister: make(chan *particle),
    }
}

func (g *gas) run() {
    for {
        select {
        case part := <- g.register:
            // Le damos un índice a la partícula
            part.idx = g.getNewIndex()
            // Y la agregamos al gas
            g.addParticle(part)
        // case part := <- g.unregister:
        //
        }
    }
}

func (g *gas) addParticle(part *particle) {
    g.numParticles++
    g.particles = append(g.particles, part)
}

func (g *gas) getNewIndex() int {
    /* Esto se puede hacer más complicado para tener en cuenta que partículas
    de en medio de la lista van a salir */
    return g.numParticles
}
