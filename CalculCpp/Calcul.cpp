#include <iostream>
#include <math.h>
#include <stdio.h>
#include <string>
#include <thread>

double doCalc(double x) { return x * x; }

void calculPortion(const double min, const double max, const unsigned long segments)
{
	double resultat, tmp;
	double x = min;

	double pas = (max - min) / (double) segments;

	for (unsigned long i = 0; i < segments; i++) 
	{
		resultat += doCalc(x) * pas;
		x += pas;
	}

	//printf("De %e à %e : %f\n", min, max, resultat);
}

int main(int argc, char *argv[])
{
	
	// verify command line arguments
	unsigned long nbSegments, nbCores;
	double        min, max;
	
	for (int i = 0 ; i < argc ; i++)
	{
		printf("%s\n", argv[i]);
	}

	if(argc == 5)
	{
		nbCores    = std::stol(argv[1]);
		min        = std::stod(argv[2]);
		max        = std::stod(argv[3]);
		nbSegments = std::stol(argv[4]);
	}
	else
	{
		printf("Usage : %s <nbCores> <min> <max> <nbSegments>\n", argv[0]);
		exit(1);
	}

	if(min > max) 
	{
		printf("Wrong input, maximum is inferior to minimum...");
		exit(2);
	}

	if (nbCores < 0) {
		printf("Impossible to use less than 1 core...");
		exit(3);
	}

	double increment = (max - min) / (double) nbCores;
	max = min + increment;

	// on crée autant de thread fils que de cœurs sur la machine
	std::thread tabThreads[nbCores];

	for(unsigned int i = 0; i < nbCores; i++)
	{
		// creation of child threads for computing
		// store it in an array for easier access
		tabThreads[i] = std::thread(calculPortion, min, max, nbSegments);

		min += increment;
		max += increment;
	}

	for(unsigned int i = 0 ; i < nbCores; i++)
	{
		tabThreads[i].join();
	}
}
