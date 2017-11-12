import matplotlib.pyplot as plt
import numpy as np

# Setteo titulo del grafico
plt.title('Go vs C++')

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
go_times = [0.54, 2.03, 14.44, 5.48, 2.17, 3.94, 21.37, 14.98, 28.49, 34.42]
cpp_times = [0.60, 1.89, 10.35, 1.73, 1.48, 1.99, 9.31, 7.18, 17.14, 2.36]

# Calculo cantidad de posiciones en el eje x
index = np.arange(len(benchs_names))

# Setteo variables del grafico de barras
bar_width = 0.2
opacity = 0.5

# Armo las barras de cada lenguaje
go_bars = plt.bar(index+bar_width, go_times, bar_width,
                  alpha=0.7,
                  color='g',
                  label='Go 1.9')

cpp_bars = plt.bar(index, cpp_times, bar_width,
                    alpha=0.7,
                    color='r',
                    label='C++ GCC 6.3.0')

plt.grid(True, axis='y', linewidth=0.2)
plt.xlabel('Benchmarks')
plt.ylabel('Tiempo (segundos)')
plt.xticks(index + bar_width / 2, benchs_names)
plt.yticks(np.arange(0, 37, 2))
plt.legend()

plt.tight_layout()
plt.show()
