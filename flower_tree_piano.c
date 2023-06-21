#include <stdio.h> 
#include <stdlib.h> 
#include <string.h> 
#include <limits.h> 

// define some constants 
#define MAXCHAR 1000 
#define MAXNUM 100 

// structures for storing data 
typedef struct{ 
    char name[256]; 
    double x; 
    double y; 
} Planet; 

typedef struct{ 
    Planet planet; 
    double distance; 
} PlanetDistance; 

// function declarations 
double getDistance(Planet, Planet); 
PlanetDistance getNearestPlanet(Planet [], int, Planet); 
void printPlanet(Planet); 

// main program 
int main(int argc, char *argv[]){ 
    int numPlanets; 
    Planet planets[MAXNUM]; 
    PlanetDistance nearest; 
    Planet earth;
	
    // input data 
    printf("Enter the number of planets: "); 
    scanf("%d", &numPlanets); 
 
    for (int i=0; i<numPlanets; i++){ 
        printf("Enter planet name and its x and y coordinates: "); 
        scanf("%s %lf %lf", planets[i].name, &planets[i].x, &planets[i].y); 
    } 
 
    printf("Enter the Earth's x and y coordinates: "); 
    scanf("%lf %lf", &earth.x, &earth.y); 
 
    // calculate the nearest planet 
    nearest = getNearestPlanet(planets, numPlanets, earth); 
 
    // output data 
    printf("The nearest planet is "); 
    printPlanet(nearest.planet); 
    printf("with a distance of %f\n", nearest.distance); 
 
    return 0; 
} 

// get the distance between two planets 
double getDistance(Planet p1, Planet p2){ 
    double dx = p1.x - p2.x; 
    double dy = p1.y - p2.y; 
    return sqrt(dx*dx + dy*dy); 
} 

// get the nearest planet to a given planet 
PlanetDistance getNearestPlanet(Planet planets[], int numPlanets, Planet earth){ 
    double minDistance = INT_MAX; 
    PlanetDistance nearest; 
 
    // compare each planet to the given planet 
    for (int i=0; i<numPlanets; i++){ 
        // calculate the distance between them 
        double distance = getDistance(planets[i], earth); 
 
        // if the current planet is closer, update the data 
        if (distance < minDistance){ 
            minDistance = distance; 
            nearest.planet = planets[i]; 
            nearest.distance = distance; 
        } 
    } 
    
    return nearest; 
} 

// print the data for a given planet 
void printPlanet(Planet p){ 
    printf("%s (x: %lf, y: %lf)", p.name, p.x, p.y); 
}