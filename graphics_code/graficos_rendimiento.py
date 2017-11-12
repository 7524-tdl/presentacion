import matplotlib.pyplot as plt
import numpy as np

# Setteo titulo del grafico
plt.title('Go vs All')

# Indico cantidad de benchmarks y nombres
benchs_number = 10
benchs_names = ['Reverse\n Complement',
                'Pidigits',
                'Fannkuch\n redux',
                'Mandelbrot',
                'Fasta',
                'Spectral\n norm',
                'N-body',
                'K\n nucleotide',
                'Regex\n redux',
                'Binary\n trees']

# Armo arrays con los tiempos de cada lenguaje, en el orden de benchs_names
go_times     = [0.54, 2.03, 14.44, 5.48 , 2.17, 3.94 , 21.37, 14.98, 28.49, 34.42]
java_times   = [1.03, 3.12, 17.26, 6.04 , 2.33, 4.23 , 22.46, 8.70 , 10.34, 8.34 ]
cpp_times    = [0.60, 1.89, 10.35, 1.73 , 1.48, 1.99 , 9.31 , 7.18 , 17.14, 2.36 ]
node_times   = [4.00, 0.00, 81.26, 17.69, 9.40, 15.79, 27.93, 62.14, 4.20 , 55.26]

# Calculo cantidad de posiciones en el eje x
index = np.arange(len(benchs_names))

# Setteo variables del grafico de barras
bar_width = 0.2
opacity = 0.5

# Armo las barras de cada lenguaje
go_bars = plt.bar(index, go_times, bar_width,
                  alpha=0.7,
                  color='c',
                  label='Go 1.9')

java_bars = plt.bar(index+bar_width, java_times, bar_width,
                    alpha=0.7,
                    color='b',
                    label='Java 9.0.1')

cpp_bars = plt.bar(index+2*bar_width, cpp_times, bar_width,
                    alpha=0.7,
                    color='m',
                    label='C++ GCC 6.3.0')

node_bars = plt.bar(index+3*bar_width, node_times, bar_width,
                    alpha=0.7,
                    color='g',
                    label='NodeJS 8.8.1')

plt.grid(True, axis='y', linewidth=0.2)
plt.xlabel('Benchmarks')
plt.ylabel('Tiempo (segundos)')
plt.xticks(index + 2*bar_width, benchs_names)
plt.yticks(np.arange(0, 85, 2))
plt.legend()

plt.tight_layout()
plt.show()
